package main

import (
	"chat/conf"
	"chat/db"
	"chat/handle/tcp"
	"fmt"
)

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	conf.LoadTomlConfig("config.toml")
	db.Pool("test")
	tcp.Run()
}
