package conf_test

import (
	"chat/conf"
	"testing"
)

func TestConfig(t *testing.T) {
	t.Log("test")
	conf.LoadTomlConfig("../config.toml")
	t.Log(conf.Get())
}

func Test_Mysql(t *testing.T) {
	conf.LoadTomlConfig("../config.toml")
	t.Log(conf.Get().Mysql["test"])
}
