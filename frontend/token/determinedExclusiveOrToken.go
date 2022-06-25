package token

import (
	"jsj/frontend/source"
)

type DeterminedExclusiveOrToken struct {
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

func DeterminedExclusiveOrTokenConstructor(s *source.Source) *DeterminedExclusiveOrToken {
	tokenObj := DeterminedExclusiveOrToken{
		source:   s,
		lineNum:  s.GetLineNum(),
		position: s.GetPosition(),
	}

	tokenObj.Extract()

	return &tokenObj
}

// 以^开头的token，可能是[^, ^=]
// 这里确定到底是哪个
func (dpt *DeterminedExclusiveOrToken) Extract() {
	nextC := dpt.NextChar()

	var tokenT TokenType
	var t string
	needSkip := false

	if nextC == '=' {
		t = "^="
		tokenT = EXCLUSIVE_OR_ASSIGN
		needSkip = true
	} else {
		t = "^"
		tokenT = EXCLUSIVE_OR
	}

	dpt.tokenType = tokenT
	dpt.text = t

	if needSkip {
		dpt.NextChar()
	}
}

// return the current character from the source
// call the source's CurrentChar() method
func (t *DeterminedExclusiveOrToken) CurrentChar() byte {
	return t.source.CurrentChar()
}

// return the next character from the source
func (t *DeterminedExclusiveOrToken) NextChar() byte {
	return t.source.NextChar()
}

// return the source character *following* the current character without consuming the current character
func (t *DeterminedExclusiveOrToken) PeekChar() byte {
	return t.source.PeekChar()
}

func (t *DeterminedExclusiveOrToken) GetTokenType() TokenType {
	return t.tokenType
}

func (t *DeterminedExclusiveOrToken) GetLineNumber() int {
	return t.lineNum
}

func (t *DeterminedExclusiveOrToken) GetPosition() int {
	return t.position
}

func (t *DeterminedExclusiveOrToken) GetText() string {
	return t.text
}

func (t *DeterminedExclusiveOrToken) GetValue() TokenValue {
	return t.value
}

func (t *DeterminedExclusiveOrToken) GetSource() *source.Source {
	return t.source
}
