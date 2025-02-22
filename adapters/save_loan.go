package adapters

import (
	"github.com/fabriciolfj/loan-service-go/entities"
	"github.com/fabriciolfj/loan-service-go/repositories"
	"log"
)

type SaveLoanAdapter struct {
	repository repositories.LoanRepository
}

func ProvideSaveLoanAdapter(repository *repositories.LoanRepository) *SaveLoanAdapter {
	return &SaveLoanAdapter{
		repository: *repository,
	}
}

func (adapter *SaveLoanAdapter) SaveLoan(loan *entities.Loan) error {
	data := LoanDataMapper(loan)
	log.Printf("Saving Loan %v", data)

	return adapter.repository.SaveLoan(data)
}
