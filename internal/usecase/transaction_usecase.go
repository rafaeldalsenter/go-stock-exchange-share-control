package usecase

import (
	"go-stock-exchange-shares-control/internal/domain"
	"time"
)

type transactionUseCase struct {
	transactionRepository domain.TransactionRepository
}

func NewTransactionUseCase(transactionRepository domain.TransactionRepository) domain.NewTransactionUseCase {
	return &transactionUseCase{
		transactionRepository: transactionRepository,
	}
}

func (tr *transactionUseCase) NewSale(code string, date time.Time, quantity float64, value float64, tax float64) error {

	transaction, err := domain.NewTransaction(domain.Sale, code, date, quantity, value, tax)
	if err != nil {
		return err
	}

	return tr.transactionRepository.New(transaction)
}

func (tr *transactionUseCase) NewPurchase(code string, date time.Time, quantity float64, value float64, tax float64) error {

	transaction, err := domain.NewTransaction(domain.Purchase, code, date, quantity, value, tax)
	if err != nil {
		return err
	}

	return tr.transactionRepository.New(transaction)
}
