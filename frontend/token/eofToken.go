package token

import "jsj/frontend/source"

type EofToken struct {
	Token
}

func EofTokenConstructor(s source.Source) *EofToken {
	tokenObj := &EofToken{
		Token: TokenConstructor(s),
	}

	return tokenObj
}

// because EOF means the end of file
// so there are nothing to extract
func (t *EofToken) Extract() {

}
