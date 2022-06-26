package intermediate

type SymTabStack struct {
}

func (sts *SymTabStack) SymTabStackConstructor() *SymTabStack {
	return &SymTabStack{}
}

func (sts *SymTabStack) GetCurrentNestingLevel() int {
	return 0
}

// localSymbolTable就是symbol table stack中的栈顶元素
func (sts *SymTabStack) GetLocalSymbolTable() *SymTab {
	return nil
}

func (sts *SymTabStack) EnterLocal(name string) *SymTabEntry {
	return nil
}

func (sts *SymTabStack) LookUpLocal(name string) *SymTabEntry {
	return nil
}

func (sts *SymTabStack) LookUp(name string) *SymTabEntry {
	return nil
}
