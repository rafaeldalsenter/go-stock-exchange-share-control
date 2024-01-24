package controllers

import (
	"errors"
	"go-stock-exchange-shares-control/internal/dtos"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func (c *controller) TransactionPost(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {

	code := chi.URLParam(r, "code")

	var transactionDto dtos.TransactionDto
	render.DecodeJSON(r.Body, &transactionDto)

	var err error

	if transactionDto.Type == "purchase" {
		err = c.NewTransactionUseCase.NewPurchase(code, transactionDto.Date, transactionDto.Quantity, transactionDto.Value, transactionDto.Tax)
	} else if transactionDto.Type == "sale" {
		err = c.NewTransactionUseCase.NewSale(code, transactionDto.Date, transactionDto.Quantity, transactionDto.Value, transactionDto.Tax)
	} else {
		err = errors.New("Transaction type not identified")
	}

	if err != nil {
		return nil, 500, err
	}

	return nil, 204, nil
}
