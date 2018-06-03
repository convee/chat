package main

import (
	"chat/conf"
	"chat/handle/tcp"
	"flag"
	"fmt"
)

var (
	tcpAddr = flag.String("tcpAddr", "127.0.0.1:5001", "tcp address")
)

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	flag.Parse()
	conf.LoadTomlConfig("config.toml")
	tcp.Run(&tcpAddr)
}
