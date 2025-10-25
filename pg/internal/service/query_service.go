package service

import (
	"fmt"
	"strings"

	"github.com/loadept/mcp-servers/internal/repository"
)

type QueryService struct {
	repo *repository.QueryRepository
}

func NewQueryService(repo *repository.QueryRepository) *QueryService {
	return &QueryService{repo: repo}
}

func (s *QueryService) ExecuteQuery(query string, args ...any) ([]map[string]any, error) {
	if query == "" {
		return nil, fmt.Errorf("query cannot be empty")
	}
	if !strings.HasPrefix(strings.ToUpper(strings.TrimSpace(query)), "SELECT") {
		return nil, fmt.Errorf("only SELECT queries are allowed")
	}

	data, err := s.repo.ExecuteQuery(query, args...)
	if err != nil {
		return []map[string]any{}, err
	}
	return data, nil
}
