package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	text := `este producto es de origen argentino y nos encanta`
	//text := `18 gramos`
	lexer := newLexer(text)
	yyParse(lexer)
}

type lexer struct {
	input string
	pos   int
	eof   bool
}

func newLexer(input string) *lexer {
	return &lexer{input: input}
}

func (l *lexer) Lex(lval *yySymType) int {
	for l.pos < len(l.input) {
		switch l.input[l.pos] {
		case ' ', '\t', '\n', '\r':
			l.pos++
			continue
		case ':':
			l.pos++
			return int(':')
		}
		break
	}

	if l.pos >= len(l.input) {
		if l.eof {
			return 0
		}
		l.eof = true
		return EOF
	}

	start := l.pos
	for l.pos < len(l.input) && !strings.ContainsRune(" \t\n\r:", rune(l.input[l.pos])) {
		l.pos++
	}

	word := l.input[start:l.pos]

	if _, err := strconv.Atoi(word); err == nil {
		lval.word = word
		return NUMBER
	}
	switch strings.ToLower(word) {
	case "origen":
		return ORIGIN
	case "peso":
		return WEIGHT
	case "color":
		return COLOR
	case "radio":
		return RADIUS
	case "precio":
		return PRICE
	default:
		lval.word = word
		return WORD
	}
}

func (l *lexer) Error(s string) {
	fmt.Printf("syntax error: %s\n", s)
}
