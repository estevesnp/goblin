package lexer

import (
	"testing"

	"github.com/estevesnp/dsb/pkg/token"
	"github.com/stretchr/testify/assert"
)

func TestIsLetter(t *testing.T) {
	assert.True(t, isLetter('a'), "'a' should be true")
	assert.True(t, isLetter('h'), "'h' should be true")
	assert.True(t, isLetter('z'), "'z' should be true")
	assert.True(t, isLetter('A'), "'A' should be true")
	assert.True(t, isLetter('H'), "'H' should be true")
	assert.True(t, isLetter('Z'), "'Z' should be true")

	assert.False(t, isLetter('1'), "'1' should be false")
	assert.False(t, isLetter('.'), "'.' should be false")
	assert.False(t, isLetter('/'), "'/' should be false")
	assert.False(t, isLetter('&'), "'&' should be false")
	assert.False(t, isLetter('~'), "'1' should be false")
	assert.False(t, isLetter(0), "null should be false")
}

func TestIsDigit(t *testing.T) {
	assert.True(t, isDigit('0'), "'0' should be true")
	assert.True(t, isDigit('5'), "'5' should be true")
	assert.True(t, isDigit('9'), "'9' should be true")

	assert.False(t, isDigit('a'), "'a' should be false")
	assert.False(t, isDigit('/'), "'/' should be false")
	assert.False(t, isDigit('-'), "'-' should be false")
}

func TestNextToken(t *testing.T) {
	input := `let five = 5;
let ten = 10;

let add = fn(x, y) {
    x + y;
};

let result = add(five, ten);
!-/*5;
5 < 10 > 5;

if (5 < 10) {
    return true;
} else {
    return false;
}
    `

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.EOF, ""},
	}

	l := New(input)

	for _, tt := range tests {
		tok := l.NextToken()

		assert.Equal(t, tt.expectedType, tok.Type)
		assert.Equal(t, tt.expectedLiteral, tok.Literal)
	}
}
