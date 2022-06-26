package command

import (
	"bufio"
	"flag"
	"jsj/backend"
	"jsj/backend/compiler"
	"jsj/backend/interpreter"
	"jsj/frontend/parser"
	"jsj/frontend/source"
	"jsj/intermediate"
	"os"
)

func BeginWork() {
	mode := flag.String("mode", "", "i means use intermediate mode")
	filePath := flag.String("filePath", "", "absolute file path")

	flag.Parse()

	fileObj, openFileErr := os.Open(*filePath)
	if openFileErr != nil {
		panic(openFileErr)
	}

	defer fileObj.Close()

	sourceInstance := source.SourceConstructor(*bufio.NewReader(fileObj))
	sourceInstance.AddMessageListener(&source.SourceMessageListener{})

	parserInstance := parser.ParserConstructor(sourceInstance)
	parserInstance.AddMessageListener(&parser.ParserMessageListener{})

	var backendInstance backend.Backend

	if *mode == "i" {
		backendInstance = interpreter.InterpreterConstructor()
		backendInstance.AddMessageListener(&interpreter.InterpreterMessageListener{})
	} else {
		backendInstance = compiler.CompilerConstructor()
		backendInstance.AddMessageListener(&compiler.CompilerMessageListener{})
	}

	parserInstance.Parse()

	icode := parserInstance.GetICode()
	symTab := parserInstance.GetSymTab()

	symTabStack := parserInstance.GetSymTabStack()

	cr := intermediate.CrossReferencerConstructor()
	cr.Print(symTabStack)

	backendInstance.Process(icode, symTab)
}
