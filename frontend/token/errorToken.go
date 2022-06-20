package token

import (
	"jsj/frontend/source"
)

type ErrorToken struct {
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

func ErrorTokenConstructor(s *source.Source) *ErrorToken {
	tokenObj := ErrorToken{
		source:    s,
		lineNum:   s.GetLineNum(),
		position:  s.GetPosition(),
		tokenType: ERROR,
	}

	tokenObj.Extract()

	return &tokenObj
}

// 不消费任何字符
func (t *ErrorToken) Extract() {
	// TODO:
	panic("ErrorToken")
}

// return the current character from the source
// call the source's CurrentChar() method
func (t *ErrorToken) CurrentChar() byte {
	return t.source.CurrentChar()
}

// return the next character from the source
func (t *ErrorToken) NextChar() byte {
	return t.source.NextChar()
}

// return the source character *following* the current character without consuming the current character
func (t *ErrorToken) PeekChar() byte {
	return t.source.PeekChar()
}

func (t *ErrorToken) GetTokenType() TokenType {
	return t.tokenType
}

func (t *ErrorToken) GetLineNumber() int {
	return t.lineNum
}

func (t *ErrorToken) GetPosition() int {
	return t.position
}

func (t *ErrorToken) GetText() string {
	return t.text
}

func (t *ErrorToken) GetValue() TokenValue {
	return t.value
}

func (t *ErrorToken) GetSource() *source.Source {
	return t.source
}
