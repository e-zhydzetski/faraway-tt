package domain

import (
	"context"
	"time"
)

type Quoter interface {
	Quote(ctx context.Context) (string, error)
}

type POWCheckFactoryFunc func(difficulty uint64) (POWCheck, error)

type POWCheck interface {
	Challenge() []byte
	ReasonableTimeout() time.Duration
	Verify(proof uint64) bool
}

type Client interface {
	WriteBytes(data []byte) error
	ReadUint64(ctx context.Context) (uint64, error)
	WriteString(data string) error
}
