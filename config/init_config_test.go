package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"testing"
)

func TestInitConfig(t *testing.T) {
	data, _ := ioutil.ReadFile("example_config.yaml")
	conf := Config{}
	err := yaml.Unmarshal(data, &conf)
	if err != nil {
		panic(any("Unmarshal config file err" + err.Error()))
	}

	info, err := yaml.Marshal(conf)
	if err != nil {
		panic(any(err.Error()))
	}
	fmt.Println(string(info))
}
