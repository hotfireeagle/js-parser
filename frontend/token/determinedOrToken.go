package token

import (
	"jsj/frontend/source"
)

type DeterminedOrToken struct {
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

func DeterminedOrTokenConstructor(s *source.Source) *DeterminedOrToken {
	tokenObj := DeterminedOrToken{
		source:   s,
		lineNum:  s.GetLineNum(),
		position: s.GetPosition(),
	}

	tokenObj.Extract()

	return &tokenObj
}

// 以|开头的token，可能是[|, |=, ||]
// 这里确定到底是哪个
func (dpt *DeterminedOrToken) Extract() {
	nextC := dpt.NextChar()

	var tokenT TokenType
	var t string
	needSkip := false

	if nextC == '|' {
		t = "||"
		tokenT = LOGICAL_OR
		needSkip = true
	} else if nextC == '=' {
		t = "|="
		tokenT = OR_ASSIGN
		needSkip = true
	} else {
		t = "|"
		tokenT = OR
	}

	dpt.tokenType = tokenT
	dpt.text = t

	if needSkip {
		dpt.NextChar()
	}
}

// return the current character from the source
// call the source's CurrentChar() method
func (t *DeterminedOrToken) CurrentChar() byte {
	return t.source.CurrentChar()
}

// return the next character from the source
func (t *DeterminedOrToken) NextChar() byte {
	return t.source.NextChar()
}

// return the source character *following* the current character without consuming the current character
func (t *DeterminedOrToken) PeekChar() byte {
	return t.source.PeekChar()
}

func (t *DeterminedOrToken) GetTokenType() TokenType {
	return t.tokenType
}

func (t *DeterminedOrToken) GetLineNumber() int {
	return t.lineNum
}

func (t *DeterminedOrToken) GetPosition() int {
	return t.position
}

func (t *DeterminedOrToken) GetText() string {
	return t.text
}

func (t *DeterminedOrToken) GetValue() TokenValue {
	return t.value
}

func (t *DeterminedOrToken) GetSource() *source.Source {
	return t.source
}
