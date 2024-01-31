package usecase

import (
	"context"
	"errors"
	"go-stock-exchange-shares-control/internal/domain"
)

type stockPriceUseCase struct {
	transactionRepository domain.TransactionRepository
}

func NewStockPriceUseCase(transactionRepository domain.TransactionRepository) domain.StockPriceUseCase {
	return &stockPriceUseCase{
		transactionRepository: transactionRepository,
	}
}

func (st *stockPriceUseCase) AveragePurchasePrice(c context.Context, code string) (float64, error) {
	transactions, err := st.transactionRepository.Get(c, code)
	if err != nil {
		return 0.0, errors.New("Repository error")
	}

	stock, err := domain.NewStock(code, transactions)
	if err != nil {
		return 0.0, err
	}

	return stock.AveragePurchasePrice()
}
