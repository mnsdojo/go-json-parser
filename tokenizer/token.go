package tokenizer

// tokentype represents the type of token
type TokenType int

const (
	ObjectStart TokenType = iota
	ObjectEnd
	ArrayStart
	ArrayEnd
	String
	Number
	Null
)

// token represets single tokn which holds type vand value --
type Token struct {
	Value string
	Type  TokenType
}

func (t TokenType) String() string {
	switch t {
	case ObjectStart:
		return "ObjectStart"
	case ObjectEnd:
		return "ObjectEnd"
	case ArrayStart:
		return "ArrayStart"
	case String:
		return "String"
	case Number:
		return "number"
	case Null:
		return "Null"
	default:
		return "Unknown"
	}
}
