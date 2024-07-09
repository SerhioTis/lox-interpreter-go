package scanner

import (
	"fmt"

	"github.com/codecrafters-io/interpreter-starter-go/internal/token"
)

type Scanner struct {
	source string
	tokens []token.Token

	start   int
	current int
	line    int
}

func (s Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) advance() byte {
	r := s.source[s.current]
	s.current++
	return r
}

func (s *Scanner) match(expexted byte) bool {
	if s.isAtEnd() || s.source[s.current] != expexted {
		return false
	}

	s.current++
	return true
}

func (s *Scanner) addToken(t token.TokenType, literal map[string]int) {
	text := s.source[s.start:s.current]
	s.tokens = append(s.tokens, token.Token{Type: t, Lexeme: text, Literal: literal, Line: s.line})
}

func (s *Scanner) scanToken() error {
	c := s.advance()
	switch c {
	case '(':
		s.addToken(token.LEFT_PAREN, nil)
	case ')':
		s.addToken(token.RIGHT_PAREN, nil)
	case '{':
		s.addToken(token.LEFT_BRACE, nil)
	case '}':
		s.addToken(token.RIGHT_BRACE, nil)
	case ',':
		s.addToken(token.COMMA, nil)
	case '.':
		s.addToken(token.DOT, nil)
	case '-':
		s.addToken(token.MINUS, nil)
	case '+':
		s.addToken(token.PLUS, nil)
	case ';':
		s.addToken(token.SEMICOLON, nil)
	case '*':
		s.addToken(token.STAR, nil)
	case '\n':
		s.line++
	default:
		return fmt.Errorf("%v Unexpected character.", s.line)
	}
	return nil
}

func (s *Scanner) ScanTokens() ([]token.Token, error) {
	for !s.isAtEnd() {
		s.start = s.current
		err := s.scanToken()
		if err != nil {
			return s.tokens, err
		}
	}

	s.tokens = append(s.tokens, token.Token{Type: token.EOF, Lexeme: "", Literal: nil, Line: s.line})
	return s.tokens, nil
}

func NewScanner(source string) *Scanner {
	return &Scanner{
		source,
		nil,
		0, 0, 1,
	}
}
