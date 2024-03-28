package opt

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/value"
)

type InsertMode int

const (
	NONE InsertMode = iota
	ABOVE
	BELOW
	REPLACE
)

type Builder struct {
	block       *ir.Block
	mode        InsertMode
	entry_point int
}

func NewBuilder(block *ir.Block) *Builder {
	return &Builder{block: block, mode: NONE, entry_point: 0}
}

func (b *Builder) SetInsertPoint(mode InsertMode, ep int) {
	b.mode = mode
	b.entry_point = ep
}

func (b *Builder) AddFuncCall(callee value.Value, args ...value.Value) *ir.InstCall {
	call := ir.NewCall(callee, args...)
	if b.mode == ABOVE {
		b.block.Insts = append(b.block.Insts[:b.entry_point+1], b.block.Insts[b.entry_point:]...)
		b.block.Insts[b.entry_point] = call
		b.block.Parent.ReAssignIDs()
		b.entry_point++
		return call
	} else if b.mode == BELOW {
		b.block.Insts = append(b.block.Insts, nil)
		copy(b.block.Insts[b.entry_point+1:], b.block.Insts[b.entry_point:])
		b.block.Insts[b.entry_point+1] = call
		b.block.Parent.ReAssignIDs()
		return call
	} else if b.mode == REPLACE {
		b.block.Insts[b.entry_point] = call
		b.block.Parent.ReAssignIDs()
		return call
	}

	return nil
}
