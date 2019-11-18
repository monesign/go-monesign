package bitswap

import (
	"github.com/monesign/go-monesign/peer"
	"github.com/multiformats/go-multihash"
)

type Ledger struct {
//	todo
}

type Bitswap struct {
	Ledgers map[peer.PeerId]*Ledger
	//HaveList map[*multihash.Multihash]*block.Block
	HaveList map[multihash.Multihash]*block.Block
	WantList []*multihash.Multihash
}