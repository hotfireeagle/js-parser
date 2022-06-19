// a top down parser
// parser also can be bottom up
// scanner return token, and parser's input is token, output is ast
package parser

import (
	"jsj/frontend/scanner"
	"jsj/frontend/token"
	"jsj/message"
	"time"
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

// parse a javascript source file and generate the symbol table
func (p *Parser) Parse() {
	var tokenInstance token.Token
	startTime := time.Now().UnixMilli()

	// 不断调用NextToken()，直到解析到了文件末尾为止
	for tokenInstance = p.NextToken(); tokenInstance.GetTokenType() != token.EOF; {

	}

	endTime := time.Now().UnixMilli()

	var elapsedTime float64 = (float64(endTime) - float64(startTime)) / 1000

	eventlog := message.ParserSummaryEvent{
		LineNum:     tokenInstance.GetLineNumber(),
		ErrorCount:  p.GetErrorCount(),
		ElapsedTime: elapsedTime,
	}

	message := message.MessageConstructor(message.PARSER_SUMMARY, eventlog)

	messageHandler.SendMessage(message)
}

// return the number of syntax errors found by the parser
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
