package token

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLookupIdent(t *testing.T) {
	assert.Equal(t, FUNCTION, LookupIdent("fn"))
	assert.Equal(t, LET, LookupIdent("let"))
	assert.Equal(t, IDENT, LookupIdent("x"))
	assert.Equal(t, IDENT, LookupIdent("var"))
}
