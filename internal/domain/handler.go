package domain

import (
	"context"
	"errors"
)

func NewHandler(powCheckFactory POWCheckFactoryFunc, quoter Quoter) Handler {
	return Handler{
		powCheckFactory: powCheckFactory,
		quoter:          quoter,
	}
}

type Handler struct {
	powCheckFactory POWCheckFactoryFunc
	quoter          Quoter
}

func (h Handler) Handle(ctx context.Context, client Client) error {
	powCheck, err := h.powCheckFactory(100)
	if err != nil {
		return err
	}
	answer, err := client.POWVerification(ctx, powCheck.Input())
	if err != nil {
		return err
	}
	if !powCheck.Check(answer) {
		return errors.New("verification error")
	}
	quote, err := h.quoter.Quote(ctx)
	if err != nil {
		return err
	}
	return client.SendQuote(ctx, quote)
}
