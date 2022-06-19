package parser

type SyntaxErrorCodeEnum int

const (
	// 在前面已指定过
	// TODO:
	ALREADY_FORWARD SyntaxErrorCodeEnum = iota

	// 重复定义了同一个标识符
	IDENTIFIER_REDEFINE

	// 标识符不存在
	IDENTIFIER_UNDEFINED

	// 不匹配的赋值
	// TODO:
	INCOMPATIBLE_ASSIGNMENT

	// 不匹配的类型
	// TODO:
	INCOMPATIBLE_TYPES

	// 无效赋值
	INVALID_ASSIGMENT

	// 无效字符
	INVALID_CHARACTER

	// 无效表达式
	INVALID_EXPRESSION

	// 无效数字
	INVALID_NUMBER

	// 无效语句
	INVALID_STATEMENT

	// 无效类型
	// TODO:
	INVALID_TYPE

	// 缺少:
	MISSING_COLON

	// 缺少,
	MISSING_COMMA

	// 缺少=
	MISSING_EQUALS

	// 缺少[
	MISSING_LEFT_BRACKET

	// 缺少]
	MISSING_RIGHT_BRACKET

	// 不是一个类型标识符
	NOT_TYPE_IDENTIFIER

	// 栈溢出
	STACK_OVERFLOW

	UNEXPECTED_EOF

	UNEXPECTED_TOKEN

	IO_ERROR

	TOO_MANY_ERRORS
)

type SyntaxErrorCode struct {
	errorType SyntaxErrorCodeEnum
	// exit status
	status int
	// error message
	message string
}

func SyntaxErrorCodeConstructor(e SyntaxErrorCodeEnum, m string) *SyntaxErrorCode {
	return &SyntaxErrorCode{
		errorType: e,
		message:   m,
	}
}

func (src *SyntaxErrorCode) GetStatus() int {
	return src.status
}

func (src *SyntaxErrorCode) ToString() string {
	return src.message
}
