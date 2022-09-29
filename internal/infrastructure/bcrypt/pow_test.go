package bcrypt

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPOWCheck(t *testing.T) {
	check, err := NewPOWCheck(10)
	require.NoError(t, err)
	_, err = bcrypt.Cost(check.Input())
	require.NoError(t, err)
	t.Log(string(check.Input()))

	x := Solve(check.Input())
	require.True(t, check.Check(x))
}
