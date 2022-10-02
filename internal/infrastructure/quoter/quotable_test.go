package quoter

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQuotable(t *testing.T) {
	c := NewQuotableClient()
	q, err := c.Quote(context.Background())
	require.NoError(t, err)
	require.True(t, len(q) > 0)
	t.Log(q)
}
