package entities

import "github.com/shopspring/decimal"

type Suggestion struct {
	Fees         decimal.Decimal
	Installments int
	EndValue     decimal.Decimal
	Status       StatusSuggestion
}
