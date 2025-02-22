package usecases

import (
	"github.com/fabriciolfj/loan-service-go/entities"
	"log"
)

type FindLoan interface {
	FindLoan(code string) (*entities.Loan, error)
}

type GetLoanUseCase struct {
	findLoan FindLoan
}

func ProviderGetLoanUseCase(findLoan FindLoan) *GetLoanUseCase {
	return &GetLoanUseCase{
		findLoan: findLoan,
	}
}

func (c *GetLoanUseCase) Execute(code string) (*entities.Loan, error) {
	loan, err := c.findLoan.FindLoan(code)
	if err != nil {
		log.Fatalf("Loan %s not found or problema get loan, details %s", code, err.Error())
		return nil, err
	}

	return loan, nil
}
