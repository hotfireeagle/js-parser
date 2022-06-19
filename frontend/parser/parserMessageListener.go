package parser

import "jsj/message"

type ParserMessageListener struct {
}

func (pml *ParserMessageListener) MessageReceived(messageObj message.Message) {
	messageType := messageObj.GetMessageType()
	messageBody := messageObj.GetBody()

	switch messageType {
	case message.PARSER_SUMMARY:
		{
			messageBody.(*message.ParserSummaryEvent).Log()
			break
		}
	case message.TOKEN:
		{
			messageBody.(*TokenEvent).Log()
			break
		}
	case message.SYNTAX_ERROR:
		{
			messageBody.(*SyntaxErrorEvent).Log()
			break
		}
	}
}
