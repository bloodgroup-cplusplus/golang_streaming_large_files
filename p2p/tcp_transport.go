package p2p

import (
	"net"
	"sync"
)

type TCPTransport struct {
	//make configurations
	listenAddress string
	listener      net.Listener
	// mutex for peer lock
	mu sync.RWMutex
	// put your mutex above the thing you want to protect it
	peers map[net.Addr]Peer
}

func NewTCPTransport(listenAddr string) *TCPTransport {
	return &TCPTransport{
		listenAddress: listenAddr,
	}

}
