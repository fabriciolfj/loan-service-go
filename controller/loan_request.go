package controller

import (
	"github.com/fabriciolfj/loan-service-go/entities"
	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
	"time"
)

type LoanRequest struct {
	Code      string             `json:"code"`
	Name      string             `json:"name"`
	Document  string             `json:"document"`
	BirthDate time.Time          `json:"birth_date"`
	Details   DetailsLoanRequest `json:"details"`
}

type DetailsLoanRequest struct {
	Value       decimal.Decimal `json:"value"`
	RequestDate time.Time       `json:"request_date"`
}

func (r LoanRequest) ToEntity() entities.Loan {
	return entities.Loan{
		Code: uuid.NewV4().String(),
		Customer: entities.Customer{
			Name:      r.Name,
			BirthDate: r.BirthDate,
			Document:  r.Document,
		},
		Details: entities.DetailsLoan{
			Value:       r.Details.Value,
			RequestDate: r.Details.RequestDate,
		},
		Status: entities.PENDING,
	}
}
