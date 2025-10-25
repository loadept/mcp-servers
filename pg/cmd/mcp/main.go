package main

import (
	"context"
	"log"

	"github.com/loadept/mcp-servers/internal/config"
	"github.com/loadept/mcp-servers/internal/di"
	"github.com/loadept/mcp-servers/internal/infra/persistence"
	"github.com/loadept/mcp-servers/internal/transport"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func init() {
	config.LoadEnvs()
}

func ToolEjm() (*mcp.Server, *mcp.Tool, mcp.ToolHandlerFor[any, any]) {
	return &mcp.Server{}, &mcp.Tool{}, nil
}

func main() {
	pg := persistence.NewDBPostgres()

	if err := pg.Connect(); err != nil {
		panic(err)
	}
	defer pg.Close()

	// now, err := pg.GetNow()
	// if err != nil {
	// 	panic(err)
	// }
	// log.Printf("Connection with database success, current date %s\n", now)

	db := pg.GetDB()

	implementation := &mcp.Implementation{
		Name:    "PostgreSQL MCP Server",
		Version: "0.1.0",
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
