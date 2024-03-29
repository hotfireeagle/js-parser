package scanner

import (
	"jsj/frontend/source"
	"jsj/frontend/token"
	"jsj/utils"
)

type Scanner struct {
	// source, which is code file
	source *source.Source

	// current token
	token token.Token
}

// Scanner Object Constructor
func ScannerConstructor(s *source.Source) *Scanner {
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
	scannerInstance.skipWhiteSpace()
	scannerInstance.skipEol()
	scannerInstance.skipComment()

	var tokenInstance token.Token

	// the current character determines what type of token to construct
	currentChar := scannerInstance.CurrentChar()

	if currentChar == utils.EOF {
		tokenInstance = token.EofTokenConstructor(scannerInstance.source)
	} else if utils.CheckIsJSWordPrefix(currentChar) {
		tokenInstance = token.WordTokenConstructor(scannerInstance.source)
	} else if currentChar == '"' {
		tokenInstance = token.DoubleQuoteTokenConstructor(scannerInstance.source)
	} else if token.CheckIsSingleCharacterToken(currentChar) {
		tokenInstance = token.SingleCharacterTokenConstructor(scannerInstance.source)
	} else if currentChar == '+' {
		if token.CheckIsNumberBeginCharacter(scannerInstance.PeekChar()) {
			scannerInstance.NextChar()
			tokenInstance = token.NumberTokenConstructor(scannerInstance.source, true)
		} else {
			tokenInstance = token.DeterminedPlusTokenConstructor(scannerInstance.source)
		}
	} else if currentChar == '-' {
		if token.CheckIsNumberBeginCharacter(scannerInstance.PeekChar()) {
			scannerInstance.NextChar()
			tokenInstance = token.NumberTokenConstructor(scannerInstance.source, false)
		} else {
			tokenInstance = token.DeterminedMinusTokenConstructor(scannerInstance.source)
		}
	} else if currentChar == '*' {
		tokenInstance = token.DeterminedMultiplyTokenConstructor(scannerInstance.source)
	} else if currentChar == '/' {
		// 在这个分支即提取出了/, /=, 也提取了行注释
		tokenInstance = token.DeterminedSlashTokenConstructor(scannerInstance.source)
	} else if currentChar == '%' {
		tokenInstance = token.DeterminedRemainderTokenConstructor(scannerInstance.source)
	} else if currentChar == '&' {
		// 提取 &, &^, &=, &^=, &&
		tokenInstance = token.DeterminedAndTokenConstructor(scannerInstance.source)
	} else if currentChar == '|' {
		// 提取 |, |=, ||
		tokenInstance = token.DeterminedOrTokenConstructor(scannerInstance.source)
	} else if currentChar == '^' {
		// 提取 ^, ^=
		tokenInstance = token.DeterminedExclusiveOrTokenConstructor(scannerInstance.source)
	} else if currentChar == '<' {
		// 提取<, <<, <<=,
		tokenInstance = token.DeterminedLessTokenConstructor(scannerInstance.source)
	} else if currentChar == '>' {
		// 提取>, >>, >>>, >>=, >>>=, >=
		tokenInstance = token.DeterminedGreaterTokenConstructor(scannerInstance.source)
	} else if currentChar == '=' {
		// 提取=, ==, ===
		tokenInstance = token.DeterminedEqualTokenConstructor(scannerInstance.source)
	} else if token.CheckIsNumberBeginCharacter(currentChar) {
		tokenInstance = token.NumberTokenConstructor(scannerInstance.source, true)
	} else {
		tokenInstance = token.BaseTokenConstructor(scannerInstance.source)
	}

	// TODO:
	// else if unicode.IsDigit(rune(currentChar)) {

	// } else if currentChar == '\'' {

	// }

	return tokenInstance
}

// return the current character from the source
func (scannerInstance *Scanner) CurrentChar() byte {
	return scannerInstance.source.CurrentChar()
}

// return the next character from the source
func (scannerInstance *Scanner) NextChar() byte {
	return scannerInstance.source.NextChar()
}

func (scannerInstance *Scanner) PeekChar() byte {
	return scannerInstance.source.PeekChar()
}

// 跳过whitespace, 并且把它们给consume掉
func (si *Scanner) skipWhiteSpace() {
	currentCharacter := si.CurrentChar()

	for currentCharacter == 32 {
		currentCharacter = si.NextChar()
	}
}

// skip end of line
func (si *Scanner) skipEol() {
	currentChar := si.CurrentChar()

	for currentChar == utils.EOL {
		currentChar = si.NextChar()
	}
}

// TODO: 跳过注释
func (si *Scanner) skipComment() {

}
