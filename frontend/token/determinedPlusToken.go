package token

import (
	"jsj/frontend/source"
)

type DeterminedPlusToken struct {
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

func DeterminedPlusTokenConstructor(s *source.Source) *DeterminedPlusToken {
	tokenObj := DeterminedPlusToken{
		source:   s,
		lineNum:  s.GetLineNum(),
		position: s.GetPosition(),
	}

	tokenObj.Extract()

	return &tokenObj
}

// 以+开头的token，可能是[+, +=, ++]
// 这里确定到底是哪个
func (dpt *DeterminedPlusToken) Extract() {
	nextC := dpt.NextChar()

	var tokenT TokenType
	var t string
	needSkip := false

	if nextC == '=' {
		t = "+="
		tokenT = ADD_ASSIGN
		needSkip = true
	} else if nextC == '+' {
		t = "++"
		tokenT = INCREMENT
		needSkip = true
	} else {
		t = "+"
		tokenT = PLUS
	}

	dpt.tokenType = tokenT
	dpt.text = t

	if needSkip {
		dpt.NextChar()
	}
}

// return the current character from the source
// call the source's CurrentChar() method
func (t *DeterminedPlusToken) CurrentChar() byte {
	return t.source.CurrentChar()
}

// return the next character from the source
func (t *DeterminedPlusToken) NextChar() byte {
	return t.source.NextChar()
}

// return the source character *following* the current character without consuming the current character
func (t *DeterminedPlusToken) PeekChar() byte {
	return t.source.PeekChar()
}

func (t *DeterminedPlusToken) GetTokenType() TokenType {
	return t.tokenType
}

func (t *DeterminedPlusToken) GetLineNumber() int {
	return t.lineNum
}

func (t *DeterminedPlusToken) GetPosition() int {
	return t.position
}

func (t *DeterminedPlusToken) GetText() string {
	return t.text
}

func (t *DeterminedPlusToken) GetValue() TokenValue {
	return t.value
}

func (t *DeterminedPlusToken) GetSource() *source.Source {
	return t.source
}
