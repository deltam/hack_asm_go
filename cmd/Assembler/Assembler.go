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

    fmt.Println(parser)

    // コマンド処理
    for ; parser.HasMoreCommands(); parser.Advance() {
        fmt.Println(parser.CurrentCommand())
        fmt.Println("type: " + parser.CommandType())
        if parser.CommandType() == asm.A_COMMAND || parser.CommandType() == asm.L_COMMAND {
            fmt.Println("address: " + parser.Symbol())
        }
        if parser.CommandType() == asm.C_COMMAND {
            fmt.Println("cmp: " + parser.Comp())
            fmt.Println("dst: " + parser.Dest())
            fmt.Println("jmp: " + parser.Jump())
            cmd := asm.NewCode(*parser)
            fmt.Println("BINARY:  " + cmd.BinCode())
        }
    }
}
