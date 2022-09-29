package bcrypt

import (
	"math/big"
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

	var x uint64
	for {
		err := bcrypt.CompareHashAndPassword(check.Input(), big.NewInt(int64(x)).Bytes())
		if err == nil {
			break
		}
		x++
	}
	require.True(t, check.Check(x))
}
