package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	HOST = "database"
	PORT = 5432
)

type Database struct {
	Connection *sql.DB
}

func Init(username, password, database string) (Database, error) {
	db := Database{}
	dataSource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		HOST, PORT, username, password, database)
	connection, err := sql.Open("postgres", dataSource)
	if err != nil {
		return db, err
	}
	db.Connection = connection
	err = db.Connection.Ping()
	if err != nil {
		return db, err
	}

	return db, nil
}
