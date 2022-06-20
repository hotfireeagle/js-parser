package parser

import (
	"jsj/frontend/token"
	"jsj/message"
	"os"
)

// 最大语法错误数
// 若超出直接中止解析过程
var MAX_ERRORS = 10

type ParserErrorHandler struct {
	// 当前已检测到的错误数
	errorCount int
}

func ParserErrorHandlerConstructor() *ParserErrorHandler {
	return &ParserErrorHandler{
		errorCount: 0,
	}
}

// 发布词法错误事件
func (peh *ParserErrorHandler) Flag(
	tokenObj token.Token,
	errorCode SyntaxErrorCode,
	parserIns *Parser,
) {
	syntaxErrorEvent := SyntaxErrorEvent{
		LineNumber:   tokenObj.GetLineNumber(),
		Position:     tokenObj.GetPosition(),
		Text:         tokenObj.GetText(),
		ErrorMessage: errorCode.ToString(),
	}
	messageObj := message.MessageConstructor(message.SYNTAX_ERROR, syntaxErrorEvent)
	parserIns.ParserSendMessage(messageObj)

	if peh.errorCount > MAX_ERRORS {
		errorCodeIns := SyntaxErrorCodeConstructor(TOO_MANY_ERRORS, "too many errors")
		peh.AbortTranslation(errorCodeIns, parserIns)
	}
}

// 中止解析过程
func (peh *ParserErrorHandler) AbortTranslation(errorCode *SyntaxErrorCode, parserIns *Parser) {
	syntaxErrorEvent := SyntaxErrorEvent{
		LineNumber:   0,
		Position:     0,
		Text:         "",
		ErrorMessage: errorCode.ToString(),
	}
	messageObj := message.MessageConstructor(message.SYNTAX_ERROR, syntaxErrorEvent)
	parserIns.ParserSendMessage(messageObj)
	os.Exit(1)
}

func (peh *ParserErrorHandler) GetErrorCount() int {
	return peh.errorCount
}
