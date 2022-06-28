package intermediate

type ICodeNode struct {
	nodeType   ICodeNodeType
	parent     *ICodeNode
	children   []*ICodeNode
	attributes map[ICodeKey]interface{}
}

func ICodeNodeConstructor(t ICodeNodeType) *ICodeNode {
	return &ICodeNode{
		nodeType:   t,
		parent:     nil,
		children:   make([]*ICodeNode, 0),
		attributes: make(map[ICodeKey]interface{}),
	}
}

func (icn *ICodeNode) GetType() ICodeNodeType {
	return icn.nodeType
}

// 返回当前节点的父节点
func (icn *ICodeNode) GetParent() *ICodeNode {
	return icn.parent
}

// 给icn node添加子节点
func (icn *ICodeNode) AddChild(node *ICodeNode) *ICodeNode {
	if node != nil {
		icn.children = append(icn.children, node)
		node.parent = icn
	}
	return node
}

// 获取子节点列表
func (icn *ICodeNode) GetChildren() []*ICodeNode {
	return icn.children
}

func (icn *ICodeNode) SetAttribute(k ICodeKey, v interface{}) {
	icn.attributes[k] = v
}

func (icn *ICodeNode) GetAttribute(k ICodeKey) interface{} {
	return icn.attributes[k]
}

func (icn *ICodeNode) Copy() *ICodeNode {
	copyIns := ICodeNodeConstructor(icn.nodeType)

	for k, v := range icn.attributes {
		copyIns.SetAttribute(k, v)
	}

	return copyIns
}

func (icn *ICodeNode) ToString() string {
	return icn.nodeType.ToString()
}
