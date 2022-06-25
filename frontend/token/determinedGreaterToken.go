package token

import (
	"jsj/frontend/source"
)

type DeterminedGreaterToken struct {
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

func DeterminedGreaterTokenConstructor(s *source.Source) *DeterminedGreaterToken {
	tokenObj := DeterminedGreaterToken{
		source:   s,
		lineNum:  s.GetLineNum(),
		position: s.GetPosition(),
	}

	tokenObj.Extract()

	return &tokenObj
}

// 以>开头的token，可能是[>, >>, >>>, >>=, >>>=, >=]
// 这里确定到底是哪个
func (dpt *DeterminedGreaterToken) Extract() {
	nextC := dpt.NextChar()

	var tokenT TokenType
	var t string
	needSkip := false

	if nextC == '>' {
		nc := dpt.NextChar()
		if nc == '>' {
			nc2 := dpt.NextChar()
			if nc2 == '=' {
				t = ">>>="
				tokenT = UNSIGNED_SHIFT_RIGHT_ASSIGN
				needSkip = true
			} else {
				t = ">>>"
				tokenT = UNSIGNED_SHIFT_RIGHT
			}
		} else if nc == '=' {
			t = ">>="
			tokenT = SHIFT_RIGHT_ASSIGN
			needSkip = true
		} else {
			t = ">>"
			tokenT = SHIFT_RIGHT
		}
	} else if nextC == '=' {
		t = ">="
		tokenT = GREATER_OR_EQUAL
		needSkip = true
	} else {
		t = ">"
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
func (t *DeterminedGreaterToken) CurrentChar() byte {
	return t.source.CurrentChar()
}

// return the next character from the source
func (t *DeterminedGreaterToken) NextChar() byte {
	return t.source.NextChar()
}

// return the source character *following* the current character without consuming the current character
func (t *DeterminedGreaterToken) PeekChar() byte {
	return t.source.PeekChar()
}

func (t *DeterminedGreaterToken) GetTokenType() TokenType {
	return t.tokenType
}

func (t *DeterminedGreaterToken) GetLineNumber() int {
	return t.lineNum
}

func (t *DeterminedGreaterToken) GetPosition() int {
	return t.position
}

func (t *DeterminedGreaterToken) GetText() string {
	return t.text
}

func (t *DeterminedGreaterToken) GetValue() TokenValue {
	return t.value
}

func (t *DeterminedGreaterToken) GetSource() *source.Source {
	return t.source
}
