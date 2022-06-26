package intermediate

type SymTabEntry struct {
}

func SymTabEntryConstructor() *SymTabEntry {
	return &SymTabEntry{}
}

func (ste *SymTabEntry) GetName() string {
	return ""
}

func (ste *SymTabEntry) GetSymTab() *SymTab {
	return nil
}

func (ste *SymTabEntry) AppendLineNumber(lineNumber int) {

}

func (ste *SymTabEntry) GetLineNumbers() []int {
	return nil
}

func (ste *SymTabEntry) SetAttribute(key SymTabKey, value interface{}) {

}

func (ste *SymTabEntry) GetAttribute(key SymTabKey) interface{} {
	return nil
}
