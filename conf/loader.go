package conf

import (
	"github.com/BurntSushi/toml"
)

var conf Config

func LoadTomlConfig(path string) {
	if _, err := toml.DecodeFile(path, &conf); err != nil {
		panic(err)
	}
	Set(conf)
}

// var configs map[string]interface{}

func Set(config Config) {
	conf = config
}

func Get() Config {
	return conf
}
