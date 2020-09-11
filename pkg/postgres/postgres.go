package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
	"time"
)

func New(dsn string, maxopen, maxidle, timeout int) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Duration(timeout) * time.Second)
	db.SetMaxOpenConns(maxopen)
	db.SetMaxIdleConns(maxidle)

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
