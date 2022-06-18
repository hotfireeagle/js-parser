package token

func (tkn Token) String() string {
	return "token(" + token2string[tkn] + ")"
}

func IsKeyword(literal string) (Token, bool) {
	if keyword, exists := keywordTable[literal]; exists {
		if keyword.futureKeyword {
			// 表明是新特性
			return KEYWORD, keyword.strict
		}
		return keyword.token, false
	}
	return 0, false
}
