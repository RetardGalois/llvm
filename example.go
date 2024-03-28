package main

import (
	"log"

	"github.com/kr/pretty"
	"github.com/llir/llvm/asm"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/opt"
)

func main() {
	// Parse the LLVM IR assembly file `rand.ll`.
	m, err := asm.ParseFile("testdata/test.ll")
	if err != nil {
		log.Fatalf("%+v", err)
	}
	// Pretty-print the data types of the parsed LLVM IR module.

	builder := opt.NewBuilder(m.Funcs[0].Blocks[0])
	for i, inst := range m.Funcs[0].Blocks[0].Insts {
		switch inst.(type) {
		case *ir.InstAdd:
			builder.SetInsertPoint(opt.ABOVE, i)
			teste := builder.AddFuncCall(m.Funcs[1])
			builder.AddFuncCall(m.Funcs[0], teste)
		}
	}

	//teste := m.Funcs[0].Blocks[0].NewStartCall(m.Funcs[1])
	//builder.AddFuncCallAtStart(m.Funcs[1])

	pretty.Println(m.String())
}
