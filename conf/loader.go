package conf

import (
	"github.com/BurntSushi/toml"
)

var config Config

func InitConfig(path string) {
	if _, err := toml.DecodeFile(path, &config); err != nil {
		panic(err)
	}
}

func GetConfig() Config {
	return config
}
