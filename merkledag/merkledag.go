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

// to calculate the size of links in a node
func (n *Node) Size() (uint64, error) {
	 d, err := n.Marshal()
	 if err != nil {
	 	return 0, err
	 }
	 s := uint64(len(d))
	 for _, l := range(n.Links){
	 	s += l.Size
	 }
	return s, nil
}

// protoc -I=. -I=$GOPATH/src -I=$GOPATH/src/github.com/gogo/protobuf/protobuf --{binary}_out=. node.proto
