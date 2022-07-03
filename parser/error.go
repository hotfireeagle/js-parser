package parser

import (
	"fmt"
	"os"
)

type SyntaxErrType int

const (
	UnexpectedTokenIllegal SyntaxErrType = iota
	UnexpectedEOF
)

var errCn = map[SyntaxErrType]string{
	UnexpectedTokenIllegal: "Unexpected token ILLEGAL",
	UnexpectedEOF:          "Unexpected end of file",
}

type SyntaxError struct {
	filePath    string
	lineContent string
	lineNumber  int
	lineColumn  int
	errorType   SyntaxErrType
}

func SyntaxErrorConstructor(s *Scanner, errType SyntaxErrType) *SyntaxError {
	return &SyntaxError{
		filePath:    s.filePath,
		lineContent: s.lineRaw,
		lineNumber:  s.lineNumber,
		lineColumn:  s.lineColumn,
		errorType:   errType,
	}
}

func (se *SyntaxError) Fatal() {
	fmt.Println(errCn[se.errorType])
	fmt.Println()

	var str string

	for i := 0; i < se.lineColumn; i++ {
		str += " "
	}

	str += "^"

	fmt.Println(se.filePath, "in line ", se.lineNumber, "in column ", se.lineColumn)
	fmt.Println()

	fmt.Println(se.lineContent)
	fmt.Println(str)
	os.Exit(1)
}
