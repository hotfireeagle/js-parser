package intermediate

type SymTabEntry struct {
	name        string
	symTab      *SymTab
	lineNumbers []int
	symTabs     map[SymTabKey]interface{}
}

func SymTabEntryConstructor(name string, symTab *SymTab) *SymTabEntry {
	return &SymTabEntry{
		name:        name,
		symTab:      symTab,
		lineNumbers: make([]int, 0),
		symTabs:     make(map[SymTabKey]interface{}),
	}
}

func (ste *SymTabEntry) GetName() string {
	return ste.name
}

func (ste *SymTabEntry) GetSymTab() *SymTab {
	return ste.symTab
}

func (ste *SymTabEntry) AppendLineNumber(lineNumber int) {
	ste.lineNumbers = append(ste.lineNumbers, lineNumber)
}

func (ste *SymTabEntry) GetLineNumbers() []int {
	return ste.lineNumbers
}

func (ste *SymTabEntry) SetAttribute(key SymTabKey, value interface{}) {
	ste.symTabs[key] = value
}

func (ste *SymTabEntry) GetAttribute(key SymTabKey) interface{} {
	return ste.symTabs[key]
}
