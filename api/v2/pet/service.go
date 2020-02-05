package pet

type Service interface {
	GetAllPets() ([]*Pet, error)
	GetPetById(id string) (*Pet, error)
	CreatePet(pet *Pet) (err error)
	UpdatePet(pet *Pet) (err error)
	DeletePetById(id string) (err error)
}

type service struct {
	repo Repository
}

func (s service) GetAllPets() ([]*Pet, error) {
	return s.repo.GetAll()
}

func (s service) GetPetById(id string) (*Pet, error) {
	return s.repo.GetById(id)
}

func (s service) CreatePet(pet *Pet) (err error) {
	s.repo.Create(pet)
	return
}

func (s service) UpdatePet(pet *Pet) (err error) {
	s.repo.Update(pet)
	return
}

func (s service) DeletePetById(id string) (err error) {
	s.repo.DeleteById(id)
	return
}

func NewService(repo Repository) Service {
	return &service{
		repo,
	}
}
