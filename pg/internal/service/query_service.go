package service

import (
	"context"
	"fmt"

	"github.com/loadept/mcp-servers/internal/repository"
)

type QueryService struct {
	repo *repository.QueryRepository
}

func NewQueryService(repo *repository.QueryRepository) *QueryService {
	return &QueryService{repo: repo}
}

func (s *QueryService) ExecuteQuery(ctx context.Context, query string, args ...any) ([]map[string]any, error) {
	if query == "" {
		return nil, fmt.Errorf("query cannot be empty")
	}

	data, err := s.repo.ExecuteQuery(ctx, query, args...)
	if err != nil {
		return []map[string]any{}, err
	}
	return data, nil
}
