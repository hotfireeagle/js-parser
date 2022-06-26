package intermediate

import "sort"

type SymTab struct {
	nestingLevel int
	entrys       map[string]*SymTabEntry
}

func SymTabConstructor(nestingLevel int) *SymTab {
	return &SymTab{
		nestingLevel: nestingLevel,
		entrys:       make(map[string]*SymTabEntry),
	}
}

func (st *SymTab) GetNestingLevel() int {
	return st.nestingLevel
}

func (st *SymTab) Enter(name string) *SymTabEntry {
	entry := SymTabEntryConstructor()
	st.entrys[name] = entry
	return entry
}

func (st *SymTab) Lookup(name string) *SymTabEntry {
	return st.entrys[name]
}

// 根据name升序排序SymTabEntry
func (st *SymTab) SortedEntries() []*SymTabEntry {
	names := make([]string, 0)

	for k := range st.entrys {
		names = append(names, k)
	}

	sort.Strings(names)

	result := make([]*SymTabEntry, 0)

	for _, v := range names {
		result = append(result, st.entrys[v])
	}

	return result
}
