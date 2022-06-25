package token

import (
	"jsj/frontend/source"
	"jsj/utils"
	"strings"
)

type DeterminedSlashToken struct {
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

func DeterminedSlashTokenConstructor(s *source.Source) *DeterminedSlashToken {
	tokenObj := DeterminedSlashToken{
		source:   s,
		lineNum:  s.GetLineNum(),
		position: s.GetPosition(),
	}

	tokenObj.Extract()

	return &tokenObj
}

// 以/开头的token，可能是[/, /=], 或者是注释
func (dpt *DeterminedSlashToken) Extract() {
	nextC := dpt.NextChar()

	var tokenT TokenType
	var t string
	needSkip := false

	if nextC == '=' {
		t = "/="
		tokenT = QUOTIENT_ASSIGN
		needSkip = true
	} else if nextC == '/' {
		var value strings.Builder
		nc := dpt.NextChar()

		for nc != utils.EOL {
			if value.Len() == 0 && nc == ' ' {
				nc = dpt.NextChar()
				continue
			}
			value.WriteByte(nc)
			nc = dpt.NextChar()
		}

		tokenT = COMMENT
		dpt.value = value.String()
		needSkip = true
	} else {
		t = "/"
		tokenT = SLASH
	}

	dpt.tokenType = tokenT
	dpt.text = t

	if needSkip {
		dpt.NextChar()
	}
}

// return the current character from the source
// call the source's CurrentChar() method
func (t *DeterminedSlashToken) CurrentChar() byte {
	return t.source.CurrentChar()
}

// return the next character from the source
func (t *DeterminedSlashToken) NextChar() byte {
	return t.source.NextChar()
}

// return the source character *following* the current character without consuming the current character
func (t *DeterminedSlashToken) PeekChar() byte {
	return t.source.PeekChar()
}

func (t *DeterminedSlashToken) GetTokenType() TokenType {
	return t.tokenType
}

func (t *DeterminedSlashToken) GetLineNumber() int {
	return t.lineNum
}

func (t *DeterminedSlashToken) GetPosition() int {
	return t.position
}

func (t *DeterminedSlashToken) GetText() string {
	return t.text
}

func (t *DeterminedSlashToken) GetValue() TokenValue {
	return t.value
}

func (t *DeterminedSlashToken) GetSource() *source.Source {
	return t.source
}
