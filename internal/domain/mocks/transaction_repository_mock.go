package mocks

import (
	"go-stock-exchange-shares-control/internal/domain"

	"github.com/stretchr/testify/mock"
)

type TransactionRepositoryMock struct {
	mock.Mock
}

func (r *TransactionRepositoryMock) New(transaction *domain.Transaction) (string, error) {
	args := r.Called(transaction)

	var r0 string
	if rf, ok := args.Get(0).(func(transaction *domain.Transaction) string); ok {
		r0 = rf(transaction)
	} else {
		r0 = args.Get(0).(string)
	}

	var r1 error
	if rf, ok := args.Get(1).(func(transaction *domain.Transaction) error); ok {
		r1 = rf(transaction)
	} else {
		r1 = args.Error(1)
	}

	return r0, r1
}

func (r *TransactionRepositoryMock) Get(code string) ([]domain.Transaction, error) {
	args := r.Called(code)

	var r0 []domain.Transaction
	if rf, ok := args.Get(0).(func(code string) []domain.Transaction); ok {
		r0 = rf(code)
	} else {
		r0 = args.Get(0).([]domain.Transaction)
	}

	var r1 error
	if rf, ok := args.Get(1).(func(code string) error); ok {
		r1 = rf(code)
	} else {
		r1 = args.Error(1)
	}

	return r0, r1
}
