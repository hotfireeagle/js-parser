package intermediate

type ICode struct{}

func IcodeConstructor() *ICode {
	return &ICode{}
}

func (ic *ICode) SetRoot(node *ICodeNode) *ICodeNode {
	return nil
}

func (ic *ICode) GetRoot() *ICodeNode {
	return nil
}
