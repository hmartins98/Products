package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func GetPostgresDB() (db *sql.DB, err error) {
	dbHost := "192.168.1.100"
	dbPort := 5432
	dbUser := "postgres"
	dbPass := "password"
	dbName := "Payout"
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPass, dbName)

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
	return
}
