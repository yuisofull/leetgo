package listcompanies

type Repository interface {
	GetCompanies() ([]Company, error)
}

type Service interface {
	GetCompanies() ([]Company, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetCompanies() ([]Company, error) {
	return s.repo.GetCompanies()
}
