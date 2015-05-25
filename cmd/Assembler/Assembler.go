package main

import (
	"bufio"
	"flag"
	"fmt"
	asm "github.com/deltam/hack_asm_go"
	"os"
	"strings"
)

func usage() {
	fmt.Fprintln(os.Stderr, `
Usage:
  Assembler [-s] Prog.asm

Flags:`)
	flag.PrintDefaults()
}

func main() {
	var fp *os.File
	var outfp *os.File
	var err error

	// command line option
	var s = flag.Bool("s", false, "Output to STDOUT")
	flag.Parse()
	// args
	var args = flag.Args()
	// custom usage
	flag.Usage = usage

	if flag.NArg() == 1 {
		fp, err = os.Open(args[0])
		if err != nil {
			panic(err)
		}
		defer fp.Close()
		// output to stdout
		if *s {
			outfp = os.Stdout
		} else {
			// output to file
			fInfo, _ := fp.Stat()
			outfile := strings.TrimRight(fInfo.Name(), ".asm") + ".hack"
			outfp, err = os.Create(outfile)
			if err != nil {
				panic(err)
			}
			defer outfp.Close()
		}
	} else {
		flag.Usage()
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
