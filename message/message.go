package message

type Message struct {
	body        interface{}
	messageType MessageType
}

func MessageConstructor(t MessageType, b interface{}) Message {
	return Message{
		body:        b,
		messageType: t,
	}
}
