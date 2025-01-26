package entities

import (
	"github.com/shopspring/decimal"
	"time"
)

type DetailsLoan struct {
	Value       decimal.Decimal
	requestDate time.Time
}
