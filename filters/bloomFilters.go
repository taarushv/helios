package filters

import (
	"github.com/ethereum/go-ethereum/core/types"
)

// WIP
// Useful to check for things like oracle updates efficiently
// After each block is mined, we check if a particular event was emitted in the txs
// Bloom filters for efficient lookup
func CheckBlockForEvent(block *types.Block, topic []byte) bool {
	bloom := block.Bloom()
	//fmt.Println(bloom.MarshalText())
	//fmt.Println("1", topic)
	return bloom.TestBytes(topic)
}
