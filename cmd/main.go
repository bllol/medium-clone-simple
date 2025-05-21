package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./myapp.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Create a sample table
	createTable := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        email TEXT UNIQUE NOT NULL
    );
    `
	_, err = db.Exec(createTable)
	if err != nil {
		panic(err)
	}

	fmt.Println("Database and table created successfully!")
}
