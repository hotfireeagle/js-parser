package intermediate

type SymTab struct {
}

func SymTabConstructor() *SymTab {
	return &SymTab{}
}

func (st *SymTab) GetNestingLevel() int {
	return 0
}

func (st *SymTab) Enter(name string) *SymTabEntry {
	return nil
}

func (st *SymTab) Lookup(name string) *SymTabEntry {
	return nil
}

func (st *SymTab) SortedEntries() []*SymTabEntry {
	return nil
}
