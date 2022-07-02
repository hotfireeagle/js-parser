package out

import (
	"encoding/json"
	"jsj/parser"
)

type SaveTokenType struct {
	TypeName string `json:"type"`
	Value    string `json:"value"`
}

func SaveJsonTokens(tokens []*parser.Token) {
	saveTokens := make([]*SaveTokenType, 0)

	for _, tokenObj := range tokens {
		saveTokens = append(saveTokens, &SaveTokenType{
			TypeName: tokenObj.TokenTypeToString(),
			Value:    tokenObj.GetValue(),
		})
	}

	data, err := json.MarshalIndent(saveTokens, "", "  ")

	if err != nil {
		panic(err)
	}

	writeToFile("tokens-out.json", data)
}
