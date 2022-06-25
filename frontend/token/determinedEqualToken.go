package token

import (
	"jsj/frontend/source"
)

type DeterminedEqualToken struct {
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

func DeterminedEqualTokenConstructor(s *source.Source) *DeterminedEqualToken {
	tokenObj := DeterminedEqualToken{
		source:   s,
		lineNum:  s.GetLineNum(),
		position: s.GetPosition(),
	}

	tokenObj.Extract()

	return &tokenObj
}

// 以=开头的token，可能是[=, ==, ===]
// 这里确定到底是哪个
func (dpt *DeterminedEqualToken) Extract() {
	nextC := dpt.NextChar()

	var tokenT TokenType
	var t string
	needSkip := false

	if nextC == '=' {
		nextC = dpt.NextChar()
		if nextC == '=' {
			t = "==="
			tokenT = STRICT_EQUAL
			needSkip = true
		} else {
			t = "=="
			tokenT = EQUAL
		}
	} else {
		t = "="
		tokenT = ASSIGN
	}

	dpt.tokenType = tokenT
	dpt.text = t

	if needSkip {
		dpt.NextChar()
	}
}

// return the current character from the source
// call the source's CurrentChar() method
func (t *DeterminedEqualToken) CurrentChar() byte {
	return t.source.CurrentChar()
}

// return the next character from the source
func (t *DeterminedEqualToken) NextChar() byte {
	return t.source.NextChar()
}

// return the source character *following* the current character without consuming the current character
func (t *DeterminedEqualToken) PeekChar() byte {
	return t.source.PeekChar()
}

func (t *DeterminedEqualToken) GetTokenType() TokenType {
	return t.tokenType
}

func (t *DeterminedEqualToken) GetLineNumber() int {
	return t.lineNum
}

func (t *DeterminedEqualToken) GetPosition() int {
	return t.position
}

func (t *DeterminedEqualToken) GetText() string {
	return t.text
}

func (t *DeterminedEqualToken) GetValue() TokenValue {
	return t.value
}

func (t *DeterminedEqualToken) GetSource() *source.Source {
	return t.source
}
