package logger

import (
	"github.com/ahubaoan/emage/pkg/log/logconf"
	"go.uber.org/zap"
	"testing"
	"time"
)

func TestOutput(t *testing.T) {
	conf := logconf.LogConf{
		TimeType:   0,
		FormatType: "",
		DebugLevel: 0,
		OutStdout:  false,
		OutFileRotate: logconf.OutFileRotate{
			Enable:     false,
			Filename:   "",
			MaxSize:    0,
			MaxBackups: 0,
			MaxAge:     0,
			Compress:   false,
		},
	}
	ZapLogInit(conf)
	for {
		Log.Info("aaa", zap.String("test", "test"))
		Log.Error("1234", zap.String("num", "999"))
		time.Sleep(1 * time.Millisecond)
	}
}
