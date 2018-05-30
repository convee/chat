package tcp

import (
	"chat/conf"
	"fmt"
	"net"
)

func Run() (err error) {
	fmt.Println("tcp server start running..")
	addr := conf.Get().Tcp.Address
	l, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			continue
		}
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
}
