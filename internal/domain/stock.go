package domain

import (
	"context"
	"errors"
	"go-stock-exchange-shares-control/internal/dtos"
	"math"
	"slices"
)

type Stock struct {
	Code         string
	Transactions []Transaction
}

func NewStock(code string, transactions []Transaction) (*Stock, error) {

	if code == "" {
		return nil, errors.New("code is required")
	}

	if len(transactions) == 0 {
		return nil, errors.New("stock not have transactions")
	}

	return &Stock{
		Code:         code,
		Transactions: transactions,
	}, nil
}

func (s *Stock) AveragePurchasePrice() (float64, error) {
	sumQuantity := 0.0
	sumValue := 0.0

	for _, transaction := range s.Transactions {

		if transaction.Type != Purchase {
			continue
		}

		sumValue += (transaction.Quantity * transaction.Value) + transaction.Tax
		sumQuantity += transaction.Quantity
	}

	if sumValue == 0.0 {
		return 0.0, errors.New("There are no valid purchase records")
	}

	ratio := math.Pow(10, float64(2))
	return math.Round((sumValue/sumQuantity)*ratio) / ratio, nil
}

func (s *Stock) SalesResult() ([]dtos.SalesDto, error) {

	sumPurchaseQuantity := 0.0
	sumPurchaseValue := 0.0
	ratio := math.Pow(10, float64(2))
	salesResult := make([]dtos.SalesDto, 0)

	transactionsOrdered := s.Transactions
	slices.SortFunc(transactionsOrdered, func(tr1, tr2 Transaction) int {
		return tr1.Date.Compare(tr2.Date)
	})

	for _, transaction := range transactionsOrdered {

		if transaction.Type == Purchase {
			sumPurchaseValue += (transaction.Quantity * transaction.Value) + transaction.Tax
			sumPurchaseQuantity += transaction.Quantity
			continue
		}

		if transaction.Type == Sale {
			if sumPurchaseValue == 0.0 {
				return []dtos.SalesDto{}, errors.New("Inconsistent records")
			}

			averagePurchase := math.Round((sumPurchaseValue/sumPurchaseQuantity)*ratio) / ratio
			result := math.Round((((transaction.Quantity*transaction.Value)-transaction.Tax)-(transaction.Quantity*averagePurchase))*ratio) / ratio
			salesResult = append(salesResult, dtos.SalesDto{Date: transaction.Date, Result: result})
			continue
		}

	}

	return salesResult, nil
}

type StockPriceUseCase interface {
	AveragePurchasePrice(c context.Context, code string) (float64, error)
}
