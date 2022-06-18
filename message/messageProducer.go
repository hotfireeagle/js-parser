package message

// TODO: 检查在go代码模式中用不到的话，删掉
type MessageProducer interface {
	AddMessageListener(listener MessageListener)

	RemoveMessageListener(listener MessageListener)

	SendMessage(message Message)
}
