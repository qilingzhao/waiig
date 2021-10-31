package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 1343456

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

	EQ     = "=="
	NOT_EQ = "!="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

var symbolLiteral2TokenType = map[string]TokenType{
	"=": ASSIGN, "+": PLUS, "-": MINUS, "!": BANG, "*": ASTERISK, "/": SLASH,
	"<": LT, ">": GT, "==": EQ, "!=": NOT_EQ,
	",": COMMA, ";": SEMICOLON,
	"(": LPAREN, ")": RPAREN, "{": LBRACE, "}": RBRACE,
	// string(rune(0)): EOF,
}

var keywordLiteral2TokenType = map[string]TokenType{
	"let": LET, "fn": FUNCTION, "true": TRUE, "false": FALSE,
	"if": IF, "else": ELSE, "return": RETURN,
}

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

func LookupIdent(ident string) TokenType {
	if keywordType, exist := keywordLiteral2TokenType[ident]; exist {
		return keywordType
	}
	return IDENT
}

func LookupSymbol(ident string) TokenType {
	if symbolType, exist := symbolLiteral2TokenType[ident]; exist {
		return symbolType
	}
	return ILLEGAL
}