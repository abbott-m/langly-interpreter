package lexer

import (
	"langly/internal/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	type Expected struct {
		t       token.TokenType
		literal string
	}

	type TestToken struct {
		input    string
		expected []Expected
	}

	tests := []TestToken{
		{
			input: `=+(){},;`,
			expected: []Expected{
				{token.ASSIGN, "="},
				{token.PLUS, "+"},
				{token.LPAREN, "("},
				{token.RPAREN, ")"},
				{token.LBRACE, "{"},
				{token.RBRACE, "}"},
				{token.COMMA, ","},
				{token.SEMICOLON, ";"},
				{token.EOF, ""},
			},
		},
		{
			input: `
				let five = 5;
				let ten = 10;

				let add = fn(x, y) {
				  x + y;
				};

				let result = add(five, ten);
			`,
			expected: []Expected{
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
				{token.EOF, ""},
			},
		},
	}

	for i, tt := range tests {

		l := New(tt.input)
		for _, expected := range tt.expected {
			tok := l.NextToken()
			if tok.Type != expected.t {
				t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
					i, expected.t, tok.Type)
			}

			if tok.Literal != expected.literal {
				t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
					i, expected.literal, tok.Literal)
			}
		}
	}
}
