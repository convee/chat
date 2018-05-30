package main

import (
	"chat/conf"
	"fmt"
)

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	conf.LoadTomlConfig("config.toml")

}
