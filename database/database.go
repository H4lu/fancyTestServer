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

var DB *sql.DB

func InitDB() (*sql.DB, error) {
	connInfo := fmt.Sprintf("host=%s port=%d user=%s"+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", connInfo)
	if err != nil {
		return nil, err
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(10)

	fmt.Println("Successfully connected to postgres!")

	DB = db
	return DB, nil
}

func GetDB() *sql.DB {
	return DB
}
