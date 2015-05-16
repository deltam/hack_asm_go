package hack_asm_go

import (
	"strings"
)

type Command struct {
	comp string
	dest string
	jump string
}

func NewCode(parser Parser) *Command {
	return &Command{
		comp: parser.Comp(),
		dest: parser.Dest(),
		jump: parser.Jump(),
	}
}

func (cmd Command) BinCode() string {
	return "111" + cmd.Comp() + cmd.Dest() + cmd.Jump()
}

func (cmd Command) Comp() string {
	var aBit string
	var cmpBits string

	// aビットの判定
	if strings.Index(cmd.comp, "M") > -1 {
		aBit = "1"
	} else {
		aBit = "0"
	}

	compCode := make(map[string]string)
	compCode["0"] = "101010"
	compCode["1"] = "111111"
	compCode["-1"] = "111010"
	compCode["D"] = "001100"
	compCode["A"] = "110000"
	compCode["!D"] = "001101"
	compCode["!A"] = "110001"
	compCode["-D"] = "001111"
	compCode["-A"] = "110011"
	compCode["D+1"] = "011111"
	compCode["A+1"] = "110111"
	compCode["D-1"] = "001110"
	compCode["A-1"] = "110010"
	compCode["D+A"] = "000010"
	compCode["D-A"] = "010011"
	compCode["A-D"] = "000111"
	compCode["D&A"] = "000000"
	compCode["D|A"] = "010101"

	// M
	compCode["M"] = compCode["A"]
	compCode["!M"] = compCode["!A"]
	compCode["-M"] = compCode["-A"]
	compCode["M+1"] = compCode["A+1"]
	compCode["M-1"] = compCode["A-1"]
	compCode["D+M"] = compCode["D+A"]
	compCode["D-M"] = compCode["D-A"]
	compCode["M-D"] = compCode["A-D"]
	compCode["D&M"] = compCode["D&A"]
	compCode["D|M"] = compCode["D|A"]

	bin, ok := compCode[cmd.comp]
	if ok {
		cmpBits = bin
	} else {
		panic("parse error")
	}

	return aBit + cmpBits
}

func (cmd Command) Dest() string {
	var a, m, d = "0", "0", "0"

	if strings.Index(cmd.dest, "A") > -1 {
		a = "1"
	}
	if strings.Index(cmd.dest, "M") > -1 {
		m = "1"
	}
	if strings.Index(cmd.dest, "D") > -1 {
		d = "1"
	}

	return a + d + m
}

func (cmd Command) Jump() string {
	jumpCode := make(map[string]string)
	jumpCode["JGT"] = "001"
	jumpCode["JEQ"] = "010"
	jumpCode["JGE"] = "011"
	jumpCode["JLT"] = "100"
	jumpCode["JNE"] = "101"
	jumpCode["JLE"] = "110"
	jumpCode["JMP"] = "111"

	bin, ok := jumpCode[cmd.jump]
	if !ok {
		bin = "000"
	}

	return bin
}
