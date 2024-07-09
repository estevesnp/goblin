package parser

import (
	"testing"

	"github.com/estevesnp/dsb/pkg/ast"
	"github.com/estevesnp/dsb/pkg/lexer"
	"github.com/stretchr/testify/require"
)

func TestLetStatements(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foo = 42;
`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()

	require.NotNil(t, program)
	require.Equal(t, 3, len(program.Statements))

	tests := []struct {
		expectedIdent string
	}{
		{"x"},
		{"y"},
		{"foo"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		require.Equal(t, "let", stmt.TokenLiteral())

		letStmt, ok := stmt.(*ast.LetStatement)
		require.True(t, ok)

		require.Equal(t, tt.expectedIdent, letStmt.Name.Value)
		require.Equal(t, tt.expectedIdent, letStmt.Name.TokenLiteral())
	}
}
