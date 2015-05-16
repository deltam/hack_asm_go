package hack_asm_go

import (
	"fmt"
	"strconv"
)

type SymbolTable struct {
	address map[string]int16
}

func NewSymbolTable() *SymbolTable {
	table := make(map[string]int16)

	// default
	table["SP"] = 0
	table["LCL"] = 1
	table["ARG"] = 2
	table["THIS"] = 3
	table["THAT"] = 4
	table["SCREEN"] = 16384
	table["KBD"] = 24576
	for i := 0; i < 16; i++ {
		table[fmt.Sprintf("R%d", i)] = int16(i)
	}

	return &SymbolTable{
		address: table,
	}
}

// シンボル：アドレスのテーブルを作る
func (table *SymbolTable) MakeTable(parser Parser) {
	for ; parser.HasMoreCommands(); parser.Advance() {
		if parser.CommandType() == L_COMMAND && !table.Contains(parser.Symbol()) {
			table.AddEntry(parser.Symbol(), int16(parser.CurrentLine()))
		}
	}
	parser.Reset()
}

func (table *SymbolTable) AddEntry(symbol string, address int16) {
	table.address[symbol] = address
}

func (table SymbolTable) Contains(symbol string) bool {
	_, contain := table.address[symbol]
	return contain
}

func (table SymbolTable) GetAddress(symbol string) (string, error) {
	if table.Contains(symbol) {
		val, _ := table.address[symbol]
		address := fmt.Sprintf("0%015d", val)
		return address, nil
	} else {
		address, err := strconv.ParseInt(symbol, 10, 16)
		return fmt.Sprintf("0%015b", address), err
	}
}
