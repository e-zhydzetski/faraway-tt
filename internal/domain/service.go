package domain

import (
	"context"
	"errors"
)

func NewService(powCheckFactory POWCheckFactoryFunc, quoter Quoter, difficulty uint64) Service {
	return Service{
		powCheckFactory: powCheckFactory,
		quoter:          quoter,
		difficulty:      difficulty,
	}
}

type Service struct {
	powCheckFactory POWCheckFactoryFunc
	quoter          Quoter
	difficulty      uint64
}

func (h Service) ServeClient(ctx context.Context, client Client) error {
	powCheck, err := h.powCheckFactory(h.difficulty) // TODO maybe use dynamic difficulty based on current load
	if err != nil {
		return err
	}

	err = client.WriteBytes(powCheck.Challenge())
	if err != nil {
		return err
	}
	checkCtx, cancel := context.WithTimeout(ctx, powCheck.ReasonableTimeout())
	defer cancel()
	proof, err := client.ReadUint64(checkCtx)
	if err != nil {
		return err
	}
	if !powCheck.Verify(proof) {
		return errors.New("verification error")
	}

	quote, err := h.quoter.Quote(ctx)
	if err != nil {
		return err
	}
	return client.WriteString(quote)
}
