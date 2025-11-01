package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/loadept/mcp-servers/internal/config"
	"github.com/loadept/mcp-servers/internal/di"
	"github.com/loadept/mcp-servers/internal/infra/persistence"
	"github.com/loadept/mcp-servers/internal/transport"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func init() {
	config.LoadEnvs()
}

func main() {
	pg, err := persistence.NewDBPostgres()
	if err != nil {
		fmt.Println("An error occurred while connecting to the database:", err)
		os.Exit(1)
	}
	defer pg.Close()
	db := pg.GetDB()

	implementation := &mcp.Implementation{
		Name:    "PostgreSQL MCP Server",
		Version: "0.1.1",
	}
	server := mcp.NewServer(implementation, nil)

	containerDependencies := di.NewContainer(db)
	tool := transport.NewMCPTransport(containerDependencies.QueryService, containerDependencies.DatabaseInfoService)

	mcp.AddTool(server, &mcp.Tool{Name: "execute_query", Description: "Executes a query on the postgres database"}, tool.ExecuteQuery)
	mcp.AddTool(server, &mcp.Tool{Name: "get_table_info", Description: "Get information about a table"}, tool.GetTableInfo)
	mcp.AddTool(server, &mcp.Tool{Name: "list_tables", Description: "List all available tables in a schema"}, tool.ListTables)

	log.Println("MCP server is running...")
	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
