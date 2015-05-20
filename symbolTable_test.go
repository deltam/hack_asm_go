package hack_asm_go

import (
	"testing"
)

func TestMakeTable(t *testing.T) {
	parsed := Parser{
		Src:     []string{"(TEST)"},
		Lines:   1,
		current: 0,
	}
	table := NewSymbolTable()
	table.MakeTable(parsed)
	actual := table.address["TEST"]
	if 0 != table.address["TEST"] {
		t.Errorf("first line LABEL is address 0, but %d", actual)
	}
	if 0 != parsed.current {
		t.Errorf("after MakeTable(), current line is reset, but line %d", parsed.current)
	}
}

func TestGetAddress(t *testing.T) {
	parsed := Parser{
		Src:     []string{"(TEST)", "@123", "M=1", "@Variable1", "M=1", "@TEST", "M;JGT", "@Variable2", "M=A"},
		Lines:   9,
		current: 0,
	}
	table := NewSymbolTable()
	table.MakeTable(parsed)
	// label
	actual1 := table.GetAddress("TEST")
	// 0
	if "0000000000000000" != actual1 {
		t.Errorf("Label (TEST) is themself address(0), but %d", actual1)
	}
	// A Command
	actual2 := table.GetAddress("123")
	// 123
	if "0000000001111011" != actual2 {
		t.Errorf("A Command and Number, that is Address Number. but %d", actual2)
	}
	// Variable
	actual3 := table.GetAddress("Variable1")
	// 16
	if "0000000000010000" != actual3 {
		t.Errorf("Variable is starting address 16, but %d", actual3)
	}
	// 16 + 1
	actual4 := table.GetAddress("Variable2")
	if "0000000000010001" != actual4 {
		t.Errorf("Variable is starting address 16, but %d", actual4)
	}
}
