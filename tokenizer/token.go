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
	Colon
	Comma
	Boolean
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
	case Colon:
		return "Colon"
	case Comma:
		return "Comma"
	case Boolean:
		return "Boolean"
	case Null:
		return "Null"
	default:
		return "Unknown"
	}
}
