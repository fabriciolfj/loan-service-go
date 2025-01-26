package entities

type LoanStatus int

const (
	PENDING LoanStatus = iota
	APPROVED
	DISAPPROVED
	CANCELED
	ACCEPTED
)
