package lexer

import (
	"lexing/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	rn           rune // support UTF-8
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readRune()
	return l
}

func (l *Lexer) readRune() {
	inputRunes := []rune(l.input)
	if l.readPosition >= len(inputRunes) {
		l.rn = 0
	} else {
		l.rn = inputRunes[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) NextToken() *token.Token {
	tok := &token.Token{}
	if tokenType, exist := token.ConstLiteral2TokenType[string(l.rn)]; exist {
		tok = &token.Token{
			Type:    tokenType,
			Literal: string(l.rn),
		}
	} else {
		tok = &token.Token{
			Type: token.EOF,
			Literal: "",
		}
	}
	l.readRune()
	return tok
}