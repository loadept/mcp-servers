package persistence

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/lib/pq"
	"github.com/loadept/mcp-servers/internal/config"
)

var (
	once     sync.Once
	instance *postgres
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

func (s *postgres) getNow() (string, error) {
	var now string

	if err := s.db.QueryRow("SELECT CURRENT_DATE::VARCHAR").Scan(&now); err != nil {
		return "", err
	}

	return now, nil
}

func NewDBPostgres() (*postgres, error) {
	var err error
	var now string

	once.Do(func() {
		postgresDB := &postgres{}
		if err = postgresDB.Connect(); err == nil {
			now, err = postgresDB.getNow()
			if err == nil {
				instance = postgresDB
				log.Printf("Connected to PostgreSQL database, current date %s\n", now)
			}
		}
	})
	return instance, err
}

func (s *postgres) GetDB() *sql.DB {
	return s.db
}

func (s *postgres) Close() error {
	return s.db.Close()
}
