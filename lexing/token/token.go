package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT" // x y x
	INT   = "INT"   // 1 2 3

	ASSIGN = "="
	PLUS   = "+"

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	FUNCTION = "FUNCTION"
	LET      = "LET"
)

var SymbolLiteral2TokenType = map[string]TokenType {
	"=": ASSIGN, "+": PLUS, ",": COMMA, ";": SEMICOLON,
	"(": LPAREN, ")": RPAREN, "{": LBRACE, "}": RBRACE,
	string(rune(0)): EOF,
}

var KeywordLiteral2TokenType = map[string]TokenType {
	"let": LET, "fn": FUNCTION,
}

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

func LookupIdent(ident string) TokenType {
	if keywordType, exist := KeywordLiteral2TokenType[ident]; exist {
		return keywordType
	}
	return IDENT
}