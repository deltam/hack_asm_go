package hack_asm_go

import (
	"fmt"
	"strconv"
)

type SymbolTable struct {
	address       map[string]int
	variableCount int
}

func NewSymbolTable() *SymbolTable {
	table := make(map[string]int)

	// default
	table["SP"] = 0
	table["LCL"] = 1
	table["ARG"] = 2
	table["THIS"] = 3
	table["THAT"] = 4
	table["SCREEN"] = 16384
	table["KBD"] = 24576
	for i := 0; i < 16; i++ {
		table[fmt.Sprintf("R%d", i)] = i
	}

	return &SymbolTable{
		address:       table,
		variableCount: 16,
	}
}

// シンボル->アドレスのテーブルを作る
func (table *SymbolTable) MakeTable(parser Parser) {
	var symbolCount = 0
	for ; parser.HasMoreCommands(); parser.Advance() {
		if parser.CommandType() == L_COMMAND && !table.Contains(parser.Symbol()) {
			table.AddEntry(parser.Symbol(), parser.CurrentLine()-symbolCount)
			symbolCount++
		}
	}
}

func (table *SymbolTable) AddEntry(symbol string, address int) {
	table.address[symbol] = address
}

func (table *SymbolTable) AddVariable(symbol string) {
	table.address[symbol] = table.variableCount
	table.variableCount++
}

func (table SymbolTable) Contains(symbol string) bool {
	_, contain := table.address[symbol]
	return contain
}

func (table *SymbolTable) GetAddress(symbol string) string {
	var val int
	if table.Contains(symbol) {
		val, _ = table.address[symbol]
	} else {
		_, err := strconv.ParseInt(symbol, 10, 16)
		if err != nil {
			table.AddVariable(symbol)
			val, _ = table.address[symbol]
		} else {
			tmp, _ := strconv.ParseInt(symbol, 10, 16)
			val = int(tmp)
		}
	}
	return fmt.Sprintf("0%015b", val)
}
