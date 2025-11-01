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
	"github.com/loadept/mcp-servers/internal/transport/tool"
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
		Version: "0.2.0",
	}
	server := mcp.NewServer(implementation, nil)

	containerDependencies := di.NewContainer(db)
	tool := tool.GetTools(
		containerDependencies.QueryService,
		containerDependencies.DatabaseInfoService,
	)

	transport.LoadTool(server, tool.GetTableInfo)
	transport.LoadTool(server, tool.ListTables)
	transport.LoadTool(server, tool.ExecuteQuery)

	log.Println("MCP server is running...")
	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
