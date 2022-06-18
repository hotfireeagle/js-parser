package scanner

import (
	"jsj/frontend/source"
	"jsj/frontend/token"
)

type Scanner struct {
	// source, which is code file
	source source.Source

	// current token
	token token.Token
}

// Scanner Object Constructor
func ScannerConstructor(s source.Source) *Scanner {
	return &Scanner{
		source: s,
	}
}

// return current token
func (scannerInstance *Scanner) CurrentToken() token.Token {
	return scannerInstance.token
}

// return next token from the source
func (scannerInstance *Scanner) NextToken() token.Token {
	nt := scannerInstance.extractToken()
	return nt
}

// extract : 提取
// do the actual work of extracting token and return the next token from the source
func (scannerInstance *Scanner) extractToken() token.Token {
	return 1
}

// return the current character from the source
func (scannerInstance *Scanner) CurrentChar() byte {
	return scannerInstance.source.CurrentChar()
}

// return the next character from the source
func (scannerInstance *Scanner) NextChar() byte {
	return scannerInstance.source.NextChar()
}
