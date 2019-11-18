package netmux

import "net"

type Interface struct {
	Network string
	Address string
	ResolvedAddress string

	conn *net.Conn

	Input chan *Packet
	Output chan *Packet
	Closed chan bool
	Errors chan error
}

func NewUDPInterface(net, addr string) (*Interface, error) {
	raddr, err := net.ResolveUDPAddr(net, addr)
	if err != nil {
		return nil, err
	}

	conn, err := net.ListenUDP(net, addr)
	if err != nil {
		return nil, err
	}

	i := &Interface{
		Network: net,
		Address: addr,
		ResolvedAddress: raddr,
		conn: conn,
	}

	go i.processInput()
	go i.processOutput()
	return i, nil
}

func (i *Interface) processOutput(){
//
}

func (i *Interface) processUDPInput(){

}

func (i *Interface) Read(buffer []byte) bool {
	n, err := i.Conn.Read(buffer)
	if err != nil {
		i.Errors <- err
		i.Close()
		return false
	}
	return true
}

func (i *Interface) Close() {

	err := i.conn.Close()
	if err != nil {
		i.Errors <- err
	}

	close(i.Input)
	close(i.Output)
	close(i.Closed)
	close(i.Errors)
}