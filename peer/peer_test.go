package peer

import (
	ma "github.com/multiformats/go-multiaddr"
	mh "github.com/multiformats/go-multihash"
	"testing"
)

func TestNetAddress(t *testing.T){
	tcp, err := ma.NewMultiaddr("/ip4/127.0.0.1/tcp/1234")
	if err != nil {
		t.Error(err)
		return
	}

	udp, err := ma.NewMultiaddr("ip4/127.0.0.1/udp/2345")
	if err != nil {
		t.Error(err)
		return
	}

	mh, err := mh.FromHexString("11140beec7b5ea3f0fdbc95d0dd47f3c5bc275da8a33")
	if err != nil {
		t.Error(err)
		return
	}

	p := Peer{Id: PeerId(mh)}
	p.AddAddress(tcp)
	p.AddAddress(udp)

	tcp2 := p.NetAddress("tcp")
	if tcp2 != tcp {
		t.Error("NetAddress lookup failed", tcp, tcp2)
	}

	udp2 := p.NetAddress("udp")
	if udp2 != udp {
		t.Error("NetAddress lookup failed", udp, udp2)
	}
}