package usecase

import (
	"database/sql"
	"service/internal/module/products/repository"
)

type products struct {
	repository Repository
	dependency dependency
}

type dependency struct {}

func Initialize(dbread, dbwrite *sql.DB) *products {
	return &products{
		repository: repository.NewProducts(dbread,dbwrite),
	}
}

type Usecase interface {
	Check() (string, error)
}

type Repository interface {}