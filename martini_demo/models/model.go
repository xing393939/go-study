package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Model interface {
	CreateTableSQL() string
}

type DB struct {
	SQLDB *sql.DB
}

func OpenDB(driverName, dataSourceName string) (*DB, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(8)
	return &DB{SQLDB: db}, nil
}

type Scannable interface {
	Scan(dest ...interface{}) error
}