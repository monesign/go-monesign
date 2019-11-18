package block

import (
	_ "github.com/multiformats/go-multihash"
)

type Block struct {
	Multihash []byte
	Data []byte
}