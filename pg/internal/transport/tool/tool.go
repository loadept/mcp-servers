package tool

import "github.com/loadept/mcp-servers/pg/internal/service"

type Tool struct {
	GetTableInfo *GetTableInfo
	ListTables   *ListTables
	ExecuteQuery *ExecuteQuery
}

func GetTools(queryService *service.QueryService, databaseInfoService *service.DatabaseInfoService) *Tool {
	return &Tool{
		GetTableInfo: &GetTableInfo{databaseInfoService: databaseInfoService},
		ListTables:   &ListTables{databaseInfoService: databaseInfoService},
		ExecuteQuery: &ExecuteQuery{queryService: queryService},
	}
}
