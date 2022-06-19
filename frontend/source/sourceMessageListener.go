package source

import (
	"jsj/message"
)

type SourceMessageListener struct {
}

func (s *SourceMessageListener) MessageReceived(messageObj message.Message) {
	messageType := messageObj.GetMessageType()
	messageBody := messageObj.GetBody()

	switch messageType {
	case message.SOURCE_LINE:
		{
			messageBody.(*message.SourceLineEvent).Log()
			break
		}
	}
}
