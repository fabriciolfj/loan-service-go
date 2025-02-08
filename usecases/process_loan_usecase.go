package usecases

import (
	"github.com/fabriciolfj/loan-service-go/entities"
	"log"
)

type SaveLoan interface {
	SaveLoan(loan *entities.Loan) error
}

type NotifyLoanPending interface {
	NotifyLoanPending(loan *entities.Loan) error
}

type ProcessLoanUseCase struct {
	save   SaveLoan
	notify NotifyLoanPending
}

func ProviderProcessLoanUseCase(save SaveLoan, notify NotifyLoanPending) *ProcessLoanUseCase {
	return &ProcessLoanUseCase{
		save:   save,
		notify: notify,
	}
}

func (u ProcessLoanUseCase) Execute(loan entities.Loan) error {
	if err := u.save.SaveLoan(&loan); err != nil {
		log.Fatal("failed to save loan: %w", err)
		return err
	}

	log.Printf("loan saved %s", loan.Code)

	if err := u.notify.NotifyLoanPending(&loan); err != nil {
		log.Fatal("failed to notify loan: %w", err)
	}

	log.Printf("loan notified %s", loan.Code)

	return nil
}
