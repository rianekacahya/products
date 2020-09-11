package postgres

import (
	"database/sql"
	"github.com/fortytw2/dockertest"
	"testing"
)

func TestNew(t *testing.T) {
	container, err := dockertest.RunContainer("postgres:alpine", "5432", func(addr string) error {
		db, err := sql.Open("postgres", "postgres://postgres:postgres@"+addr+"?sslmode=disable")
		if err != nil {
			return err
		}

		return db.Ping()
	}, "-e", "POSTGRES_PASSWORD=postgres", "-e", "POSTGRES_USER=postgres")
	defer container.Shutdown()
	if err != nil {
		t.Fatalf("could not start postgres, %s", err)
	}

	_, err = New("PSN", 10, 5, 3600)
	if err == nil {
		t.Error("Expected error")
	}

	_, err = New("postgres://postgres:postgres@localhost:1234/postgres?sslmode=disable", 10, 5, 3600)
	if err == nil {
		t.Error("Expected error")
	}

	_, err = New("postgres://postgres:postgres@"+container.Addr+"/postgres?sslmode=disable", 10, 5, 3600)
	if err != nil {
		t.Fatalf("Error establishing connection %v", err)
	}
}
