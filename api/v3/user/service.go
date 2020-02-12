package user

type Service interface {
	GetByEmail(email string) (User, error)
}

type service struct {
	repo Repository
}

func (s service) GetByEmail(email string) (User, error) {
	return s.repo.GetByEmail(email)
}

func NewService(repo Repository) Service {
	return &service{repo}
}
