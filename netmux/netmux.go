package netmux

type Netmux struct {
	Interface []*Interface

	Incoming <-chan *Packet
	Outgoing chan<- *Packet

	incomingSrc chan<- *Packet
	outgoingSrc <-chan *Packet
}

type Packet struct {

	NetAddrTo string
	NetAddrFrom string
	Data []byte
}

func NewNetmux() *Netmux {
	n := &Netmux{}

	och := make(chan *Packet)
	ich := make(chan *Packet)

	n.Incoming, n.incomingSrc := ich, ich
	n.Outgoing, n.outgoingSrc := och, och

	return n
}