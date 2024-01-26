package controllers

import (
	"errors"
	"go-stock-exchange-shares-control/internal/dtos"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func (c *controller) TransactionPost(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	ctx := r.Context()
	code := chi.URLParam(r, "code")

	var transactionDto dtos.TransactionDto
	err := render.DecodeJSON(r.Body, &transactionDto)

	if err != nil {
		return nil, 400, err
	}

	if transactionDto.Type == "purchase" {
		_, err = c.NewTransactionUseCase.NewPurchase(ctx, code, transactionDto.Date, transactionDto.Quantity, transactionDto.Value, transactionDto.Tax)
	} else if transactionDto.Type == "sale" {
		_, err = c.NewTransactionUseCase.NewSale(ctx, code, transactionDto.Date, transactionDto.Quantity, transactionDto.Value, transactionDto.Tax)
	} else {
		err = errors.New("Transaction type not identified")
	}

	if err != nil {
		return nil, 500, err
	}

	return nil, 204, nil
}
