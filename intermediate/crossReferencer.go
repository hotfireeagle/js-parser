package intermediate

import (
	"fmt"
	"strconv"
	"strings"
)

var NAME_WIDTH = 16
var NAME_FORMAT = "%-" + strconv.Itoa(NAME_WIDTH) + "s"
var NUMBERS_LABEL = " Line numbers "
var NUMBERS_UNDERLINE = " ---------- "
var NUMBER_FORMAT = " %03d"
var LABEL_WIDTH = len(NUMBERS_LABEL)
var INDENT_WIDTH = NAME_WIDTH + LABEL_WIDTH

var INDENT strings.Builder

func init() {
	for i := 0; i < INDENT_WIDTH; i++ {
		INDENT.WriteString(" ")
	}
}

type CrossReferencer struct {
}

func CrossReferencerConstructor() *CrossReferencer {
	return &CrossReferencer{}
}

func (cr *CrossReferencer) Print(sts *SymTabStack) {
	fmt.Println("\n===== CROSS-REFERENCE TABLE =====")
	cr.PrintColumnHeadings()
	cr.PrintSymTab(sts.GetLocalSymTab())
}

func (cr *CrossReferencer) PrintColumnHeadings() {
	fmt.Println()
	str := fmt.Sprintf(NAME_FORMAT, "Identifier")
	fmt.Println(str + NUMBERS_LABEL)
	s2 := fmt.Sprintf(NAME_FORMAT, "---------")
	fmt.Println(s2 + NUMBERS_UNDERLINE)
}

func (cr *CrossReferencer) PrintSymTab(st *SymTab) {
	sorted := st.SortedEntries()

	for _, entry := range sorted {
		lineNumbers := entry.GetLineNumbers()

		fmt.Printf(NAME_FORMAT, entry.GetName())

		for _, lineNum := range lineNumbers {
			fmt.Printf(NUMBER_FORMAT, lineNum)
		}

		fmt.Println()
	}
}
