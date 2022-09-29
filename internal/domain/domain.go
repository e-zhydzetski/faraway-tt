package domain

import "context"

type Quoter interface {
	Quote(ctx context.Context) (string, error)
}
