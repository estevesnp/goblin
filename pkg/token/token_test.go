package token

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLookupIdent(t *testing.T) {
	assert.Equal(t, FUNCTION, LookupIdent("fn"))
	assert.Equal(t, LET, LookupIdent("let"))
	assert.Equal(t, IF, LookupIdent("if"))
	assert.Equal(t, ELSE, LookupIdent("else"))
	assert.Equal(t, RETURN, LookupIdent("return"))
	assert.Equal(t, TRUE, LookupIdent("true"))
	assert.Equal(t, FALSE, LookupIdent("false"))
	assert.Equal(t, IDENT, LookupIdent("x"))
	assert.Equal(t, IDENT, LookupIdent("var"))
}
