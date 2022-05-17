GO := go
GO_BUILD = CGO_ENABLED=0 $(GO) build
GO_GENERATE = $(GO) generate
GO_TAGS ?=
TARGET=BPFDemo
BINDIR ?= /usr/local/bin
VERSION=$(shell git describe --tags --always)

$(TARGET):
	cd bpf/kprobe;$(GO_GENERATE)
#	$(GO_BUILD) $(if $(GO_TAGS),-tags $(GO_TAGS)) \
		-ldflags "-w -s \
		-X 'github.com/SimpCosm/godemo/ebpf/BPFDemo.Version=${VERSION}'"
		
clean:
	rm -f $(TARGET)
	rm -f bpfdemo_bpf*
	rm -rf ./release
