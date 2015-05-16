package main

import (
	"bufio"
	"fmt"
	asm "github.com/deltam/hack_asm_go"
	"os"
	"strings"
)

func main() {
	var fp *os.File
	var outfp *os.File
	var err error

	if len(os.Args) == 2 || len(os.Args) == 3 {
		fp, err = os.Open(os.Args[1])
		if err != nil {
			panic(err)
		}
		defer fp.Close()
		// output to stdout
		if len(os.Args) == 3 && os.Args[2] == "-s" {
			outfp = os.Stdout
		} else {
			// output to file
			name := strings.Split(os.Args[1], "/")
			outfile := strings.TrimRight(name[len(name)-1], "asm") + "hack"
			outfp, err = os.Create(outfile)
			if err != nil {
				panic(err)
			}
			defer outfp.Close()
		}
	} else {
		fmt.Println("Usage: Assembler Prog.asm [-s]\n" + "-s output to STDOUT")
		os.Exit(1)
	}

	scanner := bufio.NewScanner(fp)
	parser := asm.NewParser(scanner)

	symbols := asm.NewSymbolTable()
	symbols.MakeTable(*parser)

	// コマンド処理
	for ; parser.HasMoreCommands(); parser.Advance() {
		//		fmt.Println(parser.CurrentCommand())
		if parser.CommandType() == asm.A_COMMAND {
			address := symbols.GetAddress(parser.Symbol())
			if err == nil {
				fmt.Fprintln(outfp, address)
			} else {
				panic("error address")
			}
		} else if parser.CommandType() == asm.C_COMMAND {
			cmd := asm.NewCode(*parser)
			fmt.Fprintln(outfp, cmd.BinCode())
		}
	}
}
