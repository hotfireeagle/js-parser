package interpreter

import "jsj/message"

type InterpreterMessageListener struct {
}

func (iml *InterpreterMessageListener) MessageReceived(messageObj message.Message) {
	messageType := messageObj.GetMessageType()
	messageBody := messageObj.GetBody()

	switch messageType {
	case message.INTERPRETER_SUMMARY:
		{
			messageBody.(*message.InterpreterSummaryEvent).Log()
			break
		}
	}
}
