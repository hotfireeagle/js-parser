package parser

import (
	"fmt"
	"jsj/frontend/token"
)

type TokenEvent struct {
	// source line number
	LineNumber int
	// beginning source position
	Position int
	// token type name
	TokenTypeName string
	// token type
	TokenType token.TokenType
	// token text
	Text string
	// token value
	Value interface{}
}

func (te *TokenEvent) Log() {
	format1 := ">>> (%15s) line=%03d, pos=%2d, text=\"%s\"\n"
	fmt.Printf(format1, te.TokenTypeName, te.LineNumber, te.Position, te.Text)

	// format2 := ">>>       value=%s"

	// if te.Value != nil {
	// 	tokenValue := te.Value
	// 	if te.TokenType == token.STRING {
	// 		tokenValue = "\"" + te.Value.(string) + "\"\n"
	// 	}
	// 	fmt.Printf(format2, tokenValue)
	// }
}

type SyntaxErrorEvent struct {
	// 错误出现的行号
	LineNumber int
	// 错误列号
	Position int
	// 错误token
	Text string
	// 错误描述
	ErrorMessage string
}

func (se *SyntaxErrorEvent) Log() {
	format := "第%d行第%d列出现错误, 错误字符为: %s\n, 错误描述为: %s\n"
	fmt.Printf(format, se.LineNumber, se.Position, se.Text, se.ErrorMessage)
}
