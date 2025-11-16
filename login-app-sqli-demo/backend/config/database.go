package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	Connection *sql.DB
}

func NewDatabase() *Database {
	// Database connnection params
	username := "root"
	password := "root"
	hostname := "127.0.0.1"
	port := "3306"
	database := "login_app_sqli_demo"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", username, password, hostname, port, database)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	log.Println("Database connection established!")

	return &Database{
		Connection: db,
	}
}

func (d *Database) Close() {
	d.Connection.Close()
}
