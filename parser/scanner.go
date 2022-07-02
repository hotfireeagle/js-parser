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

	// 是否已经读取到文件末尾
	hasEOF bool
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
func (fr *Scanner) readline() {
	lineContent, readErr := fr.reader.ReadString('\n')

	if readErr != nil {
		if errors.Is(readErr, io.EOF) {
			fr.hasEOF = true
		} else {
			fr.CloseFile()
			panic(readErr)
		}
	}

	fr.lineRaw = lineContent
	fr.lineLength = len(fr.lineRaw)
	fr.lineColumn = 0
}

// 当前行读取的字符，不会对它进行消费，只是简单获取
func (fr *Scanner) CurrentChar() (byte, error) {
	if fr.lineColumn >= fr.lineLength {
		// 表示当前行已经读取完毕
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
	case "":
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

	return TokenConstructor(NumericLiteral, num.String(), s.lineNumber, s.lineColumn-len(num.String()))
}

func (s *Scanner) scanStringLiteral() *Token {
	ch, err := s.NextChar()

	if err == utils.ErrEof {
		panic(err)
	}

	var result strings.Builder

	for ch != '"' {
		result.WriteByte(ch)
		ch, err = s.NextChar()
		if err == utils.ErrEof {
			panic("Unexpected EOF")
		}
	}

	return TokenConstructor(StringLiteral, result.String(), s.lineNumber, s.lineColumn-len(result.String()))
}

func (s *Scanner) Lex() *Token {
	if s.hasEOF {
		return TokenConstructor(EOF, "", s.lineNumber, s.lineColumn)
	}

	cp, err := s.CurrentChar()

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