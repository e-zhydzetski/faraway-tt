package pow

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestBCryptPOWCheck(t *testing.T) {
	check, err := NewBCryptCheck(10)
	require.NoError(t, err)
	_, err = bcrypt.Cost(check.Challenge())
	require.NoError(t, err)
	t.Log(string(check.Challenge()))

	x, err := BCryptProve(context.Background(), check.Challenge())
	require.NoError(t, err)
	require.True(t, check.Verify(x))
}

func TestBCryptPOWCheckContextCancel(t *testing.T) {
	check, err := NewBCryptCheck(1000)
	require.NoError(t, err)
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	_, err = BCryptProve(ctx, check.Challenge())
	require.ErrorIs(t, err, context.DeadlineExceeded)
}
