package data

import (
	"github.com/shopspring/decimal"
	"time"
)

func (SuggestionData) TableName() string {
	return "suggestions"
}

func (LoanData) TableName() string {
	return "loans"
}

type LoanData struct {
	ID          uint             `gorm:"primaryKey"`
	Code        string           `gorm:"type:varchar(50);unique;not null"`
	Name        string           `gorm:"type:varchar(100);not null;column:name"`
	Document    string           `gorm:"type:varchar(20);not null;column:document"`
	BirthDate   time.Time        `gorm:"type:date;not null;column:birth_date"`
	Value       decimal.Decimal  `gorm:"type:decimal(10,2);not null;column:value"`
	RequestDate time.Time        `gorm:"type:datetime;not null;column:request_date"`
	Status      int              `gorm:"type:tinyint;not null"`
	Suggestions []SuggestionData `gorm:"foreignKey:LoanID"`
	CreatedAt   time.Time        `gorm:"type:datetime"`
	UpdatedAt   time.Time        `gorm:"type:datetime"`
}

type SuggestionData struct {
	ID           uint            `gorm:"primaryKey"`
	LoanID       uint            `gorm:"not null"`
	Fees         decimal.Decimal `gorm:"type:decimal(10,2);not null;column:fees"`
	Installments int             `gorm:"type:int;not null;column:installments"`
	EndValue     decimal.Decimal `gorm:"type:decimal(10,2);not null;column:end_value"`
	Status       int             `gorm:"type:integer;not null;column:status"`
	CreatedAt    time.Time       `gorm:"type:datetime"`
	UpdatedAt    time.Time       `gorm:"type:datetime"`
}
