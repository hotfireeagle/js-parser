package backend

import (
	"jsj/intermediate"
	"jsj/message"
)

var messageHandler = message.MessageHandlerConstructor()

type Backend struct {
	symTab *intermediate.SymTab
	icode  *intermediate.ICode
}

func (backendInstance *Backend) Process(icode *intermediate.ICode, symTab *intermediate.SymTab) {

}
