package entities

import "time"

type Customer struct {
	Name      string
	Document  string
	birthDate time.Time
}
