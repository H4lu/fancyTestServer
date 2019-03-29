package database

import (
	"database/sql"
	"fmt"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "demo"
)

type Database struct {
	db *sql.DB
}

var DB Database

func InitDB() (Database, error) {
	connInfo := fmt.Sprintf("host=%s port=%d user=%s"+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", connInfo)
	if err != nil {
		return DB, err
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		return DB, err
	}

	fmt.Println("Successfully connected to postgres!")

	DB.db = db
	return DB, nil
}

func (db Database) saveUser() error {

}
