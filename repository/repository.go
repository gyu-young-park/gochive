package repository

import "database/sql"

type Storer struct {
	db *sql.DB
}
