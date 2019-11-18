package blocks

import (
	"github.com/monesign/go-monesign/bitswap"
	"github.com/monesign/go-monesign/storage"
)

type BlockService  struct {
	Local *storage.Storage
	Remote *bitswap.Bitswap
}