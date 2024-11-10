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
	// Skip over whitespace characters
	for t.currentChar == ' ' || t.currentChar == '\t' || t.currentChar == '\n' {
		t.moveNext()
	}
}

func (t *Tokenizer) GetNextToken() (*Token, error) {
	for t.position < len(t.input) {
		t.skipWhitespace()
		fmt.Printf("Position: %d, Current Char: %c\n", t.position, t.currentChar)
		t.moveNext()
	}

	return nil, nil
}

func (t *Tokenizer) moveNext() {
	t.position++
	if t.position < len(t.input) {
		t.currentChar = t.input[t.position]
	} else {
		t.currentChar = 0
	}
}
