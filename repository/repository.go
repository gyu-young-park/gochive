package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Storer struct {
	db *sql.DB
}

func NewStorer(dsn string) *Storer {
	storer := &Storer{}
	db, err := openDB(dsn)
	if err != nil {
		fmt.Printf("error occured:%s\n", err.Error())
		return nil
	}
	storer.db = db
	return storer
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func (s *Storer) Close() {
	defer s.Close()
}
