package token

import (
	"jsj/frontend/source"
)

type BaseToken struct {
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

func BaseTokenConstructor(s *source.Source) *BaseToken {
	tokenObj := BaseToken{
		source:   s,
		lineNum:  s.GetLineNum(),
		position: s.GetPosition(),
	}

	tokenObj.Extract()

	return &tokenObj
}

// extract : 提取
// 默认的提取方法，只适用于从source文件中提取出一个character的token
// 不同的token可覆盖该实现方法
func (t *BaseToken) Extract() {
	t.text = string(t.CurrentChar())
	t.value = nil
	t.NextChar() // sideEffect: consume current character
}

// return the current character from the source
// call the source's CurrentChar() method
func (t *BaseToken) CurrentChar() byte {
	return t.source.CurrentChar()
}

// return the next character from the source
func (t *BaseToken) NextChar() byte {
	return t.source.NextChar()
}

// return the source character *following* the current character without consuming the current character
func (t *BaseToken) PeekChar() byte {
	return t.source.PeekChar()
}

func (t *BaseToken) GetTokenType() TokenType {
	return t.tokenType
}

func (t *BaseToken) GetLineNumber() int {
	return t.lineNum
}

func (t *BaseToken) GetPosition() int {
	return t.position
}

func (t *BaseToken) GetText() string {
	return t.text
}

func (t *BaseToken) GetValue() TokenValue {
	return t.value
}

func (t *BaseToken) GetSource() *source.Source {
	return t.source
}
