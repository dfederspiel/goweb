package pet

type Service interface {
	GetAll() ([]*Pet, error)
	GetById(id string) (*Pet, error)
	Create(pet *Pet) (err error)
	Update(pet *Pet) (err error)
	DeleteById(id string) (err error)
}

type service struct {
	repo Repository
}

func (s service) GetAll() ([]*Pet, error) {
	return s.repo.GetAll()
}

func (s service) GetById(id string) (*Pet, error) {
	return s.repo.GetById(id)
}

func (s service) Create(pet *Pet) (err error) {
	s.repo.Create(pet)
	return
}

func (s service) Update(pet *Pet) (err error) {
	s.repo.Update(pet)
	return
}

func (s service) DeleteById(id string) (err error) {
	s.repo.DeleteById(id)
	return
}

func NewService(repo Repository) Service {
	return &service{
		repo,
	}
}
