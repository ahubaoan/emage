package logger

import (
	"github.com/ahubaoan/emage/pkg/logger/logconf"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var ComLog *zap.Logger = nil
var BpfLog *zap.Logger = nil

// ZapLogInit
// If you want to send log to elasticsearch, use filebeat will be a good choice.
func ZapLogInit(conf *logconf.LogConf) *zap.Logger {
	encoderConfig := zap.NewProductionEncoderConfig()

	// time format, ISO8601 is '2006-01-02T15:04:05.000Z0700', and epoch is ''
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	switch conf.TimeType {
	case logconf.EncodeTimeTypeEpoch:
		encoderConfig.EncodeTime = zapcore.EpochTimeEncoder
	default:
		encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	}

	// format type
	encoder := zapcore.NewJSONEncoder(encoderConfig)
	if conf.FormatType == logconf.FormatTypeConsole {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	// output -> stdout
	var multiWs []zapcore.WriteSyncer
	if conf.OutStdout {
		ws := zapcore.AddSync(os.Stdout)
		multiWs = append(multiWs, ws)
	}
	// output -> file (auto logrotate)
	if conf.OutFileRotate.Enable {
		lumberJackLogger := &lumberjack.Logger{
			Filename:   conf.OutFileRotate.Filename,
			MaxSize:    conf.OutFileRotate.MaxSize,
			MaxBackups: conf.OutFileRotate.MaxBackups,
			MaxAge:     conf.OutFileRotate.MaxAge,
			Compress:   conf.OutFileRotate.Compress,
		}
		if lumberJackLogger.MaxSize <= 0 {
			lumberJackLogger.MaxSize = 10 // default 10M
		}
		if lumberJackLogger.MaxBackups <= 0 {
			lumberJackLogger.MaxBackups = 5
		}
		if lumberJackLogger.MaxAge <= 0 {
			lumberJackLogger.MaxAge = 60 // default 60 days
		}

		ws := zapcore.AddSync(lumberJackLogger)
		multiWs = append(multiWs, ws)
	}

	// aggregation writeSyners
	writeSyncer := zapcore.NewMultiWriteSyncer(multiWs...)
	if len(multiWs) == 1 {
		writeSyncer = multiWs[0]
	}

	core := zapcore.NewCore(encoder, writeSyncer, zapcore.Level(conf.DebugLevel))
	return zap.New(core)
}

func LogInit(conf *logconf.LogConf) {
	ComLog = ZapLogInit(conf)
}

func BpfLogInit(useComLog bool, conf *logconf.LogConf) {
	if useComLog {
		BpfLog = ComLog
		return
	}
	BpfLog = ZapLogInit(conf)
}
