package token

import (
	"jsj/frontend/source"
)

type DeterminedLessToken struct {
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

func DeterminedLessTokenConstructor(s *source.Source) *DeterminedLessToken {
	tokenObj := DeterminedLessToken{
		source:   s,
		lineNum:  s.GetLineNum(),
		position: s.GetPosition(),
	}

	tokenObj.Extract()

	return &tokenObj
}

// 以<开头的token，可能是[<, <<, <<=]
// 这里确定到底是哪个
func (dpt *DeterminedLessToken) Extract() {
	nextC := dpt.NextChar()

	var tokenT TokenType
	var t string
	needSkip := false

	if nextC == '<' {
		nc := dpt.NextChar()
		if nc == '=' {
			t = "<<="
			tokenT = SHIFT_LEFT_ASSIGN
			needSkip = true
		} else {
			t = "<<"
			tokenT = SHIFT_LEFT
		}
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
func (t *DeterminedLessToken) CurrentChar() byte {
	return t.source.CurrentChar()
}

// return the next character from the source
func (t *DeterminedLessToken) NextChar() byte {
	return t.source.NextChar()
}

// return the source character *following* the current character without consuming the current character
func (t *DeterminedLessToken) PeekChar() byte {
	return t.source.PeekChar()
}

func (t *DeterminedLessToken) GetTokenType() TokenType {
	return t.tokenType
}

func (t *DeterminedLessToken) GetLineNumber() int {
	return t.lineNum
}

func (t *DeterminedLessToken) GetPosition() int {
	return t.position
}

func (t *DeterminedLessToken) GetText() string {
	return t.text
}

func (t *DeterminedLessToken) GetValue() TokenValue {
	return t.value
}

func (t *DeterminedLessToken) GetSource() *source.Source {
	return t.source
}
