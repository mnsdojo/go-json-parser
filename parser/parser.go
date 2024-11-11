package parser

import (
	"fmt"

	"github.com/mnsdojo/go-json-parser/tokenizer"
)

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
	return p.parseValues()
}

func (p *Parser) parseValues() (interface{}, error) {
	switch p.currentToken.Type {
	case tokenizer.ObjectStart:
		return p.parseObject()
	case tokenizer.ArrayStart:
		return p.parseArray()
	case tokenizer.String:
		return p.parseString()
	case tokenizer.Number:
		return p.parseNumber()
	case tokenizer.Boolean:
		return p.parseBoolean()
	case tokenizer.Null:
		return p.parseNull()
	default:
		return nil, fmt.Errorf("unexpected error : %s", p.currentToken.Type.String())
	}
}

func (p *Parser) parseObject() (map[string]interface{}, error) {
}

func (p *Parser) parseBoolean() (bool, error) {
	if p.currentToken.Value == "true" {
		if err := p.advanceToken(); err != nil {
			return false, err
		}
		return true, nil
	} else if p.currentToken.Value == "false" {
		if err := p.advanceToken(); err != nil {
			return false, err
		}
		return false, nil
	}
	return false, fmt.Errorf("expected boolean token, got %s", p.currentToken.Value)
}

func (p *Parser) parseString() (string, error) {
	if p.currentToken.Type != tokenizer.String {
		return "", fmt.Errorf("expected string token, got %s", p.currentToken.Type)
	}
	value := p.currentToken.Value
	if err := p.advanceToken(); err != nil {
		return "", err
	}
	return value, nil
}

func (p *Parser) parseNull() (interface{}, error) {
	if p.currentToken.Type != tokenizer.Null {
		return nil, fmt.Errorf("expected null token, got %s", p.currentToken.Type)
	}
	if err := p.advanceToken(); err != nil {
		return nil, err
	}
	return nil, nil
}

func (p *Parser) parseNumber() (interface{}, error) {
	// Here you might want to parse the number properly into a float or int
	if p.currentToken.Type != tokenizer.Number {
		return nil, fmt.Errorf("expected number token, got %s", p.currentToken.Type)
	}
	value := p.currentToken.Value
	if err := p.advanceToken(); err != nil {
		return nil, err
	}
	return value, nil
}

func (p *Parser) advanceToken() error {
	token, err := p.tokenizer.GetNextToken()
	if err != nil {
		return fmt.Errorf("failed to get next token ;%w", err)
	}
	p.currentToken = token
	return nil
}
