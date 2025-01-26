package entities

type Loan struct {
	Code        string
	Customer    Customer
	Details     DetailsLoan
	Status      LoanStatus
	Suggestions []Suggestion
}
