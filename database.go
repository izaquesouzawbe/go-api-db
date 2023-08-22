package main

import (
	"database/sql"
	"fmt"
	"log"
)

func getDB() *sql.DB {

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=%s",
		config.Database.User, config.Database.Password, config.Database.DBName,
		config.Database.Host, config.Database.SSLMode)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	return db
}
