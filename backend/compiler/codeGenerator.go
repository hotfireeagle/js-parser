package compiler

import (
	"jsj/intermediate"
	"jsj/message"
	"time"
)

var messageHandler = message.MessageHandlerConstructor()

type Backend struct {
	symTab *intermediate.SymTab
	icode  *intermediate.ICode
}

func (backendInstance *Backend) Process(icode *intermediate.ICode, symTab *intermediate.SymTab) {
	startTime := time.Now().UnixMilli()

	endTime := time.Now().UnixMilli()

	var elapsedTime float64 = (float64(endTime) - float64(startTime)) / 1000

	messageLog := message.CompilerSummaryEvent{
		InstructionCount: 0,
		ElapsedTime:      elapsedTime,
	}

	messageObj := message.MessageConstructor(message.COMPILER_SUMMARY, messageLog)

	messageHandler.SendMessage(messageObj)
}
