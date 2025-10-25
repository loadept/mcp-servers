package service

import (
	"github.com/loadept/mcp-servers/internal/domain"
	"github.com/loadept/mcp-servers/internal/repository"
)

type DatabaseInfoService struct {
	repo *repository.DatabaseInfoRepository
}

func NewDatabaseInfoService(repo *repository.DatabaseInfoRepository) *DatabaseInfoService {
	return &DatabaseInfoService{repo: repo}
}

func (s *DatabaseInfoService) GetTableInfo(tableName string) ([]domain.TableInfo, error) {
	return s.repo.GetTableInfo(tableName)
}

func (s *DatabaseInfoService) ListTables(page int, schema string) ([]domain.ListTables, error) {
	return s.repo.ListTables(page, schema)
}
