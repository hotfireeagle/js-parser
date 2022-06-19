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

type CompilerSummaryEvent struct {
	// number of instructions generated
	InstructionCount int

	ElapsedTime float64
}

type InterpreterSummaryEvent struct {
	// number of statements executed
	ExecutionCount int

	// number of errors
	RuntimeErrorCount int

	// execution time
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
