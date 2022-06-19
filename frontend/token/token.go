package token

type TokenValue interface{}

type Token interface {
	Extract()

	CurrentChar() byte

	NextChar() byte

	PeekChar() byte

	GetTokenType() TokenType

	GetLineNumber() int
}
