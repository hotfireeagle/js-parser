// a top down parser
// parser also can be bottom up
// scanner return token, and parser's input is token, output is ast
package parser

import (
	"jsj/frontend/scanner"
	"jsj/frontend/source"
	"jsj/frontend/token"
	"jsj/intermediate"
	"jsj/message"
	"time"
)

var messageHandler = message.MessageHandlerConstructor()
var errHandler = ParserErrorHandlerConstructor()

var sysTab *intermediate.SymTab
var symTabStack = intermediate.SymTabStackConstructor()

type Parser struct {
	icode   *intermediate.ICode
	scanner *scanner.Scanner
}

func ParserConstructor(s *source.Source) Parser {
	scannerInstance := scanner.ScannerConstructor(s)
	return Parser{
		scanner: scannerInstance,
	}
}

// parse a javascript source file and generate the symbol table
func (p *Parser) Parse() {
	var tokenInstance token.Token
	startTime := time.Now().UnixMilli()

	// 不断调用NextToken()，直到解析到了文件末尾为止
	tokenInstance = p.NextToken()
	for tokenInstance.GetTokenType() != token.EOF {
		if tokenInstance.GetTokenType() == token.ERROR {
			errHandler.Flag(tokenInstance, tokenInstance.GetValue().(SyntaxErrorCode), p)
		} else if tokenInstance.GetTokenType() == token.IDENTIFIER {
			// 对于javascript来说，关键字是大小写敏感的，所以不做小写的处理
			identifierName := tokenInstance.GetText()
			entry := symTabStack.LookUp(identifierName)
			if entry == nil {
				entry = symTabStack.EnterLocal(identifierName)
			}
			entry.AppendLineNumber(tokenInstance.GetLineNumber())
			tokenInstance = p.NextToken()
		} else {
			tokenLog := &TokenEvent{
				LineNumber:    tokenInstance.GetLineNumber(),
				Position:      tokenInstance.GetPosition(),
				TokenTypeName: token.GetTokenTypeName(tokenInstance.GetTokenType()),
				TokenType:     tokenInstance.GetTokenType(),
				Text:          tokenInstance.GetText(),
				Value:         tokenInstance.GetValue(),
			}
			messageObj := message.MessageConstructor(message.TOKEN, tokenLog)
			p.ParserSendMessage(messageObj)
			tokenInstance = p.NextToken()
		}
	}

	endTime := time.Now().UnixMilli()

	var elapsedTime float64 = (float64(endTime) - float64(startTime)) / 1000

	eventlog := &message.ParserSummaryEvent{
		LineNum:     tokenInstance.GetLineNumber(),
		ErrorCount:  p.GetErrorCount(),
		ElapsedTime: elapsedTime,
	}

	message := message.MessageConstructor(message.PARSER_SUMMARY, eventlog)

	p.ParserSendMessage(message)
}

// return the number of syntax errors found by the parser
func (p *Parser) GetErrorCount() int {
	return 0
}

func (p *Parser) CurrentToken() token.Token {
	// TODO:
	return nil
}

func (p *Parser) NextToken() token.Token {
	return p.scanner.NextToken()
}

func (p *Parser) AddMessageListener(listener message.MessageListener) {
	messageHandler.AddListener(listener)
}

func (p *Parser) ParserSendMessage(messageObj message.Message) {
	messageHandler.SendMessage(messageObj)
}

func (p *Parser) GetICode() *intermediate.ICode {
	return p.icode
}

func (p *Parser) GetSymTab() *intermediate.SymTab {
	return sysTab
}

func (p *Parser) GetSymTabStack() *intermediate.SymTabStack {
	return symTabStack
}
