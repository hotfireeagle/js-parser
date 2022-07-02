package main

import (
	"jsj/out"
	"jsj/parser"
)

func main() {
	scanner := parser.ScannerConstructor("/Users/smallhai/Desktop/test.js")
	defer scanner.CloseFile()

	tokenizer := parser.TokenizerConstructor(scanner)

	tokens := tokenizer.Tokenize()

	out.SaveJsonTokens(tokens)
}
