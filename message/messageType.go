package message

import "fmt"

type MessageType int

type SourceLineEvent struct {
	LineNum int
	Line    string
}

func (sle *SourceLineEvent) Log() {
	format := "[ReadLineEvent]: %03d %s"
	fmt.Printf(format, sle.LineNum, sle.Line)
}

type ParserSummaryEvent struct {
	LineNum     int
	ErrorCount  int
	ElapsedTime float64
}

func (pse *ParserSummaryEvent) Log() {
	format := "\n %d source lines." +
		"\n %d syntax errors." +
		"\n %f seconds elapsed.\n"
	fmt.Printf(format, pse.LineNum, pse.ErrorCount, pse.ElapsedTime)
}

type CompilerSummaryEvent struct {
	// number of instructions generated
	InstructionCount int

	ElapsedTime float64
}

func (cse *CompilerSummaryEvent) Log() {
	format := "\n %d instructions generated." +
		"\n %f seconds total code generation time.\n"
	fmt.Printf(format, cse.InstructionCount, cse.ElapsedTime)
}

type InterpreterSummaryEvent struct {
	// number of statements executed
	ExecutionCount int

	// number of errors
	RuntimeErrorCount int

	// execution time
	ElapsedTime float64
}

func (ise *InterpreterSummaryEvent) Log() {
	format := "\n %d statements executed." +
		"\n %d runtime errors." +
		"\n %f seconds total execution time.\n"
	fmt.Printf(format, ise.ExecutionCount, ise.RuntimeErrorCount, ise.ElapsedTime)
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
