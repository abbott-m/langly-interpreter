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
