package token

import "jsj/frontend/source"

type EofToken struct {
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

func EofTokenConstructor(s *source.Source) *EofToken {
	tokenObj := EofToken{
		source:    s,
		lineNum:   s.GetLineNum(),
		position:  s.GetPosition(),
		tokenType: EOF,
	}

	tokenObj.Extract()

	return &tokenObj
}

// because EOF means the end of file
// so there are nothing to extract
func (t *EofToken) Extract() {

}

// return the current character from the source
// call the source's CurrentChar() method
func (t *EofToken) CurrentChar() byte {
	return t.source.CurrentChar()
}

// return the next character from the source
func (t *EofToken) NextChar() byte {
	return t.source.NextChar()
}

// return the source character *following* the current character without consuming the current character
func (t *EofToken) PeekChar() byte {
	return t.source.PeekChar()
}

func (t *EofToken) GetTokenType() TokenType {
	return t.tokenType
}

func (t *EofToken) GetLineNumber() int {
	return t.lineNum
}

func (t *EofToken) GetPosition() int {
	return t.position
}

func (t *EofToken) GetText() string {
	return t.text
}

func (t *EofToken) GetValue() TokenValue {
	return t.value
}
