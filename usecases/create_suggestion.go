package usecases

import (
	"github.com/fabriciolfj/loan-service-go/entities"
	"github.com/shopspring/decimal"
	"log"
)

type CreateSuggestion struct {
}

func ProviderSuggestionUseCase() *CreateSuggestion {
	return &CreateSuggestion{}
}

func (c CreateSuggestion) Execute(loan *entities.Loan) []entities.Suggestion {
	value := loan.Details.Value
	return c.processCreation(&value)
}

func (c CreateSuggestion) processCreation(value *decimal.Decimal) []entities.Suggestion {
	fee := 0.015
	installments := 12
	percent := value.Mul(decimal.NewFromFloat(0.3))
	parc := percent.Mul(decimal.NewFromFloat(fee)).Add(percent)
	total := parc.Mul(decimal.NewFromInt(int64(installments)))

	var suggestions []entities.Suggestion
	suggestion := entities.Suggestion{
		Fees:         decimal.NewFromFloat(fee),
		EndValue:     total,
		Installments: installments,
		Status:       entities.StatusSuggestion(0),
	}

	log.Println("suggestion created: ", suggestion)
	suggestions = append(suggestions, suggestion)

	return suggestions
}
