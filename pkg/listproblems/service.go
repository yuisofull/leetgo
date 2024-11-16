package listproblems

import (
	"sort"
)

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
	filter.Process()
	switch {
	case filter.ByFrequency:
		sort.Slice(problems, func(i, j int) bool {
			return problems[i].Frequency > problems[j].Frequency
		})
	case filter.ByAcceptance:
		sort.Slice(problems, func(i, j int) bool {
			return problems[i].Acceptance < problems[j].Acceptance
		})
	}
	if filter.ByDifficulty != "" {
		var difficultyProblems []Problem
		for _, p := range problems {
			if p.Difficulty == filter.ByDifficulty {
				difficultyProblems = append(difficultyProblems, p)
			}
		}
		problems = difficultyProblems
	}
	if filter.IsNotPremium {
		var nonPremiumProblems []Problem
		for _, problem := range problems {
			if !problem.IsPremium {
				nonPremiumProblems = append(nonPremiumProblems, problem)
			}
		}
		problems = nonPremiumProblems
	}
	if filter.Limit < len(problems) {
		problems = problems[:filter.Limit]
	}
	return problems, nil
}
