package constant

import (
	"github.com/llir/l/ir/types"
	"github.com/llir/l/ir/value"
)

// BlockAddress is a blockaddress constant.
type BlockAddress struct {
	// Parent function.
	Func value.Value // *ir.Function
	// Basic block to take address of.
	Block value.Value // *ir.BasicBlock
}

// NewBlockAddress returns a new blockaddress constant based on the given parent
// function and basic block.
func NewBlockAddress(f, block value.Value) *BlockAddress {
	return &BlockAddress{Func: f, Block: block}
}

// Type returns the type of the constant.
func (c *BlockAddress) Type() types.Type {
	panic("not yet implemented")
}

// Ident returns the identifier associated with the constant.
func (c *BlockAddress) Ident() string {
	panic("not yet implemented")
}
