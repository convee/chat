package conf

import (
	"testing"
)

func TestConfig(t *testing.T) {
	t.Log("test")
	LoadTomlConfig("../config.toml")

}
