package scanner

import (
	"fmt"
	"strconv"
	"unicode"

	loxinterpreter "github.com/codecrafters-io/interpreter-starter-go/internal/lox-interpreter"
	"github.com/codecrafters-io/interpreter-starter-go/internal/token"
)

var keywords map[string]token.TokenType = map[string]token.TokenType{
	"and":    token.AND,
	"class":  token.CLASS,
	"else":   token.ELSE,
	"false":  token.FALSE,
	"for":    token.FOR,
	"fun":    token.FUN,
	"if":     token.IF,
	"nil":    token.NIL,
	"or":     token.OR,
	"print":  token.PRINT,
	"return": token.RETURN,
	"super":  token.SUPER,
	"this":   token.THIS,
	"true":   token.TRUE,
	"var":    token.VAR,
	"while":  token.WHILE,
}

func isLetter(ch rune) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

type Scanner struct {
	source []rune
	tokens []token.Token

	start   int
	current int
	line    int
}

func (s Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) advance() rune {
	r := s.source[s.current]
	s.current++
	return r
}

func (s *Scanner) match(expexted rune) bool {
	if s.isAtEnd() || s.source[s.current] != expexted {
		return false
	}

	s.current++
	return true
}

func (s Scanner) peek() rune {
	if s.isAtEnd() {
		return '\000'
	}
	return s.source[s.current]
}

func (s Scanner) peekNext() rune {
	if s.current+1 >= len(s.source) {
		return '\000'
	}
	return s.source[s.current+1]
}

func (s *Scanner) addToken(t token.TokenType, literal interface{}) {
	text := s.source[s.start:s.current]
	s.tokens = append(s.tokens, token.Token{Type: t, Lexeme: string(text), Literal: literal, Line: s.line})
}

func (s *Scanner) scanToken() {
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
	case '!':
		if s.match('=') {
			s.addToken(token.BANG_EQUAL, nil)
		} else {
			s.addToken(token.BANG, nil)
		}
	case '=':
		if s.match('=') {
			s.addToken(token.EQUAL_EQUAL, nil)
		} else {
			s.addToken(token.EQUAL, nil)
		}

	case '<':
		if s.match('=') {
			s.addToken(token.LESS_EQUAL, nil)
		} else {
			s.addToken(token.LESS, nil)
		}
	case '>':
		if s.match('=') {
			s.addToken(token.GREATER_EQUAL, nil)
		} else {
			s.addToken(token.GREATER, nil)
		}
	case '/':
		if s.match('/') {
			for s.peek() != '\n' && !s.isAtEnd() {
				s.advance()
			}
		} else {
			s.addToken(token.SLASH, nil)
		}
	case '"':
		for s.peek() != '"' && !s.isAtEnd() {
			if s.peek() == '\n' {
				s.line++
			}
			s.advance()
		}

		if s.isAtEnd() {
			loxinterpreter.Error(s.line, "Unterminated string.")
			return
		}

		s.advance()
		s.addToken(token.STRING, string(s.source[s.start+1:s.current-1]))
	case ' ':
	case '\r':
	case '\t':
		break
	case '\n':
		s.line++
	default:
		if unicode.IsDigit(c) {
			for unicode.IsDigit(s.peek()) {
				s.advance()
			}

			if s.peek() == '.' && unicode.IsDigit(s.peekNext()) {
				s.advance()
				for unicode.IsDigit(s.peek()) {
					s.advance()
				}
			}

			literal, err := strconv.ParseFloat(string(s.source[s.start:s.current]), 64)
			if err != nil {
				loxinterpreter.Error(s.line, "Unable to parse number")
			} else {
				s.addToken(token.NUMBER, literal)
			}
		} else if isLetter(c){ 
			for isLetter(s.peek()) || unicode.IsDigit(s.peek()) {
				s.advance()
			}

			text := string(s.source[s.start:s.current])
			t, ok := keywords[text]
			if !ok {
				t = token.IDENTIFIER
			}

			s.addToken(t, nil)
		} else {
			loxinterpreter.Error(s.line, fmt.Sprintf("Unexpected character: %c", c))
		}
	}
}

func (s *Scanner) ScanTokens() []token.Token {
	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()
	}

	s.tokens = append(s.tokens, token.Token{Type: token.EOF, Lexeme: "", Literal: nil, Line: s.line})
	return s.tokens
}

func NewScanner(source []rune) *Scanner {
	return &Scanner{
		source,
		nil,
		0, 0, 1,
	}
}
