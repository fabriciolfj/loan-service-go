package repositories

import (
	"github.com/fabriciolfj/loan-service-go/data"
	"gorm.io/gorm"
)

type LoanRepository struct {
	db *gorm.DB
}

func ProviderLoanRepository(db *gorm.DB) *LoanRepository {
	return &LoanRepository{
		db: db,
	}
}

func (r *LoanRepository) SaveLoan(data *data.LoanData) error {
	return r.db.Save(&data).Error
}

func (r *LoanRepository) FindLoanByCode(code string) (*data.LoanData, error) {
	var loan data.LoanData
	result := r.db.Preload("Suggestions").Where(&data.LoanData{Code: code}).First(&loan)

	return &loan, result.Error
}
