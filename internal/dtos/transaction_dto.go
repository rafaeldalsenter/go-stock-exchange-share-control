package dtos

import (
	"time"
)

type TransactionDto struct {
	Date     time.Time
	Quantity float64
	Value    float64
	Tax      float64
	Type     string
}
