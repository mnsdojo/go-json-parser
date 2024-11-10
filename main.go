package main

import (
	"fmt"

	"github.com/mnsdojo/go-json-parser/tokenizer"
)

func main() {
	jsonInput := `{
        "name": "Alice",
        "age": 30,
        "isStudent": false,
        "grades": [95, 88, 92],
        "address": {
            "city": "Wonderland",
            "zipcode": null
        }
    }`

	tokenizer := tokenizer.NewTokenizer(jsonInput)
	for {
		token, err := tokenizer.GetNextToken()
		if err != nil {
			fmt.Println("error", err)
			return
		}
		if token == nil {
			break
		}
		fmt.Printf("Token Type: %s, Value: %s\n", token.Type.String(), token.Value)
	}
}
