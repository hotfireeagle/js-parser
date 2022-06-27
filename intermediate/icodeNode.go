package intermediate

type ICodeNode struct {
}

func ICodeNodeConstructor() *ICodeNode {
	return &ICodeNode{}
}

func (icn *ICodeNode) GetType() ICodeNodeType {
	return 0
}

// 返回当前节点的父节点
func (icn *ICodeNode) GetParent() *ICodeNode {
	return nil
}

// 新增子节点
func (icn *ICodeNode) AddChild(node *ICodeNode) *ICodeNode {
	return nil
}

// 获取子节点列表
func (icn *ICodeNode) GetChildren() []*ICodeNode {
	return nil
}

func (icn *ICodeNode) SetAttribute(k ICodeKey, v interface{}) {

}

func (icn *ICodeNode) GetAttribute(k ICodeKey) interface{} {
	return nil
}

func (icn *ICodeNode) Copy() *ICodeNode {
	return nil
}
