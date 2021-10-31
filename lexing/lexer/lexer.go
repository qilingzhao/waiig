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
		tok = &token.Token{
			Type:    tokenType,
			Literal: string(l.rn),
		}
	} else if l.rn == 0 {
		tok = &token.Token{
			Type: token.EOF,
			Literal: "",
		}
	} else if !unicode.IsDigit(l.rn) { // 变量名和关键字不以数字开头
		word := ""
		for ; unicode.IsLetter(l.rn) || unicode.IsDigit(l.rn) || l.rn == '_'; l.readRune() {
			word = word + string(l.rn)
		}
		if keywordTokenType, existKeyword := token.KeywordLiteral2TokenType[word]; existKeyword {
			tok = &token.Token{
				Type:    keywordTokenType,
				Literal: word,
			}
		} else if word != "" {
			tok = &token.Token{
				Type:    token.IDENT,
				Literal: word,
			}
		} else {
			tok = &token.Token{
				Type:  token.ILLEGAL,
				Literal: word,
			}
			l.readRune()
		}
		return tok
	} else if unicode.IsDigit(l.rn) {
		word := ""
		for ; unicode.IsDigit(l.rn); l.readRune() {
			word = word + string(l.rn)
		}
		tok = &token.Token{
			Type:    token.INT,
			Literal: word,
		}
		return tok
	}
	l.readRune()
	return tok
}
