//go:build linux
// +build linux

package execsnoop

import (
	"bytes"
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/ahubaoan/emage/pkg/bpf"
	"github.com/ahubaoan/emage/pkg/bpf/module"
	"github.com/ahubaoan/emage/pkg/logger"
	"github.com/cilium/ebpf"
	"github.com/cilium/ebpf/link"
	"github.com/cilium/ebpf/ringbuf"
	"go.uber.org/zap"
	"golang.org/x/sys/unix"
)

// $BPF_CLANG and $BPF_CFLAGS are set by the Makefile.
//go:generate go run github.com/cilium/ebpf/cmd/bpf2go -cc $BPF_CLANG -cflags $BPF_CFLAGS bpf execsnoop.c -- -I../headers

type execSnoopBpfEvent struct {
	Pid      uint32
	PPid     uint32
	Comm     [16]uint8
	FileName [256]uint8
}

func RewriteConstants(obj *bpfObjects, spec *ebpf.CollectionSpec) error {
	rodata := spec.Maps[".rodata"]
	if rodata != nil && rodata.BTF != nil {
		err := spec.RewriteConstants(obj.Consts)
		if err != nil {
			fmt.Println("RewriteConstants err=", err.Error())
			return err
		}
	}
	return nil
}

func Start(ctx context.Context, c module.ExecSnoopKern) {

	// Load pre-compiled programs and maps into the kernel.
	objs := bpfObjects{Consts: bpf.KernConstantsGen(c)}
	if err := loadBpfObjects(&objs, nil); err != nil {
		logger.ComLog.Error("loading objects err", zap.Error(err), zap.Any("obj", objs))
		return
	}
	defer objs.Close()

	kp, err := link.Tracepoint("syscalls", "sys_enter_execve", objs.SysEnterExecve)
	if err != nil {
		logger.ComLog.Error("opening tracepoint", zap.Error(err))
		return
	}
	defer kp.Close()

	rd, err := ringbuf.NewReader(objs.Events)
	if err != nil {
		logger.ComLog.Error("opening ringbuf reader", zap.Error(err))
	}
	defer rd.Close()

	go func() {
		<-ctx.Done()

		if err := rd.Close(); err != nil {
			logger.ComLog.Error("closing ringbuf reader", zap.Error(err))
			return
		}
	}()

	logger.ComLog.Info("Waiting for events..")

	var event execSnoopBpfEvent
	for {
		record, err := rd.Read()
		if err != nil {
			if errors.Is(err, ringbuf.ErrClosed) {
				logger.ComLog.Error("Received signal, exiting..")
				return
			}
			continue
		}

		// Parse the ringbuf event entry into a bpfEvent structure.
		if err := binary.Read(bytes.NewBuffer(record.RawSample), binary.LittleEndian, &event); err != nil {
			logger.ComLog.Error("parsing ringbuf event", zap.Error(err))
			continue
		}

		logger.BpfLog.Info("execsnoop output",
			zap.String("bpf", "execsnoop"),
			zap.Uint32("pid", event.Pid),
			zap.Uint32("ppid", event.PPid),
			zap.String("comm", unix.ByteSliceToString(event.Comm[:])),
			zap.String("filename", unix.ByteSliceToString(event.FileName[:])))
	}

}
