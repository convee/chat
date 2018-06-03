package tcp

import (
	"fmt"
	"net"
)

func Run(tcpAddr string) {
	fmt.Println("tcp server start running..")
	l, err := net.Listen("tcp", tcpAddr)
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

func read(conn net.Conn) {

}

func write(conn net.Conn)
