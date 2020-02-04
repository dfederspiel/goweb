package v2

import "rsi.com/go-training/api/v1/pet"

type Owner struct {
	ID        int64     `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Pets      []pet.Pet `json:"pets"`
}
