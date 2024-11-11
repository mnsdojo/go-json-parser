package parser

import (
	"fmt"
	"strconv"

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
	// initialize curent token
	if err := p.advanceToken(); err != nil {
		return nil, err
	}
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
	obj := make(map[string]interface{})

	// Advance past the opening brace
	if err := p.advanceToken(); err != nil {
		return nil, err
	}
	// loop till end of the object }
	for p.currentToken.Type != tokenizer.ObjectEnd {
		// Parse the key as a string
		key, err := p.parseString()
		if err != nil {
			return nil, fmt.Errorf("failed to parse object key: %w", err)
		}

		// Expect a colon after the key
		if p.currentToken.Type != tokenizer.Colon {
			return nil, fmt.Errorf("expected colon after object key, got %s", p.currentToken.Type)
		}
		if err := p.advanceToken(); err != nil {
			return nil, err
		}

		// Parse the value
		value, err := p.parseValues()
		if err != nil {
			return nil, fmt.Errorf("failed to parse object value for key '%s': %w", key, err)
		}
		obj[key] = value

		// Check for comma or end of object
		if p.currentToken.Type == tokenizer.Comma {
			if err := p.advanceToken(); err != nil {
				return nil, err
			}
		} else if p.currentToken.Type != tokenizer.ObjectEnd {
			return nil, fmt.Errorf("expected comma or end of object, got %s", p.currentToken.Type)
		}
	}

	// Advance past the closing brace
	if err := p.advanceToken(); err != nil {
		return nil, err
	}
	return obj, nil
}

func (p *Parser) parseArray() ([]interface{}, error) {
	var array []interface{}
	if err := p.advanceAndCheckToken(tokenizer.ArrayStart); err != nil {
		return nil, err
	}
	for p.currentToken.Type != tokenizer.ArrayEnd {
		// Parse the next value in the array
		element, err := p.parseValues()
		if err != nil {
			return nil, err
		}
		array = append(array, element)

		// Check for comma or end of array
		if p.currentToken.Type == tokenizer.Comma {
			if err := p.advanceToken(); err != nil {
				return nil, err
			}
		} else if p.currentToken.Type != tokenizer.ArrayEnd {
			return nil, fmt.Errorf("expected comma or end of array, got %s", p.currentToken.Type)
		}
	}

	// Advance past the closing bracket
	if err := p.advanceAndCheckToken(tokenizer.ArrayEnd); err != nil {
		return nil, err
	}
	return array, nil
}

func (p *Parser) parseBoolean() (bool, error) {
	switch p.currentToken.Value {
	case "true":
		if err := p.advanceToken(); err != nil {
			return false, err
		}
		return true, nil
	case "false":
		if err := p.advanceToken(); err != nil {
			return false, err
		}
		return false, nil
	}
	return false, fmt.Errorf("expected boolean token, got %s", p.currentToken.Value)
}

func (p *Parser) parseString() (string, error) {
	if err := p.checkTokenType(tokenizer.String); err != nil {
		return "", err
	}
	value := p.currentToken.Value
	if err := p.advanceToken(); err != nil {
		return "", err
	}
	return value, nil
}

func (p *Parser) parseNull() (interface{}, error) {
	if err := p.checkTokenType(tokenizer.Null); err != nil {
		return nil, err
	}
	if err := p.advanceToken(); err != nil {
		return nil, err
	}
	return nil, nil
}

func (p *Parser) parseNumber() (interface{}, error) {
	// Here you might want to parse the number properly into a float or int

	if err := p.checkTokenType(tokenizer.Number); err != nil {
		return nil, err
	}
	value, err := strconv.ParseFloat(p.currentToken.Value, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse number: %s", p.currentToken.Value)
	}
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

func (p *Parser) advanceAndCheckToken(expectedType tokenizer.TokenType) error {
	if err := p.advanceToken(); err != nil {
		return err
	}
	if p.currentToken.Type != expectedType {
		return fmt.Errorf("expected %s token, got %s", expectedType, p.currentToken.Type)
	}
	return nil
}

func (p *Parser) checkTokenType(expectedType tokenizer.TokenType) error {
	if p.currentToken.Type != expectedType {
		return fmt.Errorf("expected %s  token, got %s", expectedType, p.currentToken.Type)
	}
	return nil
}
