package parser

import "github.com/mnsdojo/go-json-parser/tokenizer"

type Parser struct {
	tokenizer    *tokenizer.Tokenizer
	currentToken *tokenizer.Token
}

func NewParser(tokenizer *tokenizer.Tokenizer) *Parser {
	return &Parser{
		tokenizer: tokenizer,
	}
}

func (p *Parser) Parse() (interface{}, error) {
	return nil, nil
}
