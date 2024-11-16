package listproblems

type Repository interface {
	GetProblemsFromCompany(company string) ([]Problem, error)
}

type Service interface {
	GetProblemsFromCompany(company string, filter Filter) ([]Problem, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetProblemsFromCompany(company string, filter Filter) ([]Problem, error) {
	problems, err := s.repo.GetProblemsFromCompany(company)
	if err != nil {
		return nil, err
	}

	return problems, nil
}
