package pet

import (
	"rsi.com/go-training/interfaces"
	"rsi.com/go-training/models"
)

type Service interface {
	GetAll() ([]*models.Pet, error)
	GetById(id string) (*models.Pet, error)
	Create(pet *models.Pet) (err error)
	Update(pet *models.Pet) (err error)
	DeleteById(id string) (err error)
}

type service struct {
	repo interfaces.PetRepository
}

func (s service) GetAll() ([]*models.Pet, error) {
	return s.repo.GetAll()
}

func (s service) GetById(id string) (*models.Pet, error) {
	return s.repo.GetById(id)
}

func (s service) Create(pet *models.Pet) (err error) {
	err = s.repo.Create(pet)
	return
}

func (s service) Update(pet *models.Pet) (err error) {
	err = s.repo.Update(pet)
	return
}

func (s service) DeleteById(id string) (err error) {
	err = s.repo.DeleteById(id)
	return
}

func NewService(repo interfaces.PetRepository) Service {
	return &service{
		repo,
	}
}
