package intermediate

type SymTabKey byte

const (
	// constant
	CONSTANT_VALUE SymTabKey = iota

	// function
	ROUTINE_CODE
	ROUTINE_SYMTAB
	ROUTINE_ICODE
	ROUTINE_PARMS
	ROUTINE_ROUTINES

	// variable
	DATA_VALUE
)
