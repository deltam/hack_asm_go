package main

import (
	"bufio"
	"fmt"
	asm "github.com/deltam/hack_asm_go"
	"os"
)

func main() {
	var fp *os.File
	var err error

	if len(os.Args) < 2 {
		panic("Usage: Assembler Prog.asm")
	} else {
		fp, err = os.Open(os.Args[1])
		if err != nil {
			panic(err)
		}
		defer fp.Close()
	}

	scanner := bufio.NewScanner(fp)
	parser := asm.NewParser(scanner)

	// コマンド処理
	for ; parser.HasMoreCommands(); parser.Advance() {
		if parser.CommandType() == asm.A_COMMAND || parser.CommandType() == asm.L_COMMAND {
			adrs, err := parser.Address()
			if err == nil {
				fmt.Println(adrs)
			} else {
				panic("error address")
			}
		} else if parser.CommandType() == asm.C_COMMAND {
			cmd := asm.NewCode(*parser)
			fmt.Println(cmd.BinCode())
		}
	}
}
