package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func GetDB() *sql.DB {

	switch ConfigVar.Database.TypeDB {
	case "postgres":
		return getPostgres()
	case "mysql":
		return getMysql()
	default:
		return nil
	}

}

func getPostgres() *sql.DB {

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

func getMysql() *sql.DB {

	connStr := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s",
		ConfigVar.Database.User, ConfigVar.Database.Password,
		ConfigVar.Database.Host, ConfigVar.Database.DBName)

	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
