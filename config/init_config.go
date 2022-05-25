package config

import (
	"github.com/ahubaoan/emage/pkg/bpf/module"
	"github.com/ahubaoan/emage/pkg/logger"
	"github.com/ahubaoan/emage/pkg/logger/logconf"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Config struct {
	Log    logconf.LogConf `json:"log" yaml:"log"`
	BpfLog logconf.LogConf `json:"bpf_log" yaml:"bpf_log"`
	Kern   module.BpfKern  `json:"kern" yaml:"kern"`
}

var GlobalConfigFilePath = "./example_config.yaml"
var GlobalConfig Config

func InitConfig(FileName string) {
	GlobalConfigFilePath = FileName
	data, err := ioutil.ReadFile(GlobalConfigFilePath)
	if err != nil {
		panic(err)
	}
	if err = yaml.Unmarshal(data, &GlobalConfig); err != nil {
		panic(err)
	}
	logger.LogInit(&GlobalConfig.Log)
	logger.BpfLogInit(GlobalConfig.BpfLog.UseLog, &GlobalConfig.BpfLog)
}

func InitConfigDefault() {
	conf := &logconf.LogConf{
		TimeType:   1, // iso8601
		FormatType: "json",
		DebugLevel: 0,
		OutStdout:  true,
		OutFileRotate: logconf.OutFileRotate{
			Enable: false,
		},
	}
	logger.LogInit(conf)
	logger.BpfLogInit(true, conf)
}
