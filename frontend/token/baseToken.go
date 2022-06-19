package token

import (
	"jsj/frontend/source"
)

type TokenValue interface{}

type Token struct {
	// javascript token type
	tokenType TokenType

	// the token text
	text string

	// the token value
	value TokenValue

	// the token which belong to file
	source source.Source

	// the token in source file's line number
	lineNum int

	// position of the token first character
	position int
}

func TokenConstructor(s source.Source) Token {
	tokenObj := Token{
		source:   s,
		lineNum:  s.GetLineNum(),
		position: s.GetPosition(),
	}

	tokenObj.Extract()

	return tokenObj
}

// extract : 提取
// 默认的提取方法，只适用于从source文件中提取出一个character的token
// 不同的token可覆盖该实现方法
func (t *Token) Extract() {
	t.text = string(t.CurrentChar())
	t.value = nil
	t.NextChar() // sideEffect: consume current character
}

// return the current character from the source
// call the source's CurrentChar() method
func (t *Token) CurrentChar() byte {
	return t.source.CurrentChar()
}

// return the next character from the source
func (t *Token) NextChar() byte {
	return t.source.NextChar()
}

// return the source character *following* the current character without consuming the current character
func (t *Token) PeekChar() byte {
	return t.source.PeekChar()
}

func (t *Token) GetTokenType() TokenType {
	return t.tokenType
}

func (t *Token) GetLineNumber() int {
	return t.lineNum
}
