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

func (m *Message) GetMessageType() MessageType {
	return m.messageType
}

func (m *Message) GetBody() interface{} {
	return m.body
}
