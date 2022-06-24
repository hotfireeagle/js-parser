package token

import (
	"jsj/frontend/source"
	"strings"
)

type DoubleQuoteToken struct {
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

func DoubleQuoteTokenConstructor(s *source.Source) *DoubleQuoteToken {
	tokenObj := DoubleQuoteToken{
		source:   s,
		lineNum:  s.GetLineNum(),
		position: s.GetPosition(),
	}

	tokenObj.Extract()

	return &tokenObj
}

// 提取出token
// TODO: error handler, string的后面可以是什么才合法
func (dqt *DoubleQuoteToken) Extract() {
	currentChar := dqt.NextChar()
	var textBuilder strings.Builder

	// TODO: 字符串可换行，兼容该场景
	// TODO: 考虑""未正确闭合的错误场景
	for currentChar != '"' {
		textBuilder.WriteByte(currentChar)
		currentChar = dqt.NextChar()
	}

	// 跳过结尾的 "
	dqt.NextChar()

	dqt.tokenType = STRING
	dqt.text = textBuilder.String()
	dqt.value = dqt.text
}

// return the current character from the source
// call the source's CurrentChar() method
func (t *DoubleQuoteToken) CurrentChar() byte {
	return t.source.CurrentChar()
}

// return the next character from the source
func (t *DoubleQuoteToken) NextChar() byte {
	return t.source.NextChar()
}

// return the source character *following* the current character without consuming the current character
func (t *DoubleQuoteToken) PeekChar() byte {
	return t.source.PeekChar()
}

func (t *DoubleQuoteToken) GetTokenType() TokenType {
	return t.tokenType
}

func (t *DoubleQuoteToken) GetLineNumber() int {
	return t.lineNum
}

func (t *DoubleQuoteToken) GetPosition() int {
	return t.position
}

func (t *DoubleQuoteToken) GetText() string {
	return t.text
}

func (t *DoubleQuoteToken) GetValue() TokenValue {
	return t.value
}

func (t *DoubleQuoteToken) GetSource() *source.Source {
	return t.source
}
