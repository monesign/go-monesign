package peer

import (
	mh "github.com/multiformats/go-multihash"
	ma "github.com/multiformats/go-multiaddr"
)

type PeerId mh.Multihash

type PeerBook map[string]*Peer

type Peer struct {
	Id mh.Multihash
	Addresses []*ma.Multiaddr
}

func (p *Peer) AddAddress(a *ma.Multiaddr){
	p.Addresses = append(p.Addresses, a)
}

func (p *Peer) NetAddress(n string) *ma.Multiaddr{
	for _, a := range p.Addresses{
		ps, err = a.Protocols()
		if err != nil {
			continue // address is invalid
		}
	}

	for _, p := range ps {
		if p.Name == n {
			return a
		}
	}

	return nil
}

