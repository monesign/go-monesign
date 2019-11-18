package merkledag

import (
	mh "github.com/multiformats/go-multihash"
)
// Monesign nodes have opaque data and navigable links
type Node struct {
	Links []*Link
	Data []byte
}

// Merkle DAG link
type Link struct {
	Name string
	Size uint64
	Hash mh.Multihash
}

type EncodedNode []byte

func (n *Node) Size() uint64 {
	 s uint64 = len(n.Encode())
	for _, l = range(n.Links) {
		s += l.Size
	}
//	no return ?
}
