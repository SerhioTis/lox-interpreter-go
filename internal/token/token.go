package token

import "fmt"

type Token struct {
	Type    TokenType
	Lexeme  string
	Literal interface{}
	Line    int
}

func (t Token) String() string {
	literal := t.Literal
	if t.Literal == nil {
		literal = "null"

	} else if v, ok := t.Literal.(float64); ok {
		if v == float64(int(v)) {
			literal = fmt.Sprintf("%.1f", v)
		} else {
			literal = fmt.Sprintf("%g", v)
		}

	}
	return fmt.Sprintf("%s %s %s", t.Type, t.Lexeme, literal)
}
