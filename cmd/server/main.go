package main

import (
	"log"

	"github.com/satioO/tcp/v2/p2p"
)

func main() {
	tcpConfig := p2p.TCPTransportOptions{Addr: ":4000", Decoder: p2p.DefaultDecoder{}}
	tr := p2p.NewTCPTransport(tcpConfig)
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}
}
