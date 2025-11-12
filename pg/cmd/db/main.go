package main

import (
	"github.com/loadept/mcp-servers/pg/internal/config"
	"github.com/loadept/mcp-servers/pg/internal/infra/persistence"
)

func init() {
	config.LoadEnvs()
}

func main() {
	pg, err := persistence.NewDBPostgres()
	if err != nil {
		panic(err)
	}
	defer pg.Close()
}
