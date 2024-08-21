package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransportSuccess(t *testing.T) {
	tcpConfig := TCPTransportOptions{
		Addr: ":4000",
	}
	tr := NewTCPTransport(tcpConfig).(*TCPTransport)

	assert.Equal(t, tcpConfig.Addr, tr.Addr)
	assert.Nil(t, tr.ListenAndAccept())
}

func TestTCPTransportFailure(t *testing.T) {
	tcpConfig := TCPTransportOptions{
		Addr: ":4000",
	}
	tr := NewTCPTransport(tcpConfig).(*TCPTransport)

	assert.Error(t, tr.ListenAndAccept())
}
