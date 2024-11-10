package tokenizer

import (
	"fmt"
	"unicode"
)

type Tokenizer struct {
	input       string
	position    int
	currentChar byte
}

func NewTokenizer(input string) *Tokenizer {
	return &Tokenizer{
		input:       input,
		position:    0,
		currentChar: input[0],
	}
}

// skipWhitespace skips any whitespace characters (spaces, tabs, newlines).
func (t *Tokenizer) skipWhitespace() {
	for t.currentChar == ' ' || t.currentChar == '\t' || t.currentChar == '\n' {
		t.moveNext()
	}
}

func (t *Tokenizer) GetNextToken() (*Token, error) {
	for t.position < len(t.input) {
		t.skipWhitespace()
		if t.position >= len(t.input) {
			return nil, nil
		}

		fmt.Printf("Position: %d, Current Char: %c\n", t.position, t.currentChar)

		// handle different token types
		switch t.currentChar {
		case '{':
			t.moveNext()
			return &Token{Value: "{", Type: ObjectStart}, nil
		case '}':
			t.moveNext()
			return &Token{Value: "}", Type: ObjectEnd}, nil
		case '[':
			t.moveNext()
			return &Token{Value: "[", Type: ArrayStart}, nil
		case ']':
			t.moveNext()
			return &Token{Value: "]", Type: ArrayEnd}, nil

		case ':':
			t.moveNext()
			return &Token{Value: ":", Type: Colon}, nil
		case ',':
			t.moveNext()
			return &Token{Value: ",", Type: Comma}, nil
		case 't', 'f':
			token, err := t.readBoolean()
			if err != nil {
				return nil, err
			}
			return token, nil
		case 'n':
			token, err := t.readNull()
			if err != nil {
				return nil, err
			}
			return token, nil
		case '"':
			token, err := t.readString()
			if err != nil {
				return nil, err
			}
			return token, nil
		case '-', '+', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9': // Numbers
			token, err := t.readNumber()
			if err != nil {
				return nil, err
			}
			return token, nil

		}

		t.moveNext()
	}

	return nil, nil
}

func (t *Tokenizer) readNumber() (*Token, error) {
	var numValue string

	if t.currentChar == '-' || t.currentChar == '+' {
		numValue += string(t.currentChar)
		t.moveNext()
	}

	// handling integers
	if !unicode.IsDigit(rune(t.currentChar)) {
		return nil, fmt.Errorf("invalid character '%c' found in number", t.currentChar)
	}

	for t.position < len(t.input) && unicode.IsDigit(rune(t.currentChar)) {
		numValue += string(t.currentChar)
		t.moveNext()
	}

	// handling fractional part of the number if found

	if t.currentChar == '.' {
		numValue += "."
		t.moveNext()
		if !unicode.IsDigit(rune(t.currentChar)) {
			return nil, fmt.Errorf("invalid fractional part in number")
		}
		for t.position < len(t.input) && unicode.IsDigit(rune(t.currentChar)) {
			numValue += string(t.currentChar)
			t.moveNext()
		}

	}
	return &Token{Value: numValue, Type: Number}, nil
}

func (t *Tokenizer) readString() (*Token, error) {
	var strValue string

	// Move past the opening quote.
	t.moveNext()

	for t.currentChar != '"' && t.position < len(t.input) {
		if t.currentChar == '\\' {
			// Move past the escape character.
			t.moveNext()
			// Handle different escape sequences.
			switch t.currentChar {
			case '"':
				strValue += "\""
			case '\\':
				strValue += "\\"
			case 'n':
				strValue += "\n"
			case 't':
				strValue += "\t"
			case 'r':
				strValue += "\r"
			default:
				// Handle invalid escape sequences.
				return nil, fmt.Errorf("invalid escape sequence \\%c", t.currentChar)
			}
			t.moveNext()
		} else {
			// Add regular character to string value.
			strValue += string(t.currentChar)
			t.moveNext()
		}
	}

	// If the string ends without a closing quote, handle it.
	if t.currentChar != '"' {
		return nil, fmt.Errorf("unterminated string literal")
	}

	// Move past the closing quote.
	t.moveNext()

	return &Token{Value: strValue, Type: String}, nil
}

func (t *Tokenizer) readNull() (*Token, error) {
	if !t.matchString("null") {
		return nil, fmt.Errorf("invalid null value")
	}
	return &Token{Value: "null", Type: Null}, nil
}

func (t *Tokenizer) readBoolean() (*Token, error) {
	if t.currentChar == 't' {
		if !t.matchString("true") {
			return nil, fmt.Errorf("invalid boolean value")
		}
		return &Token{Value: "true", Type: Boolean}, nil
	}
	if t.currentChar == 'f' {
		if !t.matchString("false") {
			return nil, fmt.Errorf("invalid boolean value")
		}
		return &Token{Value: "false", Type: Boolean}, nil

	}
	return nil, fmt.Errorf("invalid boolean value starting with %c", t.currentChar)
}

func (t *Tokenizer) moveNext() {
	t.position++
	if t.position < len(t.input) {
		t.currentChar = t.input[t.position]
	} else {
		t.currentChar = 0
	}
}

func (t *Tokenizer) matchString(str string) bool {
	for i := 0; i < len(str); i++ {
		if t.currentChar != str[i] {
			return false
		}
		t.moveNext()
	}
	return true
}
