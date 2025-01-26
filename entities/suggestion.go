package entities

import "github.com/shopspring/decimal"

type Suggestion struct {
	Fees         decimal.Decimal
	Installments int
	endValue     decimal.Decimal
	Status       StatusSuggestion
}
