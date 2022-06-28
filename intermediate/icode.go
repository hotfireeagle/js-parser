package intermediate

type ICode struct {
	root *ICodeNode
}

func IcodeConstructor() *ICode {
	return &ICode{}
}

func (ic *ICode) SetRoot(node *ICodeNode) *ICodeNode {
	ic.root = node
	return ic.root
}

func (ic *ICode) GetRoot() *ICodeNode {
	return ic.root
}
