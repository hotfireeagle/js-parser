package parser

import (
	"jsj/frontend/scanner"
	"jsj/frontend/token"
	"jsj/message"
)

type Parser struct {
	scanner scanner.Scanner
	message.MessageHandler
}

func ParserConstructor(s scanner.Scanner) Parser {
	return Parser{
		scanner:        s,
		MessageHandler: message.MessageHandlerConstructor(),
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
