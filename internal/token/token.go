package token

import "fmt"

type Token struct {
	Type       TokenType
	Lexeme  string
	Literal map[string]int
	Line    int
}

func (t Token) String() string {
	if t.Literal == nil {
		return  fmt.Sprintf("%v %v %v", t.Type, t.Lexeme, "null")

	}
	return fmt.Sprintf("%v %v %v", t.Type, t.Lexeme, t.Literal)
}
