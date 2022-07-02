package parser

import (
	"fmt"
	"jsj/utils"
	"testing"
)

var fr = FileReaderConstructor("/Users/smallhai/Desktop/test.js")

func TestCurrentChar(t *testing.T) {
	cr, err := fr.CurrentChar()

	for err != utils.ErrEof {
		fmt.Println("cr:", cr, string(cr))
		cr, err = fr.NextChar()
	}
}
