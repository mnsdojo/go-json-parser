package tokenizer

import "fmt"

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
		case '"':
			token, err := t.readString()
			if err != nil {
				return nil, err
			}
			return token, nil
		}

		t.moveNext()
	}

	return nil, nil
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

func (t *Tokenizer) moveNext() {
	t.position++
	if t.position < len(t.input) {
		t.currentChar = t.input[t.position]
	} else {
		t.currentChar = 0
	}
}
