package hack_asm_go

import (
	"bufio"
	"strings"
	"testing"
)

func TestNewParser(t *testing.T) {
	var code string
	code = `(START)
  @10
	M=D
D=1  // hogehoge
@START
D;JGT
`
	codeReader := strings.NewReader(code)
	scanner := bufio.NewScanner(codeReader)

	parser := NewParser(scanner)
	// lines
	if parser.Lines != 6 {
		t.Errorf("code lines 6, but %d", parser.Lines)
	}
	// remove space
	if parser.Src[1] != "@10" {
		t.Errorf("line 2: [@10], but [%s]", parser.Src[1])
	}
	// remove tab
	if parser.Src[2] != "M=D" {
		t.Errorf("line 3: [M=D], but [%s]", parser.Src[2])
	}
	// remove comment
	if parser.Src[3] != "D=1" {
		t.Errorf("line 4: [D=1], but [%s]", parser.Src[2])
	}
}

func TestHasMoreCommandsOk(t *testing.T) {
	parsed := Parser{
		Lines:   2,
		Src:     []string{"@10", "M=D"},
		current: 0,
	}
	actual := parsed.HasMoreCommands()
	expected := true
	if actual != expected {
		t.Errorf("commands is exists more, but false")
	}
}

func TestHasMoreCommandsFail(t *testing.T) {
	parsed := Parser{
		Lines:   2,
		Src:     []string{"@10", "M=D"},
		current: 2,
	}
	actual := parsed.HasMoreCommands()
	expected := false
	if actual != expected {
		t.Errorf("commands is no exists, but true")
	}
}

func TestAdvance(t *testing.T) {
	parsed := Parser{
		Lines:   2,
		Src:     []string{"@10", "M=D"},
		current: 0,
	}
	parsed.Advance()
	if parsed.current != 1 {
		t.Errorf("current=1, but %d", parsed.current)
	}
}

func TestReset(t *testing.T) {
	parsed := Parser{
		Lines:   2,
		Src:     []string{"@10", "M=D"},
		current: 2,
	}
	parsed.Reset()
	if parsed.current != 0 {
		t.Errorf("current=0, but %d", parsed.current)
	}
}

func TestCommandType(t *testing.T) {
	parsed := Parser{
		Lines: 3,
		Src:   []string{"@10", "M=D;JLT", "(HOGE)"},
	}

	parsed.current = 0
	actualA := parsed.CommandType()
	if actualA != A_COMMAND {
		t.Errorf("%s is A_COMMAND, but %s", parsed.Src[0], actualA)
	}

	parsed.current = 1
	actualC := parsed.CommandType()
	if actualC != C_COMMAND {
		t.Errorf("%s is C_COMMAND, but %s", parsed.Src[1], actualC)
	}

	parsed.current = 2
	actualL := parsed.CommandType()
	if actualL != L_COMMAND {
		t.Errorf("%s is L_COMMAND, but %s", parsed.Src[2], actualL)
	}
}

func TestSymbol(t *testing.T) {
	parsed := Parser{
		Lines: 3,
		Src:   []string{"@10", "(HOGE)", "@Variable"},
	}

	parsed.current = 0
	if "10" != parsed.Symbol() {
		t.Errorf("%s -> Symbol 10, but %s", parsed.Src[0], parsed.Symbol())
	}

	parsed.current = 1
	if "HOGE" != parsed.Symbol() {
		t.Errorf("%s -> Symbol HOGE, but %s", parsed.Src[1], parsed.Symbol())
	}

	parsed.current = 2
	if "Variable" != parsed.Symbol() {
		t.Errorf("%s -> Symbol Variable, but %s", parsed.Src[2], parsed.Symbol())
	}
}

func TestSplit(t *testing.T) {
	parsed := Parser{
		Lines: 3,
		Src:   []string{"MD=A", "D;JLT", "A=D;JGT"},
	}
	var tokens [3]string

	parsed.current = 0
	tokens = parsed.split()
	if tokens[0] != "A" || tokens[1] != "MD" || tokens[2] != "" {
		t.Errorf("%s split []string{A, MD, _}, but %s", parsed.CurrentCommand(), tokens)
	}

	parsed.current = 1
	tokens = parsed.split()
	if tokens[0] != "D" || tokens[1] != "" || tokens[2] != "JLT" {
		t.Errorf("%s split []string{D, _, JLT}, but %s", parsed.CurrentCommand(), tokens)
	}

	parsed.current = 2
	tokens = parsed.split()
	if tokens[0] != "D" || tokens[1] != "A" || tokens[2] != "JGT" {
		t.Errorf("%s split []string{D, A, JGT}, but %s", parsed.CurrentCommand(), tokens)
	}
}
