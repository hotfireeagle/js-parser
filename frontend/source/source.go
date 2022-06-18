package source

import (
	"bufio"
	"jsj/message"
	"jsj/utils"
)

type Source struct {
	reader     bufio.Reader
	line       string
	lineNum    int
	currentPos int
	message.MessageHandler
}

func SourceConstructor(r bufio.Reader) *Source {
	return &Source{
		reader:         r,
		currentPos:     -2,
		MessageHandler: message.MessageHandlerConstructor(),
	}
}

// return the source character at current positon
func (sourceInstance *Source) CurrentChar() byte {
	if sourceInstance.currentPos == -2 {
		// first time
		sourceInstance.readLine()
		return sourceInstance.NextChar()
	} else if sourceInstance.line == "" {
		// at end of file
		return utils.EOF
	} else if (sourceInstance.currentPos == -1) || sourceInstance.currentPos == len(sourceInstance.line) {
		// at end of line
		return utils.EOL
	} else if sourceInstance.currentPos > len(sourceInstance.line) {
		// need read the next line
		sourceInstance.readLine()
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

func (sourceInstance *Source) readLine() {
	l, err := sourceInstance.reader.ReadString('\n')
	if err != nil {
		// TODO:
		panic(err)
	}

	sourceInstance.line = l

	sourceInstance.currentPos = -1

	if l != "" {
		sourceInstance.SendMessage(message.MessageConstructor(message.SOURCE_LINE, message.SourceLineEvent{
			LineNum: sourceInstance.lineNum,
			Line:    l,
		}))
		sourceInstance.lineNum += 1
	}
}

func (sourceInstance *Source) Close() {

}
