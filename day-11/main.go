package main

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	// Data Source Name
	// hardcoded, bad for production grade
	// dsn := "host=localhost port=5435 user=postgres password=password dbname=postgres sslmode=disable"

	// DSN can be set when running application
	// DSN="host=localhost port=5435 user=postgres password=password dbname=db1 sslmode=disable" go run main.go

	// better, lookup env variable
	// set on your shell (.zshrc, .bashrc, etc)
	// export DSN="..."
	dsn := os.Getenv("DSN")
	if dsn == "" {
		log.Fatal("DSN not found!")
	}

	// create pool
	db, err := sql.Open("postgres", dsn)

	// pool config
	db.SetMaxOpenConns(5)                  // max active connections
	db.SetMaxIdleConns(2)                  // keep 2 connection ready
	db.SetConnMaxLifetime(5 * time.Minute) // recycle connection every 5 minute

	if err != nil {
		log.Fatalf("failed to open db: %v", err)
	}

	defer db.Close()

	// connect from pool using db.Ping()
	if err := db.Ping(); err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}

	log.Println("connected to postgresql")

	userId := 1

	// `$1` parameter placeholder to avoid SQL Injection
	row := db.QueryRow("select id, name, is_active, weight from users where id = $1", userId)

	var u User
	// map record to struct
	errScan := row.Scan(&u.ID, &u.Name, &u.IsActive, &u.Weight) // why use `&` because Scan will modify value via reference

	if errScan == sql.ErrNoRows {
		log.Fatalf("user %d not found", 1)
	}

	if errScan != nil {
		log.Fatalf("fallback error: %v", errScan)
	}

	log.Printf("user: %+v", u)
}
