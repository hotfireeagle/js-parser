package main

import (
	"fmt"
	"jsj/parser"
)

func main() {
	scaner := parser.ScannerConstructor("/Users/smallhai/Desktop/test.js")

	tokenizer := parser.TokenizerConstructor(scaner)

	tokens := tokenizer.Tokenize()

	for _, token := range tokens {
		fmt.Println(token.TokenTypeToString(), token.GetValue())
	}
}
