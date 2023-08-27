package config

import (
	"database/sql"
	"fmt"
	"log"
)

func GetDB() *sql.DB {

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=%s",
		ConfigVar.Database.User, ConfigVar.Database.Password, ConfigVar.Database.DBName,
		ConfigVar.Database.Host, ConfigVar.Database.SSLMode)

	fmt.Println(connStr)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
