package token

import (
	"jsj/frontend/source"
	"strconv"
	"strings"
	"unicode"
)

type NumberToken struct {
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

	// 是否是正数
	isPositive bool
}

func NumberTokenConstructor(s *source.Source, isPositive bool) *NumberToken {
	tokenObj := NumberToken{
		source:     s,
		lineNum:    s.GetLineNum(),
		position:   s.GetPosition(),
		isPositive: isPositive,
	}

	tokenObj.Extract()

	return &tokenObj
}

func (dpt *NumberToken) Extract() {
	currentChar := dpt.CurrentChar()

	dotHasExists := false

	var str strings.Builder

	for (dotHasExists && unicode.IsNumber(rune(currentChar))) || (!dotHasExists && CheckIsNumberBeginCharacter(currentChar)) {
		str.WriteByte(currentChar)

		if currentChar == '.' {
			dotHasExists = true
		}

		currentChar = dpt.NextChar()
	}

	numStr := str.String()

	if !dpt.isPositive {
		numStr = "-" + numStr
	}

	dpt.tokenType = NUMBER
	dpt.text = numStr

	if dotHasExists {
		// f64
		f64v, err := strconv.ParseFloat(numStr, 64)
		if err != nil {
			// TODO:
			panic(err)
		}
		dpt.value = f64v
	} else {
		// int64
		i64v, err := strconv.ParseInt(numStr, 10, 64)
		if err != nil {
			// TODO:
			panic(err)
		}
		dpt.value = i64v
	}
}

// return the current character from the source
// call the source's CurrentChar() method
func (t *NumberToken) CurrentChar() byte {
	return t.source.CurrentChar()
}

// return the next character from the source
func (t *NumberToken) NextChar() byte {
	return t.source.NextChar()
}

// return the source character *following* the current character without consuming the current character
func (t *NumberToken) PeekChar() byte {
	return t.source.PeekChar()
}

func (t *NumberToken) GetTokenType() TokenType {
	return t.tokenType
}

func (t *NumberToken) GetLineNumber() int {
	return t.lineNum
}

func (t *NumberToken) GetPosition() int {
	return t.position
}

func (t *NumberToken) GetText() string {
	return t.text
}

func (t *NumberToken) GetValue() TokenValue {
	return t.value
}

func (t *NumberToken) GetSource() *source.Source {
	return t.source
}
