package mocks

import (
	"go-stock-exchange-shares-control/internal/domain"

	"github.com/stretchr/testify/mock"
)

type TransactionRepositoryMock struct {
	mock.Mock
}

func (r *TransactionRepositoryMock) New(transaction *domain.Transaction) error {
	args := r.Called(transaction)
	return args.Error(0)
}
