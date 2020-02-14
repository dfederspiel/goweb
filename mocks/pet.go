package mocks

import (
	"rsi.com/go-training/interfaces"
	"rsi.com/go-training/models"
)

type testablePetRepository struct{}

func (t testablePetRepository) GetAll() (pets []*models.Pet, err error) {
	return []*models.Pet{{
		ID:    "1",
		Name:  "Buddy",
		Age:   4,
		Legs:  true,
		Color: "Green",
	}}, nil
}

func (t testablePetRepository) GetById(id string) (p *models.Pet, err error) {
	return &models.Pet{
		ID:    "1",
		Name:  "Buddy",
		Age:   4,
		Legs:  true,
		Color: "Green",
	}, nil
}

func (t testablePetRepository) Create(pet *models.Pet) (err error) {
	return nil
}

func (t testablePetRepository) Update(pet *models.Pet) (err error) {
	return nil
}

func (t testablePetRepository) DeleteById(id string) (err error) {
	return nil
}

func NewTestablePetRepository() interfaces.PetRepository {
	return &testablePetRepository{}
}
