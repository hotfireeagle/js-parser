package parser

import (
	"jsj/frontend/scanner"
	"jsj/frontend/token"
	"jsj/message"
)

var messageHandler = message.MessageHandlerConstructor()

type Parser struct {
	scanner scanner.Scanner
}

func ParserConstructor(s scanner.Scanner) Parser {
	return Parser{
		scanner: s,
	}
}

func (p *Parser) Parse() {

}

func (p *Parser) GetErrorCount() int {
	return 0
}

func (p *Parser) CurrentToken() token.Token {
	// TODO:
	return token.Token{}
}

func (p *Parser) NextToken() token.Token {
	return p.scanner.NextToken()
}
