package domain

import "context"

type Quoter interface {
	Quote(ctx context.Context) (string, error)
}

type POWCheckFactoryFunc func(difficulty uint64) (POWCheck, error)

type POWCheck interface {
	Challenge() []byte
	Verify(proof uint64) bool
}

type Client interface {
	WriteBytes(data []byte) error
	ReadUint64() (uint64, error)
	WriteString(data string) error
}
