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
}

var KeywordLiteral2TokenType = map[string]TokenType {
	"let": LET, "fn": FUNCTION,
}

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}
