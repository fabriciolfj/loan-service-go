package entities

import "time"

type Customer struct {
	Name      string    `json:"name"`
	Document  string    `json:"document"`
	BirthDate time.Time `json:"birth_date"`
}
