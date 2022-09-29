package pow

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestBCryptPOWCheck(t *testing.T) {
	check, err := NewBCryptCheck(10)
	require.NoError(t, err)
	_, err = bcrypt.Cost(check.Input())
	require.NoError(t, err)
	t.Log(string(check.Input()))

	x := BCryptSolve(check.Input())
	require.True(t, check.Check(x))
}
