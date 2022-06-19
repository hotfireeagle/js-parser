package source

import (
	"bufio"
	"errors"
	"io"
	"jsj/message"
	"jsj/utils"
)

var messageHandler = message.MessageHandlerConstructor()

type Source struct {
	reader     bufio.Reader
	line       string
	lineNum    int
	currentPos int
}

func SourceConstructor(r bufio.Reader) *Source {
	return &Source{
		reader:     r,
		currentPos: -2,
	}
}

// return the source character at current positon
func (sourceInstance *Source) CurrentChar() byte {
	if sourceInstance.currentPos == -2 {
		// first time
		isEof := sourceInstance.readLine()

		if isEof {
			return utils.EOF
		}

		return sourceInstance.NextChar()
	} else if (sourceInstance.currentPos == -1) || sourceInstance.currentPos == len(sourceInstance.line) {
		// at end of line
		return utils.EOL
	} else if sourceInstance.currentPos > len(sourceInstance.line) {
		// need read the next line
		isEof := sourceInstance.readLine()
		if isEof {
			return utils.EOF
		}
		return sourceInstance.NextChar()
	} else {
		return sourceInstance.line[sourceInstance.currentPos]
	}
}

// consumer the current source character and return the next character
func (sourceInstance *Source) NextChar() byte {
	sourceInstance.currentPos += 1
	return sourceInstance.CurrentChar()
}

// return the source character *following* the current character without consuming the current character
func (sourceInstance *Source) PeekChar() byte {
	// TODO: why we need first call it
	sourceInstance.CurrentChar()

	if sourceInstance.line == "" {
		return utils.EOF
	}

	nextPos := sourceInstance.currentPos + 1

	if nextPos < len(sourceInstance.line) {
		return sourceInstance.line[nextPos]
	}

	return utils.EOL
}

func (sourceInstance *Source) GetLineNum() int {
	return sourceInstance.lineNum
}

func (sourceInstance *Source) GetPosition() int {
	return sourceInstance.currentPos
}

func (sourceInstance *Source) readLine() bool {
	l, err := sourceInstance.reader.ReadString('\n')
	if err != nil {
		if errors.Is(err, io.EOF) {
			return true
		}
		panic(err)
	}

	sourceInstance.line = l

	sourceInstance.currentPos = -1

	if l != "" {
		messageHandler.SendMessage(message.MessageConstructor(message.SOURCE_LINE, &message.SourceLineEvent{
			LineNum: sourceInstance.lineNum,
			Line:    l,
		}))
		sourceInstance.lineNum += 1
	}

	return false
}

func (sourceInstance *Source) AddMessageListener(listener message.MessageListener) {
	messageHandler.AddListener(listener)
}
