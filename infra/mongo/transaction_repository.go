package mongo

import "go-stock-exchange-shares-control/internal/domain"

type transactionRepository struct {
}

func NewTransactionRepositoryMongo(connectionString string) domain.TransactionRepository {

	return &transactionRepository{}
}

func (tr *transactionRepository) New(transaction *domain.Transaction) error {

}

func (tr *transactionRepository) Get(code string) ([]domain.Transaction, error) {

}
