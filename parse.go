package hack_asm_go

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Parser struct {
	Lines   int
	Src     []string
	current int
}

const (
	A_COMMAND = "A"
	C_COMMAND = "C"
	L_COMMAND = "L"
)

func NewParser(scanner *bufio.Scanner) *Parser {
	var src []string
	var lines int
	var text string
	for scanner.Scan() {
		text = scanner.Text()
		text = strings.Replace(text, " ", "", -1) // remove space
		text = strings.Split(text, "//")[0]       // remove comment
		if text != "" {
			src = append(src, text)
			lines++
		}
	}
	return &Parser{
		Lines:   lines,
		Src:     src,
		current: 0,
	}
}

// まだコマンドが残っているか
func (parser Parser) HasMoreCommands() bool {
	return parser.Lines > parser.current
}

// 現在のコマンドをひとつ進める
func (parser *Parser) Advance() {
	parser.current++
}

// 現在のコマンド文字列を返す
func (parser Parser) CurrentCommand() string {
	return parser.Src[parser.current]
}

// コマンドの種類を返す
func (parser Parser) CommandType() string {
	var cmd = parser.CurrentCommand()
	switch {
	case cmd[0] == '@':
		return A_COMMAND
	case cmd[0] == '(':
		return L_COMMAND
	default:
		return C_COMMAND
	}
}

func (parser Parser) Symbol() string {
	if parser.CommandType() == A_COMMAND {
		return parser.CurrentCommand()[1:]
	} else if parser.CommandType() == L_COMMAND {
		return strings.Trim(parser.CurrentCommand(), "()")
	}
	return "NULL"
}

// 命令を分割する
// tokens[0] = cmp
// tokens[1] = dest
// tokens[2] = jmp
func (parser Parser) split(command string) [3]string {
	var tokens [3]string
	if strings.Index(command, "=") != -1 {
		tokens[1] = strings.Split(command, "=")[0]
		command = strings.Split(command, "=")[1]
	}
	if strings.Index(command, ";") != -1 {
		tokens[0] = strings.Split(command, ";")[0]
		tokens[2] = strings.Split(command, ";")[1]
	} else {
		tokens[0] = command
	}
	return tokens
}

func (parser Parser) Dest() string {
	tokens := parser.split(parser.CurrentCommand())
	return tokens[1]
}

func (parser Parser) Comp() string {
	tokens := parser.split(parser.CurrentCommand())
	return tokens[0]
}

func (parser Parser) Jump() string {
	tokens := parser.split(parser.CurrentCommand())
	return tokens[2]
}

// A命令のバイナリ表現を返す
func (parser Parser) Address() (string, error) {
	adrs, err := strconv.ParseInt(parser.Symbol(), 10, 16)
	return fmt.Sprintf("0%015b", adrs), err
}
