package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	typeTr   = "purchase"
	date     = "2019-10-12T07:20:50.52Z"
	quantity = "12.02"
	value    = "0.22"
	tax      = "9.90"
)

func TestConverterTransactionValuesWithInvalidDate(t *testing.T) {
	assert := assert.New(t)
	invalidDate := "121/asasa/222"

	_, err := ConverterTransactionValues(typeTr, invalidDate, quantity, value, tax)

	assert.Error(err)
}

func TestConverterTransactionValuesWithInvalidQuantity(t *testing.T) {
	assert := assert.New(t)
	invalidQuantity := "sadssda"

	_, err := ConverterTransactionValues(typeTr, date, invalidQuantity, value, tax)

	assert.Error(err)
}

func TestConverterTransactionValuesWithInvalidValue(t *testing.T) {
	assert := assert.New(t)
	invalidValue := "asdasdds"

	_, err := ConverterTransactionValues(typeTr, date, quantity, invalidValue, tax)

	assert.Error(err)
}

func TestConverterTransactionValuesWithInvalidTax(t *testing.T) {
	assert := assert.New(t)
	invalidTax := "asdsad"

	_, err := ConverterTransactionValues(typeTr, date, quantity, value, invalidTax)

	assert.Error(err)
}

func TestConverterTransactionValuesWithValidValues(t *testing.T) {
	assert := assert.New(t)

	tr, err := ConverterTransactionValues(typeTr, date, quantity, value, tax)

	assert.NoError(err)
	assert.Equal(typeTr, tr.Type)
	assert.Equal(time.Time(time.Date(2019, time.October, 12, 7, 20, 50, 520000000, time.UTC)), tr.Date)
	assert.Equal(12.02, tr.Quantity)
	assert.Equal(0.22, tr.Value)
	assert.Equal(9.90, tr.Tax)
}
