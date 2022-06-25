package token

import (
	"jsj/frontend/source"
)

type SingleCharacterToken struct {
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

func SingleCharacterTokenConstructor(s *source.Source) *SingleCharacterToken {
	tokenObj := SingleCharacterToken{
		source:   s,
		lineNum:  s.GetLineNum(),
		position: s.GetPosition(),
	}

	tokenObj.Extract()

	return &tokenObj
}

// 单一字符token还可进行细化
// 在这里会具体细化它到底对应哪个token类型
func (dqt *SingleCharacterToken) Extract() {
	currentChar := dqt.CurrentChar()

	var tokenT TokenType

	for k, v := range SingleToken {
		if currentChar == k {
			tokenT = v
			break
		}
	}

	dqt.tokenType = tokenT
	dqt.text = string(currentChar)

	dqt.NextChar() // skip current
}

// return the current character from the source
// call the source's CurrentChar() method
func (t *SingleCharacterToken) CurrentChar() byte {
	return t.source.CurrentChar()
}

// return the next character from the source
func (t *SingleCharacterToken) NextChar() byte {
	return t.source.NextChar()
}

// return the source character *following* the current character without consuming the current character
func (t *SingleCharacterToken) PeekChar() byte {
	return t.source.PeekChar()
}

func (t *SingleCharacterToken) GetTokenType() TokenType {
	return t.tokenType
}

func (t *SingleCharacterToken) GetLineNumber() int {
	return t.lineNum
}

func (t *SingleCharacterToken) GetPosition() int {
	return t.position
}

func (t *SingleCharacterToken) GetText() string {
	return t.text
}

func (t *SingleCharacterToken) GetValue() TokenValue {
	return t.value
}

func (t *SingleCharacterToken) GetSource() *source.Source {
	return t.source
}
