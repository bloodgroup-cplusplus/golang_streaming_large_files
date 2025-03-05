package p2p
import "net"

type TCPTransport struct {
	//make configurations 
	listenAddress string
	listener  net.Listender
	peers map[net.Addr ]
}