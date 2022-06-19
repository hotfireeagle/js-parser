package compiler

import "jsj/message"

type CompilerMessageListener struct {
}

func (cml *CompilerMessageListener) MessageReceived(messageObj message.Message) {
	messageType := messageObj.GetMessageType()
	messageBody := messageObj.GetBody()

	switch messageType {
	case message.COMPILER_SUMMARY:
		{
			messageBody.(*message.CompilerSummaryEvent).Log()
			break
		}
	}
}
