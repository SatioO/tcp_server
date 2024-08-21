package p2p

import (
	"fmt"
	"net"
)

type TCPTransportOptions struct {
	Addr     string
	Listener net.Listener
	Decoder  Decoder
}

type TCPTransport struct {
	TCPTransportOptions
}

func NewTCPTransport(config TCPTransportOptions) Transport {
	return &TCPTransport{
		TCPTransportOptions: config,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error
	t.Listener, err = net.Listen("tcp", t.Addr)

	if err != nil {
		return err
	}

	go t.acceptConn()

	return nil
}

func (t *TCPTransport) acceptConn() {
	for {
		conn, err := t.Listener.Accept()

		if err != nil {
			fmt.Println("error while accepting new connection:", err)
			break
		}

		go t.handleConn(conn)
	}
}

func (t *TCPTransport) handleConn(conn net.Conn) {
	msg := &Message{}

	for {
		if err := t.Decoder.Decode(conn, msg); err != nil {
			fmt.Println("error decoding msg: ", err)
		}

		msg.From = conn.RemoteAddr()

		fmt.Printf("recieved: %+v\n", msg)
	}
}
