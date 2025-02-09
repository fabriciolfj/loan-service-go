package adapters

import "github.com/fabriciolfj/loan-service-go/entities"

type SaveLoanAdapter struct {
}

func ProvideSaveLoanAdapter() *SaveLoanAdapter {
	return &SaveLoanAdapter{}
}

func (adapter *SaveLoanAdapter) SaveLoan(loan *entities.Loan) error {
	println("Saving Loan", loan)
	return nil
}
