package intermediate

type SymTabStack struct {
	currentNestingLevel int
	// 使用slice作为底层来实现这个栈
	symTabs []*SymTab
}

func SymTabStackConstructor() *SymTabStack {
	symTabStackIns := SymTabStack{
		currentNestingLevel: 0,
		symTabs:             make([]*SymTab, 0),
	}

	symTabStackIns.Add(SymTabConstructor(symTabStackIns.currentNestingLevel))

	return &symTabStackIns
}

func (sts *SymTabStack) GetCurrentNestingLevel() int {
	return sts.currentNestingLevel
}

// localSymbolTable就是symbol table stack中的栈顶元素
func (sts *SymTabStack) GetLocalSymTab() *SymTab {
	return sts.symTabs[sts.currentNestingLevel]
}

func (sts *SymTabStack) EnterLocal(name string) *SymTabEntry {
	return sts.symTabs[sts.currentNestingLevel].Enter(name)
}

func (sts *SymTabStack) LookUpLocal(name string) *SymTabEntry {
	return sts.symTabs[sts.currentNestingLevel].Lookup(name)
}

func (sts *SymTabStack) LookUp(name string) *SymTabEntry {
	return sts.LookUpLocal(name)
}

func (sts *SymTabStack) Add(symTab *SymTab) {
	sts.symTabs = append(sts.symTabs, symTab)
}
