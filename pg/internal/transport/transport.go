package transport

import "github.com/modelcontextprotocol/go-sdk/mcp"

type MCPTransport[In, Out any] interface {
	MCPTool() (metadata *mcp.Tool, handler mcp.ToolHandlerFor[In, Out])
}

func LoadTool[In, Out any](server *mcp.Server, tool MCPTransport[In, Out]) {
	metadata, handler := tool.MCPTool()
	mcp.AddTool(server, metadata, handler)
}
