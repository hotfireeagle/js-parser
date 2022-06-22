package token

type TokenType int

// Token类型
const (
	ILLEGAL TokenType = iota
	EOF
	ERROR
	COMMENT
	KEYWORD

	STRING
	BOOLEAN
	NULL
	NUMBER
	IDENTIFIER

	PLUS      // +
	MINUS     // -
	MULTIPLY  // *
	SLASH     // /
	REMAINDER // %

	AND                  // &
	OR                   // |
	EXCLUSIVE_OR         // ^
	SHIFT_LEFT           // <<
	SHIFT_RIGHT          // >>
	UNSIGNED_SHIFT_RIGHT // >>>
	AND_NOT              // &^

	ADD_ASSIGN       // +=
	SUBTRACT_ASSIGN  // -=
	MULTIPLY_ASSIGN  // *=
	QUOTIENT_ASSIGN  // /=
	REMAINDER_ASSIGN // %=

	AND_ASSIGN                  // &=
	OR_ASSIGN                   // |=
	EXCLUSIVE_OR_ASSIGN         // ^=
	SHIFT_LEFT_ASSIGN           // <<=
	SHIFT_RIGHT_ASSIGN          // >>=
	UNSIGNED_SHIFT_RIGHT_ASSIGN // >>>=
	AND_NOT_ASSIGN              // &^=

	LOGICAL_AND // &&
	LOGICAL_OR  // ||
	INCREMENT   // ++
	DECREMENT   // --

	EQUAL        // ==
	STRICT_EQUAL // ===
	LESS         // <
	GREATER      // >
	ASSIGN       // =
	NOT          // !

	BITWISE_NOT // ~

	NOT_EQUAL        // !=
	STRICT_NOT_EQUAL // !==
	LESS_OR_EQUAL    // <=
	GREATER_OR_EQUAL // >=

	LEFT_PARENTHESIS // (
	LEFT_BRACKET     // [
	LEFT_BRACE       // {
	COMMA            // ,
	PERIOD           // .

	RIGHT_PARENTHESIS // )
	RIGHT_BRACKET     // ]
	RIGHT_BRACE       // }
	SEMICOLON         // ;
	COLON             // :
	QUESTION_MARK     // ?

	firstKeyword // nolint: deadcode TODO:
	IF
	IN
	DO

	VAR
	LET
	CONST
	FOR
	NEW
	TRY

	THIS
	ELSE
	CASE
	VOID
	WITH

	WHILE
	BREAK
	CATCH
	THROW

	RETURN
	TYPEOF
	DELETE
	SWITCH

	DEFAULT
	FINALLY

	FUNCTION
	CONTINUE
	DEBUGGER

	INSTANCEOF

	CLASS
	ENUM
	EXPORT
	EXTENDS
	IMPORT
	SUPER
	IMPLEMENTS
	INTERFACE
	PACKAGE
	PRIVATE
	PROTECTED
	PUBLIC
	STATIC
	lastKeyword // nolint: deadcode TODO: del
)

// ECMAScript规范的保留关键字
var KeyWordMap map[string]TokenType = map[string]TokenType{
	"if":         IF,
	"in":         IN,
	"do":         DO,
	"var":        VAR,
	"new":        NEW,
	"try":        TRY,
	"this":       THIS,
	"else":       ELSE,
	"case":       CASE,
	"void":       VOID,
	"with":       WITH,
	"while":      WHILE,
	"break":      BREAK,
	"catch":      CATCH,
	"throw":      THROW,
	"return":     RETURN,
	"typeof":     TYPEOF,
	"delete":     DELETE,
	"switch":     SWITCH,
	"default":    DEFAULT,
	"finally":    FINALLY,
	"function":   FUNCTION,
	"continue":   CONTINUE,
	"debugger":   DEBUGGER,
	"instanceof": INSTANCEOF,
	"const":      CONST,
	"class":      CLASS,
	"enum":       ENUM,
	"export":     EXPORT,
	"extends":    EXTENDS,
	"import":     IMPORT,
	"super":      SUPER,
	"implements": IMPLEMENTS,
	"interface":  INTERFACE,
	"let":        LET,
	"package":    PACKAGE,
	"private":    PRIVATE,
	"protected":  PROTECTED,
	"public":     PUBLIC,
	"static":     STATIC,
}

var token2strMap = map[TokenType]string{
	ILLEGAL:                     "ILLEGAL",
	EOF:                         "EOF",
	COMMENT:                     "COMMENT",
	KEYWORD:                     "KEYWORD",
	STRING:                      "STRING",
	BOOLEAN:                     "BOOLEAN",
	NULL:                        "NULL",
	NUMBER:                      "NUMBER",
	IDENTIFIER:                  "IDENTIFIER",
	PLUS:                        "+",
	MINUS:                       "-",
	MULTIPLY:                    "*",
	SLASH:                       "/",
	REMAINDER:                   "%",
	AND:                         "&",
	OR:                          "|",
	EXCLUSIVE_OR:                "^",
	SHIFT_LEFT:                  "<<",
	SHIFT_RIGHT:                 ">>",
	UNSIGNED_SHIFT_RIGHT:        ">>>",
	AND_NOT:                     "&^",
	ADD_ASSIGN:                  "+=",
	SUBTRACT_ASSIGN:             "-=",
	MULTIPLY_ASSIGN:             "*=",
	QUOTIENT_ASSIGN:             "/=",
	REMAINDER_ASSIGN:            "%=",
	AND_ASSIGN:                  "&=",
	OR_ASSIGN:                   "|=",
	EXCLUSIVE_OR_ASSIGN:         "^=",
	SHIFT_LEFT_ASSIGN:           "<<=",
	SHIFT_RIGHT_ASSIGN:          ">>=",
	UNSIGNED_SHIFT_RIGHT_ASSIGN: ">>>=",
	AND_NOT_ASSIGN:              "&^=",
	LOGICAL_AND:                 "&&",
	LOGICAL_OR:                  "||",
	INCREMENT:                   "++",
	DECREMENT:                   "--",
	EQUAL:                       "==",
	STRICT_EQUAL:                "===",
	LESS:                        "<",
	GREATER:                     ">",
	ASSIGN:                      "=",
	NOT:                         "!",
	BITWISE_NOT:                 "~",
	NOT_EQUAL:                   "!=",
	STRICT_NOT_EQUAL:            "!==",
	LESS_OR_EQUAL:               "<=",
	GREATER_OR_EQUAL:            ">=",
	LEFT_PARENTHESIS:            "(",
	LEFT_BRACKET:                "[",
	LEFT_BRACE:                  "{",
	COMMA:                       ",",
	PERIOD:                      ".",
	RIGHT_PARENTHESIS:           ")",
	RIGHT_BRACKET:               "]",
	RIGHT_BRACE:                 "}",
	SEMICOLON:                   ";",
	COLON:                       ":",
	QUESTION_MARK:               "?",
	IF:                          "if",
	IN:                          "in",
	DO:                          "do",
	VAR:                         "var",
	FOR:                         "for",
	NEW:                         "new",
	TRY:                         "try",
	THIS:                        "this",
	ELSE:                        "else",
	CASE:                        "case",
	VOID:                        "void",
	WITH:                        "with",
	WHILE:                       "while",
	BREAK:                       "break",
	CATCH:                       "catch",
	THROW:                       "throw",
	RETURN:                      "return",
	TYPEOF:                      "typeof",
	DELETE:                      "delete",
	SWITCH:                      "switch",
	DEFAULT:                     "default",
	FINALLY:                     "finally",
	FUNCTION:                    "function",
	CONTINUE:                    "continue",
	DEBUGGER:                    "debugger",
	INSTANCEOF:                  "instanceof",
	CONST:                       "const",
}

func GetTokenTypeName(t TokenType) string {
	v, ok := token2strMap[t]

	if !ok {
		// TODO:
		return "unknown and should fix"
	}

	return v
}
