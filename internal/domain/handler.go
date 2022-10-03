package domain

import (
	"context"
	"errors"
)

func NewHandler(powCheckFactory POWCheckFactoryFunc, quoter Quoter, difficulty uint64) Handler {
	return Handler{
		powCheckFactory: powCheckFactory,
		quoter:          quoter,
		difficulty:      difficulty,
	}
}

type Handler struct {
	powCheckFactory POWCheckFactoryFunc
	quoter          Quoter
	difficulty      uint64
}

func (h Handler) Handle(ctx context.Context, client Client) error {
	powCheck, err := h.powCheckFactory(h.difficulty) // TODO maybe use dynamic difficulty based on current load
	if err != nil {
		return err
	}

	err = client.WriteBytes(powCheck.Challenge())
	if err != nil {
		return err
	}
	proof, err := client.ReadUint64()
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
