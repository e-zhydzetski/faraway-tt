package quotable

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQuote(t *testing.T) {
	c := NewClient()
	q, err := c.Quote(context.Background())
	require.NoError(t, err)
	require.True(t, len(q) > 0)
	t.Log(q)
}
