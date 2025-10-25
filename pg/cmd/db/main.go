package main

import (
	"log"

	"github.com/loadept/mcp-servers/internal/config"
	"github.com/loadept/mcp-servers/internal/infra/persistence"
)

func init() {
	config.LoadEnvs()
}

func main() {
	pg := persistence.NewDBPostgres()

	if err := pg.Connect(); err != nil {
		panic(err)
	}
	defer pg.Close()

	now, err := pg.GetNow()
	if err != nil {
		panic(err)
	}
	log.Printf("Connection with database success, current date %s\n", now)
}
