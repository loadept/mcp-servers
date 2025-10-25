package transport

import (
	"context"
	"fmt"

	"github.com/loadept/mcp-servers/internal/domain"
	"github.com/loadept/mcp-servers/internal/service"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type MCPTransport struct {
	queryService        *service.QueryService
	databaseInfoService *service.DatabaseInfoService
}

func NewMCPTransport(queryService *service.QueryService, databaseInfoService *service.DatabaseInfoService) *MCPTransport {
	return &MCPTransport{queryService: queryService, databaseInfoService: databaseInfoService}
}

func (m *MCPTransport) ExecuteQuery(ctx context.Context, req *mcp.CallToolRequest, input domain.QueryToolInput) (
	*mcp.CallToolResult,
	domain.QueryToolOutput,
	error,
) {
	results, err := m.queryService.ExecuteQuery(input.Query, input.Args...)
	if err != nil {
		output := domain.QueryToolOutput{
			Detail:   fmt.Sprintf("An error occurred while executing the query: %v", err.Error()),
			RowCount: 0,
			Results:  []map[string]any{},
		}
		return nil, output, err
	}

	output := domain.QueryToolOutput{
		Detail:   "Query executed successfully.",
		RowCount: len(results),
		Results:  results,
	}
	return nil, output, nil
}

func (m *MCPTransport) GetTableInfo(ctx context.Context, req *mcp.CallToolRequest, input domain.TableInfoInput) (
	*mcp.CallToolResult,
	domain.TableInfoOutput,
	error,
) {
	results, err := m.databaseInfoService.GetTableInfo(input.TableName)
	if err != nil {
		output := domain.TableInfoOutput{
			Detail:  fmt.Sprintf("An error occurred while retrieving table information: %v", err.Error()),
			Results: []domain.TableInfo{},
		}
		return nil, output, err
	}

	output := domain.TableInfoOutput{
		Detail:  "The table information has been retrieved successfully.",
		Results: results,
	}
	return nil, output, nil
}

func (m *MCPTransport) ListTables(ctx context.Context, req *mcp.CallToolRequest, input domain.ListTablesInput) (
	*mcp.CallToolResult,
	domain.ListTablesOutput,
	error,
) {
	results, err := m.databaseInfoService.ListTables(input.Page, input.Schema)
	if err != nil {
		output := domain.ListTablesOutput{
			Detail:  fmt.Sprintf("An error occurred while listing tables: %v", err.Error()),
			Results: []domain.ListTables{},
		}
		return nil, output, err
	}

	output := domain.ListTablesOutput{
		Detail:  "The table listing has been executed successfully.",
		Results: results,
	}
	return nil, output, nil
}
