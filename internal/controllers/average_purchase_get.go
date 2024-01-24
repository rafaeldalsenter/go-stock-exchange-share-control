package controllers

import (
	"net/http"

	"github.com/go-chi/chi"
)

func (c *controller) AveragePurchaseGet(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {

	code := chi.URLParam(r, "code")

	result, err := c.StockPriceUseCase.AveragePurchasePrice(code)

	if err != nil {
		return nil, 500, err
	}

	return result, 200, nil
}
