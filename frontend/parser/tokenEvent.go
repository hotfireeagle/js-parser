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
	// token type
	TokenTypeName token.TokenType
	// token text
	Text string
	// token value
	Value interface{}
}

func (te *TokenEvent) Log() {
	// TODO: token type print
	format1 := ">>> %s line=%03d, pos=%2d, text=\"%s\"\n"
	fmt.Printf(format1, te.TokenTypeName, te.LineNumber, te.Position, te.Text)

	format2 := ">>>       value=%s"

	if te.Value != nil {
		tokenValue := te.Value
		if te.TokenTypeName == token.STRING {
			tokenValue = "\"" + te.Value.(string) + "\"\n"
		}
		fmt.Printf(format2, tokenValue)
	}
}
