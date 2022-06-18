package message

type MessageListener interface {
	// called to receive a message sent by message producer
	MessageReceived(message Message)
}
