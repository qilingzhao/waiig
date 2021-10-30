package lexer

import (
	`lexing/token`
)

type Lexer struct {
	tokens []*token.Token
	index int
}

func New(input string) *Lexer {
	tokens := make([]*token.Token ,0)
	for _, field := range input {
		if tokenType, exist := token.ConstLiteral2TokenType[string(field)]; exist {
			tokens = append(tokens, &token.Token{
				Type:    tokenType,
				Literal: string(field),
			})
		}
	}


	return &Lexer{
		tokens: append(tokens, &token.Token{
			Type:    token.EOF,
			Literal: "",
		}),
		index: 0,
	}
}

func (l *Lexer) NextToken() *token.Token {
	now := l.index
	if now  < len(l.tokens) {
		l.index++
		return  l.tokens[now]
	}
	return nil
}