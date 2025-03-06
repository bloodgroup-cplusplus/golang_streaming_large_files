package p2p

import (
	"fmt"
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

func (t *TCPTransport) ListenAndAccept() error {

	var err error

	t.listener, err = net.Listen("tcp", t.listenAddress)
	if err != nil {
		return err
	}
	go t.startAcceptLoop()
	return nil

}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("TCP accept error: %s\n", err)
		}
		go t.handleConn(conn)
	}
}

func (t *TCPTransport) handleConn(conn net.Conn) {
	fmt.Printf(" New incoming connection%+v\n", conn)

}
