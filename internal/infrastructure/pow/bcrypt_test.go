package pow

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestBCryptPOWCheck(t *testing.T) {
	check, err := NewBCryptCheck(10)
	require.NoError(t, err)
	_, err = bcrypt.Cost(check.Challenge())
	require.NoError(t, err)
	t.Log(string(check.Challenge()))

	x := BCryptProve(check.Challenge())
	require.True(t, check.Verify(x))
}
