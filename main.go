package main

import (
	"fmt"
	"jsj/frontend/scanner"
)

func main() {
	scanner := scanner.New("/Users/smallhai/Desktop/empty")

	fmt.Println(scanner.CurrentChar())
}
