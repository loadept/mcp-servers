package persistence

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/loadept/mcp-servers/internal/config"
)

type postgres struct {
	db *sql.DB
}

func (s *postgres) Connect() error {
	pgURI := config.GetEnv("POSTGRES_URI")

	db, err := sql.Open("postgres", pgURI)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db
	return nil
}

func (s *postgres) GetNow() (string, error) {
	var now string

	if err := s.db.QueryRow("SELECT TO_CHAR(TODAY) FROM systables WHERE tabid = 1").Scan(&now); err != nil {
		return "", err
	}

	return now, nil
}

func NewDBPostgres() *postgres {
	return &postgres{}
}

func (s *postgres) GetDB() *sql.DB {
	return s.db
}

func (s *postgres) Close() error {
	return s.db.Close()
}
