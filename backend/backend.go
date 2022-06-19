package backend

import (
	"jsj/intermediate"
	"jsj/message"
)

type Backend interface {
	Process(i *intermediate.ICode, s *intermediate.SymTab)
	AddMessageListener(listener message.MessageListener)
}
