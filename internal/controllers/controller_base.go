package controllers

import (
	"go-stock-exchange-shares-control/internal/domain"
	"net/http"

	"github.com/go-chi/render"
)

type controller struct {
	StockPriceUseCase     domain.StockPriceUseCase
	NewTransactionUseCase domain.NewTransactionUseCase
}

func NewController(stockPriceUseCase domain.StockPriceUseCase, newTransactionUseCase domain.NewTransactionUseCase) *controller {
	return &controller{
		StockPriceUseCase:     stockPriceUseCase,
		NewTransactionUseCase: newTransactionUseCase,
	}
}

type ControllerFunc func(w http.ResponseWriter, r *http.Request) (interface{}, int, error)

func ControllerBase(controllerFunc ControllerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		result, statusCode, err := controllerFunc(w, r)
		render.Status(r, statusCode)

		if err != nil {
			render.JSON(w, r, map[string]string{"error": err.Error()})
			return
		}

		render.JSON(w, r, result)
	})
}
