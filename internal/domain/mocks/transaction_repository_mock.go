package mocks

import (
	"context"
	"go-stock-exchange-shares-control/internal/domain"

	"github.com/stretchr/testify/mock"
)

type TransactionRepositoryMock struct {
	mock.Mock
}

func (r *TransactionRepositoryMock) New(c context.Context, transaction *domain.Transaction) (string, error) {
	args := r.Called(c, transaction)

	var r0 string
	if rf, ok := args.Get(0).(func(c context.Context, transaction *domain.Transaction) string); ok {
		r0 = rf(c, transaction)
	} else {
		r0 = args.Get(0).(string)
	}

	var r1 error
	if rf, ok := args.Get(1).(func(c context.Context, transaction *domain.Transaction) error); ok {
		r1 = rf(c, transaction)
	} else {
		r1 = args.Error(1)
	}

	return r0, r1
}

func (r *TransactionRepositoryMock) Get(c context.Context, code string) ([]domain.Transaction, error) {
	args := r.Called(c, code)

	var r0 []domain.Transaction
	if rf, ok := args.Get(0).(func(c context.Context, code string) []domain.Transaction); ok {
		r0 = rf(c, code)
	} else {
		r0 = args.Get(0).([]domain.Transaction)
	}

	var r1 error
	if rf, ok := args.Get(1).(func(c context.Context, code string) error); ok {
		r1 = rf(c, code)
	} else {
		r1 = args.Error(1)
	}

	return r0, r1
}
