package adapters

import (
	"github.com/fabriciolfj/loan-service-go/data"
	"github.com/fabriciolfj/loan-service-go/entities"
)

func LoanDataMapper(entity *entities.Loan) *data.LoanData {
	var suggestions []data.SuggestionData
	for _, sugestion := range entity.Suggestions {
		suggestions = append(suggestions, data.SuggestionData{
			Fees:         sugestion.Fees,
			Installments: sugestion.Installments,
			EndValue:     sugestion.EndValue,
			Status:       int(sugestion.Status),
		})
	}

	return &data.LoanData{
		Code:        entity.Code,
		Name:        entity.Customer.Name,
		Document:    entity.Customer.Document,
		BirthDate:   entity.Customer.BirthDate,
		Value:       entity.Details.Value,
		RequestDate: entity.Details.RequestDate,
		Status:      int(entity.Status),
		Suggestions: suggestions,
	}
}

func LoanEntityMapper(data *data.LoanData) *entities.Loan {
	var suggestions []entities.Suggestion
	for _, suggestion := range data.Suggestions {
		suggestions = append(suggestions, entities.Suggestion{
			Fees:         suggestion.Fees,
			Installments: suggestion.Installments,
			EndValue:     suggestion.EndValue,
			Status:       entities.StatusSuggestion(suggestion.Status),
		})
	}
	return &entities.Loan{
		Code: data.Code,
		Customer: entities.Customer{
			Name:     data.Name,
			Document: data.Document,
		},
		Details: entities.DetailsLoan{
			Value:       data.Value,
			RequestDate: data.RequestDate,
		},
		Status: entities.LoanStatus(data.Status),
	}
}
