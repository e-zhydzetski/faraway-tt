package domain

import "context"

type Quoter interface {
	Quote(ctx context.Context) (string, error)
}

type POWCheckFactoryFunc func(complexity uint64) (POWCheck, error)

type POWCheck interface {
	Input() []byte
	Check(answer uint64) bool
}

type Client interface {
	POWVerification(ctx context.Context, input []byte) (uint64, error)
	SendQuote(ctx context.Context, quote string) error
}
