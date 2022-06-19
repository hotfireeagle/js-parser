package interpreter

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

	messageLog := message.InterpreterSummaryEvent{
		ExecutionCount:    0,
		RuntimeErrorCount: 0,
		ElapsedTime:       elapsedTime,
	}

	messageObj := message.MessageConstructor(message.INTERPRETER_SUMMARY, messageLog)

	messageHandler.SendMessage(messageObj)
}
