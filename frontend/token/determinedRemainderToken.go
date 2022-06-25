package token

import (
	"jsj/frontend/source"
)

type DeterminedRemainderToken struct {
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

func DeterminedRemainderTokenConstructor(s *source.Source) *DeterminedRemainderToken {
	tokenObj := DeterminedRemainderToken{
		source:   s,
		lineNum:  s.GetLineNum(),
		position: s.GetPosition(),
	}

	tokenObj.Extract()

	return &tokenObj
}

// 以%开头的token，可能是[%, %=]
// 这里确定到底是哪个
func (dpt *DeterminedRemainderToken) Extract() {
	nextC := dpt.NextChar()

	var tokenT TokenType
	var t string
	needSkip := false

	if nextC == '=' {
		t = "%="
		tokenT = REMAINDER_ASSIGN
		needSkip = true
	} else {
		t = "%"
		tokenT = REMAINDER
	}

	dpt.tokenType = tokenT
	dpt.text = t

	if needSkip {
		dpt.NextChar()
	}
}

// return the current character from the source
// call the source's CurrentChar() method
func (t *DeterminedRemainderToken) CurrentChar() byte {
	return t.source.CurrentChar()
}

// return the next character from the source
func (t *DeterminedRemainderToken) NextChar() byte {
	return t.source.NextChar()
}

// return the source character *following* the current character without consuming the current character
func (t *DeterminedRemainderToken) PeekChar() byte {
	return t.source.PeekChar()
}

func (t *DeterminedRemainderToken) GetTokenType() TokenType {
	return t.tokenType
}

func (t *DeterminedRemainderToken) GetLineNumber() int {
	return t.lineNum
}

func (t *DeterminedRemainderToken) GetPosition() int {
	return t.position
}

func (t *DeterminedRemainderToken) GetText() string {
	return t.text
}

func (t *DeterminedRemainderToken) GetValue() TokenValue {
	return t.value
}

func (t *DeterminedRemainderToken) GetSource() *source.Source {
	return t.source
}
