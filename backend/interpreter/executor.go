package interpreter

import (
	"jsj/intermediate"
	"jsj/message"
	"time"
)

var messageHandler = message.MessageHandlerConstructor()

type Interpreter struct {
	symTab *intermediate.SymTab
	icode  *intermediate.ICode
}

func InterpreterConstructor() *Interpreter {
	return &Interpreter{}
}

func (InterpreterInstance *Interpreter) Process(icode *intermediate.ICode, symTab *intermediate.SymTab) {
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

func (interpreterInstance *Interpreter) GetSymTab() *intermediate.SymTab {
	return interpreterInstance.symTab
}

func (interpreterInstance *Interpreter) GetICode() *intermediate.ICode {
	return interpreterInstance.icode
}

func (interpreterInstance *Interpreter) AddMessageListener(listener message.MessageListener) {
	messageHandler.AddListener(listener)
}
