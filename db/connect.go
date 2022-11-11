package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func NewConnect(host, name, user, password string, port uint) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable", host, port, name, user, password)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
