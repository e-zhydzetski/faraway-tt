package domain

import "context"

type Quoter interface {
	Quote(ctx context.Context) (string, error)
}

type POWCheck interface {
	Input() []byte
	Check(answer uint64) bool
}
