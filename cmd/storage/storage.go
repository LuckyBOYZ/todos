package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Storage struct {
	user     string
	password string
	database string
	db       *sql.DB
}

func NewStorage(user, password, database string) *Storage {
	return &Storage{
		user:     user,
		password: password,
		database: database,
	}
}

func (s *Storage) OpenConnection() error {
	dsn := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s", s.user, s.password, s.database)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("failed to open connection: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return fmt.Errorf("failed to ping database: %v", err)
	}
	db.SetMaxOpenConns(2)
	db.SetMaxIdleConns(1)
	db.SetConnMaxLifetime(1 * time.Minute)
	s.db = db
	return nil
}

func (s *Storage) CloseConnection() error {
	if s.db != nil {
		err := s.db.Close()
		if err != nil {
			return fmt.Errorf("failed to close connection: %v", err)
		}
		s.db = nil
	}
	return nil
}
