package di

import (
	"database/sql"

	"github.com/loadept/mcp-servers/pg/internal/repository"
	"github.com/loadept/mcp-servers/pg/internal/service"
)

type Container struct {
	DatabaseInfoRepository *repository.DatabaseInfoRepository
	QueryRepository        *repository.QueryRepository

	DatabaseInfoService *service.DatabaseInfoService
	QueryService        *service.QueryService
}

func NewContainer(db *sql.DB) *Container {
	dbInfoRepo := repository.NewDatabaseInfoRepository(db)
	queryRepo := repository.NewQueryRepository(db)

	dbInfoService := service.NewDatabaseInfoService(dbInfoRepo)
	queryService := service.NewQueryService(queryRepo)

	return &Container{
		DatabaseInfoRepository: dbInfoRepo,
		QueryRepository:        queryRepo,
		DatabaseInfoService:    dbInfoService,
		QueryService:           queryService,
	}
}
