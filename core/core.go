package core

import "github.com/ethereum/go-ethereum/signer/storage"

type MonesignNode struct {

	Identity *peer.Peer

	PeerBook *peerbook.PeerBook

	Storage *storage.Storage

	Network *netmux.Netux

	Routing *routing.Routing

	Bitswap *bitswap.Bitswap

	Blocks *blocks.BlockService

	Resolver *resolver.PathResolver

	Namesys *namesys.Namesys
}