package config

import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
	"log"
)

const (
	DB_HOST = "localhost"
	DB_PORT = "5434"
	DB_USER = "postgres"
	DB_PASSWORD = "password"
	DB_NAME = "testserver"
)

var db *sql.DB

func InitDB() (*sql.DB, error) {
	var err error
	dbinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Successfully connected!")
	}
	return db, nil
}

func main() {
	InitDB()
	defer db.Close()
}