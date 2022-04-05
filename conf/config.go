package conf

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Port   int           `yaml:"Port"`
	Themis *ThemisConfig `yaml:"themis"`
}

type ThemisConfig struct {
	ServerName string `yaml:"server_name"`
	Address    string `yaml:"address"`
}

var config = &Config{}

func Setup() {
	f, err := ioutil.ReadFile("./conf/app.yml")
	if err != nil {
		panic(err)
	}

	if err := yaml.Unmarshal(f, config); err != nil {
		panic(err)
	}
}

func GetThemisServer() (string, string) {
	return config.Themis.ServerName, config.Themis.Address
}
