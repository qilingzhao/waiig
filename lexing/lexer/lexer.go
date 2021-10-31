package lexer

import (
	`unicode`

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
	for ; unicode.IsSpace(l.rn); l.readRune() {

	}

	// log.Printf("rune is %c", l.rn)
	if tokenType, exist := token.SymbolLiteral2TokenType[string(l.rn)]; exist {
		literal := string(l.rn)
		if tokenType == token.EOF {
			literal = ""
		}
		tok = &token.Token{
			Type:    tokenType,
			Literal: literal,
		}
	}  else if unicode.IsLetter(l.rn) { // 变量名和关键字以字母开头
		ident := l.readIdentifier()
		identType := token.LookupIdent(ident)
		tok = &token.Token{
			Type:    identType,
			Literal: ident,
		}
		return tok
	} else if unicode.IsDigit(l.rn) {
		num := l.readNumber()
		tok = &token.Token{
			Type:    token.INT,
			Literal: num,
		}
		return tok
	} else {
		tok = &token.Token{
			Type:    token.ILLEGAL,
			Literal: string(l.rn),
		}
	}
	l.readRune()
	return tok
}

func (l *Lexer) readIdentifier() string {
	start := l.position
	for ; unicode.IsLetter(l.rn) || l.rn == '_'; l.readRune() {

	}
	rns := []rune(l.input)
	return string(rns[start:l.position])
}

func (l *Lexer) readNumber() string {
	start := l.position
	for ; unicode.IsNumber(l.rn); l.readRune() {

	}
	rns := []rune(l.input)
	return string(rns[start:l.position])
}