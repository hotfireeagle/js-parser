package token

import (
	"jsj/frontend/source"
)

type DeterminedMinusToken struct {
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

func DeterminedMinusTokenConstructor(s *source.Source) *DeterminedMinusToken {
	tokenObj := DeterminedMinusToken{
		source:   s,
		lineNum:  s.GetLineNum(),
		position: s.GetPosition(),
	}

	tokenObj.Extract()

	return &tokenObj
}

// 以-开头的token，可能是[-, -=, --]
// 这里确定到底是哪个
func (dpt *DeterminedMinusToken) Extract() {
	nextC := dpt.NextChar()

	var tokenT TokenType
	var t string
	needSkip := false

	if nextC == '=' {
		t = "-="
		tokenT = SUBTRACT_ASSIGN
		needSkip = true
	} else if nextC == '-' {
		t = "--"
		tokenT = DECREMENT
		needSkip = true
	} else {
		t = "-"
		tokenT = MINUS
	}

	dpt.tokenType = tokenT
	dpt.text = t

	if needSkip {
		dpt.NextChar()
	}
}

// return the current character from the source
// call the source's CurrentChar() method
func (t *DeterminedMinusToken) CurrentChar() byte {
	return t.source.CurrentChar()
}

// return the next character from the source
func (t *DeterminedMinusToken) NextChar() byte {
	return t.source.NextChar()
}

// return the source character *following* the current character without consuming the current character
func (t *DeterminedMinusToken) PeekChar() byte {
	return t.source.PeekChar()
}

func (t *DeterminedMinusToken) GetTokenType() TokenType {
	return t.tokenType
}

func (t *DeterminedMinusToken) GetLineNumber() int {
	return t.lineNum
}

func (t *DeterminedMinusToken) GetPosition() int {
	return t.position
}

func (t *DeterminedMinusToken) GetText() string {
	return t.text
}

func (t *DeterminedMinusToken) GetValue() TokenValue {
	return t.value
}

func (t *DeterminedMinusToken) GetSource() *source.Source {
	return t.source
}
