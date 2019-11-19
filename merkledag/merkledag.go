package merkledag

import (
	mh "github.com/multiformats/go-multihash"
)
// Monesign nodes have opaque data and navigable links
type Node struct {
	Links []*Link
	Data []byte

	// cache encoded/marshaled value
	encoded []byte
}

// Merkle DAG link
type Link struct {
	Name string
	Size uint64
	Hash mh.Multihash
}

func (n *Node) AddNodeLink(name string, that *Node) error {
	s, err := that.Size()
	if err != nil {
		return err
	}

	h, err := that.Multihash()
	if err != nil {
		return err
	}

	n.Links = append(n.Links, &Link{
		Name: name,
		Size: s,
		Hash: h,
	})
	return nil
}

// to calculate the size of links in a node
func (n *Node) Size() (uint64, error) {
	 b, err := n.Encoded(false)
	 if err != nil {
	 	return 0, err
	 }
	 s := uint64(len(b))
	 for _, l := range(n.Links){
	 	s += l.Size
	 }
	return s, nil
}

func (n *Node) Multihash() (mh.Multihash, error) {
	b, err := n.Encoded(false)
	if err != nil {
		return nil, err
	}

	return mh.Sum(b, mh.SHA2_256, -1)
}

// protoc -I=. -I=$GOPATH/src -I=$GOPATH/src/github.com/gogo/protobuf/protobuf --{binary}_out=. node.proto

// protoc --go_out=. **/*.proto
