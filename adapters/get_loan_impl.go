package adapters

import (
	"github.com/fabriciolfj/loan-service-go/entities"
	"github.com/fabriciolfj/loan-service-go/repositories"
	"log"
)

type FindLoanAdapter struct {
	repository repositories.LoanRepository
}

func ProvideFindLoanAdapter(repository *repositories.LoanRepository) *FindLoanAdapter {
	return &FindLoanAdapter{
		repository: *repository,
	}
}

func (find *FindLoanAdapter) FindLoan(code string) (*entities.Loan, error) {
	log.Printf("find loan %s", code)

	data, err := find.repository.FindLoanByCode(code)
	if err != nil {
		return nil, err
	}

	return LoanEntityMapper(data), nil
}
