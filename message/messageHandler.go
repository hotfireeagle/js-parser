package message

type MessageHandler struct {
	message   Message
	listeners []MessageListener
}

func MessageHandlerConstructor() MessageHandler {
	return MessageHandler{
		listeners: make([]MessageListener, 0),
	}
}

// add a message listener
func (messageHandlerInstance *MessageHandler) AddListener(listener MessageListener) {
	messageHandlerInstance.listeners = append(messageHandlerInstance.listeners, listener)
}

// remove a message listener
func (messageHandlerInstance *MessageHandler) RemoveListener(listener MessageListener) {
	// Notation: O(n)
	// is necessary to add map to check, so that we can arrive at O(1)
	for i, l := range messageHandlerInstance.listeners {
		// TODO: 确定可不可以这样进行比较，不能的话加ID
		if l == listener {
			messageHandlerInstance.listeners = append(messageHandlerInstance.listeners[:i], messageHandlerInstance.listeners[i+1:]...)
			return
		}
	}
}

// send a message to all listeners
func (messageHandlerInstance *MessageHandler) SendMessage(message Message) {
	messageHandlerInstance.message = message
	messageHandlerInstance.notifyListeners()
}

// notify all listeners
func (messageHandlerInstance *MessageHandler) notifyListeners() {
	for _, listener := range messageHandlerInstance.listeners {
		listener.MessageReceived(messageHandlerInstance.message)
	}
}
