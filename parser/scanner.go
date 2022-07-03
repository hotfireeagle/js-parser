package parser

import (
	"bufio"
	"errors"
	"io"
	"jsj/utils"
	"os"
	"strings"
)

// TODO: 支持中文
// TODO: 支持二进制、八进制、十六进制
// TODO: 支持一些基础的语法检查
// TODO: 支持单引号字符串
// TODO: 支持template
// TODO: 支持正则表达式

type Scanner struct {
	// 文件读取器
	reader *bufio.Reader

	file *os.File

	// 要读取的文件的绝对路径
	filePath string

	// 当前读取的文件行内容
	lineRaw string

	// 当前读取的行号
	lineNumber int

	// 当前行的长度
	lineLength int

	// 当前读取行的列索引
	lineColumn int

	// 当前是否已经读取到了文件最后一行
	hasEOF bool

	// 是否命中需要进行多行注释检查的校验
	hitByMultiLineComment bool
}

func ScannerConstructor(fp string) *Scanner {
	fileObj, openErr := os.Open(fp)

	if openErr != nil {
		panic(openErr)
	}

	fr := &Scanner{
		filePath: fp,
		file:     fileObj,
		reader:   bufio.NewReader(fileObj),
	}

	// 文件一打开即开始读取一行
	fr.readline()
	return fr
}

// 对文件进行按行读取
func (s *Scanner) readline() {
	lineContent, readErr := s.reader.ReadString('\n')

	if readErr != nil {
		if errors.Is(readErr, io.EOF) {
			s.hasEOF = true
		} else {
			s.CloseFile()
			panic(readErr)
		}
	}

	s.setLineRaw(lineContent)
}

// 当前行读取的字符，不会对它进行消费，只是简单获取
func (fr *Scanner) CurrentChar() (byte, error) {
	if fr.lineColumn >= fr.lineLength {
		// 表示当前行已经读取完毕
		if fr.hasEOF {
			return 0, utils.ErrEof
		}
		return 0, utils.ErrEol
	} else {
		return fr.lineRaw[fr.lineColumn], nil
	}
}

// 当前行读取的字符，不会对它进行消费，只是简单获取
func (fr *Scanner) NextChar() (byte, error) {
	fr.lineColumn += 1
	if fr.lineColumn >= fr.lineLength {
		// 表示当前行已经读取完毕
		// 那么读取下一行
		if fr.hasEOF {
			// 表示文件已经读取完毕，没有读取下一行的机会了
			return 0, utils.ErrEof
		} else {
			// 读取下一行
			fr.readline()
			return fr.CurrentChar()
		}
	} else {
		return fr.lineRaw[fr.lineColumn], nil
	}
}

// 获取下几个字符，但不consume
func (s *Scanner) Peek(skip int) (byte, error) {
	nextPos := s.lineColumn + skip
	if nextPos >= s.lineLength {
		return 0, utils.ErrEol
	}
	return s.lineRaw[nextPos], nil
}

func (s *Scanner) SubStr(length int) string {
	start := s.lineColumn
	endIdx := start + length

	if endIdx > len(s.lineRaw) {
		return s.lineRaw[start:]
	}
	return s.lineRaw[start:endIdx]
}

// 提取出标识符，可能是变量标识符，也可能是关键字
func (s *Scanner) getIdentifier() string {
	var identifier strings.Builder

	currentChar, err := s.CurrentChar()

	if err != nil {
		panic(err)
	}

	for err != utils.ErrEof && utils.IsIdentifierPart(currentChar) {
		identifier.WriteByte(currentChar)
		currentChar, err = s.NextChar()

		if err == utils.ErrEof {
			break
		}

		if err != nil {
			panic(err)
		}
	}

	return identifier.String()
}

// 提取出单词，单词可能是变量标识符，也可能是关键字，也可能是一些字面量
// https://tc39.github.io/ecma262/#sec-names-and-keywords
// TODO: 支持以0x5c反破折号开始的特殊标识符
func (s *Scanner) scanIdentifier() *Token {
	id := s.getIdentifier()

	var tokenType TokenType

	if len(id) == 1 {
		// 特殊边界，在ECMAScript中，没有字符串关键字是一个字符的，所以它肯定是变量名
		tokenType = Identifier
	} else if s.isKeyword(id) {
		tokenType = Keyword
	} else if id == "null" {
		tokenType = NullLiteral
	} else if id == "true" || id == "false" {
		tokenType = BooleanLiteral
	} else {
		tokenType = Identifier
	}

	return TokenConstructor(tokenType, id, s.lineNumber, s.lineColumn-len(id))
}

// 提取出各种各样的标点符号，比如说运算操作符，又比如；，：{} [] ()等
func (s *Scanner) scanPunctuator() *Token {
	c, err := s.CurrentChar()

	str := string(c)

	if err != nil {
		panic(err)
	}

	switch str {
	case "(":
	case "{":
		// TODO: 入栈，检查括号是否匹配
		break
	case ".":
		temp := s.SubStr(3)

		if temp == "..." {
			// Spread operator: ...
			str = "..."
		}
		break
	case "}":
		// TODO: 出栈，检查括号是否匹配
		break
	case "?":
		temp := s.SubStr(2)
		if temp == "??" {
			str = "??"
		} else if temp == "?." {
			// TODO: "?." in "foo?.3:0" should not be treated as optional chaining.
			// See https://github.com/tc39/proposal-optional-chaining#notes
			str = "?."
		}
		break
	case ")":
	case ";":
	case ",":
	case "[":
	case "]":
	case ":":
	case "~":
		break
	default:
		// 4-character punctuator.
		str = s.SubStr(4)

		if str != ">>>=" {
			// 3-character punctuators.
			str = s.SubStr(3)

			if !(str == "==" || str == "!==" || str == ">>>" || str == "<<=" || str == ">>=" || str == "**=") {
				// 2-character punctuators.
				str = s.SubStr(2)
				m2 := map[string]bool{
					"&&": true,
					"||": true,
					"??": true,
					"==": true,
					"!=": true,
					"+=": true,
					"-=": true,
					"*=": true,
					"/=": true,
					"++": true,
					"--": true,
					"<<": true,
					">>": true,
					"&=": true,
					"|=": true,
					"^=": true,
					"%=": true,
					"<=": true,
					">=": true,
					"=>": true,
					"**": true,
				}
				if !m2[str] {
					// 1-character punctuators.
					str = s.SubStr(1)
					m1 := map[string]bool{
						"<": true,
						">": true,
						"=": true,
						"!": true,
						"+": true,
						"-": true,
						"*": true,
						"%": true,
						"&": true,
						"|": true,
						"^": true,
						"/": true,
					}
					if !m1[str] {
						panic("Unexpected punctuator: " + str)
					}
				}
			}
		} else {
			panic("Unexpected punctuator: " + str)
		}
	}

	// comsume
	strLength := len(str)
	for i := 0; i < strLength; i++ {
		s.NextChar()
	}

	return TokenConstructor(Punctuator, str, s.lineNumber, s.lineColumn-len(str))
}

// 提取出十进制数字
// TODO: 支持二进制、八进制、16进制
func (s *Scanner) scanNumericLiteral() *Token {
	ch, err := s.CurrentChar()

	if err != nil {
		panic(err)
	}

	var num strings.Builder
	var dotHasBeenSeen bool

	for utils.IsNumberPart(ch) {
		if ch == '.' {
			if !dotHasBeenSeen {
				dotHasBeenSeen = true
				num.WriteString(".")
				ch, err = s.NextChar()

				if err == utils.ErrEof {
					break
				}
			} else {
				panic("Unexpected dot")
			}
		} else {
			num.WriteString(string(ch))
			ch, err = s.NextChar()
			if err == utils.ErrEof {
				break
			}
		}
	}

	if utils.IsIdentifierBegin(ch) {
		syntaxErr := SyntaxErrorConstructor(s, UnexpectedTokenIllegal)
		syntaxErr.Fatal()
	}

	return TokenConstructor(NumericLiteral, num.String(), s.lineNumber, s.lineColumn-len(num.String()))
}

func (s *Scanner) scanStringLiteral() *Token {
	ch, err := s.NextChar()

	if errors.Is(err, utils.ErrEof) {
		panic(err)
	}

	var result strings.Builder

	for ch != '"' {
		result.WriteByte(ch)
		ch, err = s.NextChar()
		if errors.Is(err, utils.ErrEof) {
			// TODO:
			panic("Unexpected EOF")
		}
	}

	if ch == '"' {
		s.NextChar() //  consume the end "
	}

	return TokenConstructor(StringLiteral, result.String(), s.lineNumber, s.lineColumn-len(result.String()))
}

// TODO: 提取匹配模式
func (s *Scanner) scanRegExp(steps int) *Token {
	var regexpStr strings.Builder

	c, _ := s.CurrentChar()
	regexpStr.WriteByte(c)

	for j := 0; j < steps; j++ {
		c, _ = s.NextChar()
		regexpStr.WriteByte(c)
	}

	s.NextChar()

	return TokenConstructor(RegularExpression, regexpStr.String(), s.lineNumber, s.lineColumn-len(regexpStr.String()))
}

func (s *Scanner) Lex() *Token {
	cp, err := s.CurrentChar()

	if errors.Is(err, utils.ErrEof) {
		if s.hitByMultiLineComment {
			// 文件结尾了，多行注释还没有正确闭合的话，那么报错
			syntaxErr := SyntaxErrorConstructor(s, UnexpectedEOF)
			syntaxErr.Fatal()
		}
		return TokenConstructor(EOF, "", s.lineNumber, s.lineColumn)
	}

	if err != nil {
		panic(err)
	}

	if cp == utils.BREAK {
		s.NextChar()
		return s.Lex()
	}

	if cp == utils.SPACE {
		s.NextChar()
		return s.Lex()
	}

	// 提取多行注释的结尾标志，遇不到的话，那么把他当成注释进行考虑
	if s.hitByMultiLineComment {
		return s.skipMultiLineCommentWhenHitAlready()
	}

	// 跳过单行注释的处理
	if cp == '/' {
		nextCp, err2 := s.Peek(1)
		if !errors.Is(err2, utils.ErrEol) && nextCp == '/' {
			// 单行注释的开始标志
			// 直接开始读取下一行
			s.readline()
			return s.Lex()
		}
	}

	// 多行注释开始标志的判断
	if cp == '/' {
		nextCp, perr2 := s.Peek(1)
		if !errors.Is(perr2, utils.ErrEol) && nextCp == '*' {
			// 设置当前已经进入多行注释校验，期望见到多行注释的结尾标志
			s.setHitByMultiLineComment()
			return s.Lex()
		}
	}

	// 在前面已经过滤掉了注释
	// 所以这里的/要么是正则表达式，要么是操作符
	// 如果是操作符的话，那么前面不可能是空token、等号token、圆括号token，中括号token
	if cp == '/' {
		if prevExtractToken == nil || (prevExtractToken.GetValue() == "=" || prevExtractToken.GetValue() == "(" || prevExtractToken.GetValue() == "[") {
			// 判断是不是正则表达式
			for i := s.lineColumn + 1; i < s.lineLength; i++ {
				ccp := s.lineRaw[i]
				if ccp == '/' {
					// 是正则表达式
					return s.scanRegExp(i - s.lineColumn)
				}
			}
		}
	}

	if utils.IsIdentifierBegin(cp) {
		return s.scanIdentifier()
	}

	if cp == '(' || cp == ')' || cp == ';' {
		return s.scanPunctuator()
	}

	if cp == '"' {
		return s.scanStringLiteral()
	}

	if cp == '.' {
		n1, err := s.Peek(1)
		if err != nil {
			panic(err)
		}
		if utils.IsDecimalDigit(n1) {
			return s.scanNumericLiteral()
		}
		return s.scanPunctuator()
	}

	// 16进制的处理
	// if cp == '0' {
	// 	ncp, ncperr := s.Peek(1)
	// 	if !errors.Is(ncperr, utils.ErrEol) && (ncp == 'x' || ncp == 'X') {

	// 	}
	// }

	if utils.IsDecimalDigit(cp) {
		return s.scanNumericLiteral()
	}

	return s.scanPunctuator()
}

// 判断一个单词字符串是否是ECMAScript关键字
func (s *Scanner) isKeyword(id string) bool {
	switch len(id) {
	case 2:
		return (id == "if") || (id == "in") || (id == "do")
	case 3:
		return (id == "var") || (id == "for") || (id == "new") ||
			(id == "try") || (id == "let")
	case 4:
		return (id == "this") || (id == "else") || (id == "case") ||
			(id == "void") || (id == "with") || (id == "enum")
	case 5:
		return (id == "while") || (id == "break") || (id == "catch") ||
			(id == "throw") || (id == "const") || (id == "yield") ||
			(id == "class") || (id == "super")
	case 6:
		return (id == "return") || (id == "typeof") || (id == "delete") ||
			(id == "switch") || (id == "export") || (id == "import")
	case 7:
		return (id == "default") || (id == "finally") || (id == "extends")
	case 8:
		return (id == "function") || (id == "continue") || (id == "debugger")
	case 10:
		return (id == "instanceof")
	default:
		return false
	}
}

// 携带副作用去更新行内容
// 副作用1: 更新行号
// 副作用2: 更新列号
// 副作用3: 如果是单行注释的话，那么跳过
func (s *Scanner) setLineRaw(str string) {
	s.lineRaw = str
	s.lineLength = len(str)
	s.lineNumber += 1
	s.lineColumn = 0
}

func (s *Scanner) skipMultiLineCommentWhenHitAlready() *Token {
	cr, _ := s.CurrentChar()

	if cr == '*' {
		nxc, perr := s.Peek(1)
		if !errors.Is(perr, utils.ErrEol) && nxc == '/' {
			// 多行注释的结束标志
			s.hitByMultiLineComment = false
			s.NextChar()
			s.NextChar()
			return s.Lex()
		} else {
			s.NextChar()
			return s.Lex()
		}
	} else {
		s.NextChar()
		return s.Lex()
	}
}

func (s *Scanner) setHitByMultiLineComment() {
	s.hitByMultiLineComment = true
	s.NextChar() // consume /
	s.NextChar() // consume *
}

func (fr *Scanner) GetFilePath() string {
	return fr.filePath
}

func (fr *Scanner) GetLineRaw() string {
	return fr.lineRaw
}

func (fr *Scanner) GetlineNumber() int {
	return fr.lineNumber
}

func (fr *Scanner) GetLineColumn() int {
	return fr.lineColumn
}

func (fr *Scanner) GetHasEOF() bool {
	return fr.hasEOF
}

func (fr *Scanner) CloseFile() {
	fr.file.Close()
}
