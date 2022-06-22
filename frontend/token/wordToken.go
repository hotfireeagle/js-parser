package token

import (
	"jsj/frontend/source"
	"jsj/utils"
	"strings"
)

type WordToken struct {
	// javascript token type
	tokenType TokenType

	// the token text
	text string

	// the token value
	value TokenValue

	// the token which belong to file
	source *source.Source

	// the token in source file's line number
	lineNum int

	// position of the token first character
	position int
}

func WordTokenConstructor(s *source.Source) *WordToken {
	tokenObj := WordToken{
		source:   s,
		lineNum:  s.GetLineNum(),
		position: s.GetPosition(),
	}

	tokenObj.Extract()

	return &tokenObj
}

// 提取token
func (wt *WordToken) Extract() {
	var textBuffer strings.Builder

	currentChar := wt.CurrentChar()

	for utils.CheckIsJSWordSuffix(currentChar) {
		textBuffer.WriteByte(currentChar)
		currentChar = wt.NextChar()
	}

	text := textBuffer.String()

	var tokenType TokenType

	if tt, ok := KeyWordMap[text]; ok {
		// 表明这个提取出来的word属于保留关键字
		tokenType = tt
	} else {
		// 表明这个提取出来的word是一个变量标识符
		tokenType = IDENTIFIER
	}

	wt.tokenType = tokenType
	wt.text = text
}

// return the current character from the source
// call the source's CurrentChar() method
func (t *WordToken) CurrentChar() byte {
	return t.source.CurrentChar()
}

// return the next character from the source
func (t *WordToken) NextChar() byte {
	return t.source.NextChar()
}

// return the source character *following* the current character without consuming the current character
func (t *WordToken) PeekChar() byte {
	return t.source.PeekChar()
}

func (t *WordToken) GetTokenType() TokenType {
	return t.tokenType
}

func (t *WordToken) GetLineNumber() int {
	return t.lineNum
}

func (t *WordToken) GetPosition() int {
	return t.position
}

func (t *WordToken) GetText() string {
	return t.text
}

func (t *WordToken) GetValue() TokenValue {
	return t.value
}
