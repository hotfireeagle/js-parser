package parser

type TokenType int

const (
	// 布尔值字面量
	BooleanLiteral TokenType = iota

	// 文件结束符，文件的最后的空白行
	EOF

	// 变量标识符
	Identifier

	// 语言关键字
	Keyword

	// null字面量
	NullLiteral

	// 数字字面量
	NumericLiteral

	// 标点符合
	Punctuator

	// 字符串字面量
	StringLiteral

	// 正则表达式
	// TODO:
	RegularExpression

	// 模板字符串
	Template
)

var TokenName = map[TokenType]string{
	BooleanLiteral:    "Boolean",
	EOF:               "<end>",
	Identifier:        "Identifier",
	Keyword:           "Keyword",
	NullLiteral:       "Null",
	NumericLiteral:    "Numeric",
	Punctuator:        "Punctuator",
	StringLiteral:     "String",
	RegularExpression: "RegularExpression",
	Template:          "Template",
}

type Token struct {
	// token类型
	tokenType TokenType

	// token值
	value string

	// token所在的行号
	lineNumber int

	// token所在的列号
	lineColumn int
}

func TokenConstructor(
	tokenType TokenType,
	value string,
	ln int,
	lc int,
) *Token {
	return &Token{
		tokenType:  tokenType,
		value:      value,
		lineNumber: ln,
		lineColumn: lc,
	}
}

func (t *Token) GetTokenType() TokenType {
	return t.tokenType
}

func (t *Token) GetValue() string {
	return t.value
}

func (t *Token) TokenTypeToString() string {
	return TokenName[t.tokenType]
}

func (t *Token) GetLineNumber() int {
	return t.lineNumber
}
