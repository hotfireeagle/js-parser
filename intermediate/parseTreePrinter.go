package intermediate

import (
	"fmt"
	"strings"
)

var INDENT_WIDTH_2 = 4
var LINE_WIDTH = 80

type ParseTreePrinter struct {
	length      int
	indent      string
	indentation string
	line        strings.Builder
}

func ParseTreePrinterConstructor() *ParseTreePrinter {
	s := ""

	for i := 0; i < INDENT_WIDTH_2; i++ {
		s += " "
	}

	return &ParseTreePrinter{
		indent: s,
	}
}

func (ptp *ParseTreePrinter) Print(icode *ICode) {
	fmt.Println("\n ========== INTERMEDIATE CODE ==========\n")
	ptp.PrintNode(icode.GetRoot())
	ptp.PrintLine()
}

func (ptp *ParseTreePrinter) PrintNode(node *ICodeNode) {
	ptp.Append(ptp.indentation)
	ptp.Append("<" + node.ToString())

	ptp.PrintAttributes(node)

	ptp.PrintTypeSpec(node)

	childNodes := node.GetChildren()

	if len(childNodes) > 0 {
		ptp.Append(">")
		ptp.PrintLine()

		ptp.PrintChildNodes(childNodes)

		ptp.Append(ptp.indentation)
		ptp.Append("</" + node.ToString() + ">")
	} else {
		ptp.Append(" />")
	}

	ptp.PrintLine()
}

func (ptp *ParseTreePrinter) PrintAttributes(node *ICodeNode) {
	saveIndentation := ptp.indentation
	ptp.indentation += ptp.indent

	for _, attribute := range node.GetAttributes() {
		ptp.PrintAttribute(attribute)
	}

	ptp.indentation = saveIndentation
}

// TODO:
func (ptp *ParseTreePrinter) PrintAttribute(attribute interface{}) {

}

func (ptp *ParseTreePrinter) PrintChildNodes(childNodes []*ICodeNode) {
	saveIndentation := ptp.indentation
	ptp.indentation += ptp.indent
	for _, childNode := range childNodes {
		ptp.PrintNode(childNode)
	}
	ptp.indentation = saveIndentation
}

// TODO:
func (ptp *ParseTreePrinter) PrintTypeSpec(node *ICodeNode) {

}

func (ptp *ParseTreePrinter) Append(text string) {
	textLength := len(text)

	lineBreak := false

	if ptp.length+textLength > LINE_WIDTH {
		ptp.PrintLine()
		ptp.line.WriteString(ptp.indentation)
		ptp.length = len(ptp.indentation)
		lineBreak = true
	}

	if !(lineBreak && text == " ") {
		ptp.line.WriteString(text)
		ptp.length += textLength
	}
}

func (ptp *ParseTreePrinter) PrintLine() {
	if ptp.length > 0 {
		fmt.Println(ptp.line.String())
		ptp.line.Reset()
		ptp.length = 0
	}
}
