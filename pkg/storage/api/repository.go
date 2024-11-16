package apistore

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gocarina/gocsv"
	"github.com/yuisofull/leetgo/pkg/listcompanies"
	"github.com/yuisofull/leetgo/pkg/listproblems"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Storage struct {
	client *http.Client
}

func NewStorage() *Storage {
	return &Storage{
		client: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func (s *Storage) GetCompanies() ([]listcompanies.Company, error) {
	type response struct {
		Name string `json:"name"`
	}
	/*
		curl -L \
		  -H "Accept: application/vnd.github+json" \
		  -H "X-GitHub-Api-Version: 2022-11-28" \
		  https://api.github.com/repos/yuisofull/leetcode-companywise-interview-questions/contents/
	*/
	req, err := http.NewRequest("GET", "https://api.github.com/repos/yuisofull/leetcode-companywise-interview-questions/contents/", nil)
	if err != nil {
		return nil, errors.New("failed to fetch companies: " + err.Error())
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, errors.New("failed to fetch companies: " + err.Error())
	}
	defer resp.Body.Close()
	var companies []response
	if err := json.NewDecoder(resp.Body).Decode(&companies); err != nil {
		return nil, errors.New("failed to decode response: " + err.Error())
	}
	var result []listcompanies.Company
	for _, c := range companies {
		if !strings.HasSuffix(c.Name, ".csv") {
			continue
		}
		result = append(result, listcompanies.Company{
			Name: strings.TrimSuffix(c.Name, ".csv"),
		})
	}
	return result, nil
}

func (s *Storage) GetProblemsFromCompany(company string) ([]listproblems.Problem, error) {
	/*
		curl -L \                                                                                         [15:49]   %
		  -H "Accept: application/vnd.github+json" \
		  -H "X-GitHub-Api-Version: 2022-11-28" \
		  https://api.github.com/repos/yuisofull/leetcode-companywise-interview-questions/contents/facebook.csv
	*/
	req, err := http.NewRequest("GET", "https://api.github.com/repos/yuisofull/leetcode-companywise-interview-questions/contents/"+company+".csv", nil)
	if err != nil {
		return nil, errors.New("failed to fetch problems: " + err.Error())
	}
	req.Header.Set("Accept", "application/vnd.github.raw+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, errors.New("failed to fetch problems: " + err.Error())
	}
	defer resp.Body.Close()

	var ps []Problem
	if err := gocsv.Unmarshal(resp.Body, &ps); err != nil {
		return nil, errors.New("failed to decode response: " + err.Error())
	}
	var problems []listproblems.Problem
	for _, p := range ps {
		problems = append(problems, listproblems.Problem{
			ID:         p.ID,
			Title:      p.Title,
			URL:        fmt.Sprintf("https://leetcode.com%s", p.URL),
			Difficulty: strings.ToLower(p.Difficulty),
		})
		acc, err := strconv.ParseFloat(strings.TrimSuffix(p.Acceptance, "%"), 64)
		if err != nil {
			return nil, errors.New("failed to parse acceptance rate: " + err.Error())
		}
		problems[len(problems)-1].Acceptance = acc
		freq, err := strconv.ParseFloat(strings.TrimSuffix(p.Frequency, "%;"), 64)
		if err != nil {
			return nil, errors.New("failed to parse frequency rate: " + err.Error())
		}
		problems[len(problems)-1].Frequency = freq
		if p.IsPremium == "N" {
			problems[len(problems)-1].IsPremium = false
		} else {
			problems[len(problems)-1].IsPremium = true
		}
	}
	return problems, nil
}
