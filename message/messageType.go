package message

type MessageType int

type SourceLineEvent struct {
	LineNum int
	Line    string
}

type ParserSummaryEvent struct {
	LineNum     int
	ErrorCount  int
	ElapsedTime float64
}

const (
	SOURCE_LINE MessageType = iota
	SYNTAX_ERROR
	PARSER_SUMMARY
	INTERPRETER_SUMMARY
	COMPILER_SUMMARY
	MISCELLANEOUS
	TOKEN
	ASSIGN
	FETCH
	BREAKPOINT
	RUNTIME_EERROR
	CALL
	RETURN
)
