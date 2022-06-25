package token

func CheckIsSingleCharacterToken(c byte) bool {
	_, ok := SingleToken[c]
	return ok
}
