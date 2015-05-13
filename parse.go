package hack_asm_go

import (
    "bufio"
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
        text = strings.Split(text, "//")[0] // remove comment
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

func (parser Parser) Dest() string {
    if strings.Index(parser.CurrentCommand(), "=") != -1 {
        return strings.Split(parser.CurrentCommand(), "=")[0]
    }
    return ""
}

func (parser Parser) Comp() string {
    if strings.Index(parser.CurrentCommand(), "=") != -1 {
        cmpJmp := strings.Split(parser.CurrentCommand(), "=")[1]
        return strings.Split(cmpJmp, ";")[0]
    }
    return ""
}

func (parser Parser) Jump() string {
    if strings.Index(parser.CurrentCommand(), ";") != -1 {
        return strings.Split(parser.CurrentCommand(), ";")[1]
    }
    return ""
}
