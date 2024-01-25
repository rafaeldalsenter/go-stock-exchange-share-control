package utils

import (
	"errors"
	"go-stock-exchange-shares-control/internal/dtos"
	"strconv"
	"time"
)

func ConverterTransactionValues(typeTr string, date string, quantity string, value string, tax string) (*dtos.TransactionDto, error) {

	qtd, err := strconv.ParseFloat(quantity, 64)
	if err != nil {
		return nil, errors.New("Quantity invalid format")
	}

	val, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return nil, errors.New("Value invalid format")
	}

	tx, err := strconv.ParseFloat(tax, 64)
	if err != nil {
		return nil, errors.New("Tax invalid format")
	}

	tim, err := time.Parse(time.RFC3339, date)
	if err != nil {
		return nil, errors.New("Date invalid format")
	}

	return &dtos.TransactionDto{
		Type:     typeTr,
		Quantity: qtd,
		Value:    val,
		Tax:      tx,
		Date:     tim,
	}, nil
}
