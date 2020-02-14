package interfaces

import "rsi.com/go-training/models"

type PetRepository interface {
	GetAll() (pets []*models.Pet, err error)
	GetById(id string) (pet *models.Pet, err error)
	Create(pet *models.Pet) (err error)
	Update(pet *models.Pet) (err error)
	DeleteById(id string) (err error)
}
