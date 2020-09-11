package repository

import "database/sql"

type products struct{
	dbread, dbwrite *sql.DB
}

func NewProducts(dbread, dbwrite *sql.DB) *products {
	return &products{dbread, dbwrite}
}