package parser

type Tokenizer struct {
	scanner *Scanner
}

func TokenizerConstructor(scanner *Scanner) *Tokenizer {
	return &Tokenizer{
		scanner: scanner,
	}
}

func (t *Tokenizer) GetNextToken() *Token {
	token := t.scanner.Lex()
	return token
}

func (t *Tokenizer) Tokenize() []*Token {
	result := make([]*Token, 0)

	for {
		token := t.GetNextToken()
		if token == nil || token.GetTokenType() == EOF {
			break
		}
		result = append(result, token)
	}

	return result
}
