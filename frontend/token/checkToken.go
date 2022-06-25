package token

import "unicode"

func CheckIsSingleCharacterToken(c byte) bool {
	_, ok := SingleToken[c]
	return ok
}

func CheckIsNumberBeginCharacter(c byte) bool {
	return unicode.IsNumber(rune(c)) || c == '.'
}
