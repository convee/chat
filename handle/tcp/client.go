package tcp

import (
	"chat/pb"
	"net"
)

type Client struct {
	conn net.Conn
	buf  [8192]byte
}

func (p *Client) read() {
	n, err := p.conn.Read(p.buf[0:4])
	if n != 4 {

	}
}

func (p *Client) Process() {
	for {
		var msg pb.Message
		msg, err = p.read()
	}
}
