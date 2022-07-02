package intermediate

type ICodeNodeType int

const (
	// structure
	FUNCTION ICodeNodeType = iota

	// statement
	COMPOUND
	ASSIGN
	LOOP
	IF

	EQ
	NE
	LT
	LE
	GT
	GE
	NOT

	ADD
	OR
	NEGATE
	SUBTRACT

	MULTIPLY
	MOD
	AND
	INTEGER_DIVIDE
	FLOAT_DIVIDE

	VARIABLE
	INTEGER_CONSTANT
	REAL_CONSTANT
	STRING_CONSTANT
	BOOLEAN_CONSTANT
)

func (icnt *ICodeNodeType) ToString() string {
	return ""
}
