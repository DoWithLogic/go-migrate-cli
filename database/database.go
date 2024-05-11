package database

import (
	"database/sql"
	"fmt"
)

func Open(dsn string) *sql.DB {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("Error Opening DB: %v \n", err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Printf("Error Pinging DB: %v \n", err)
	}

	fmt.Println("Connected to db!")

	return db
}
