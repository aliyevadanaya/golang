package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func InitPostgres() *sql.DB {
	connStr := "user=danaya_user password=DanayaKrasotka dbname=practice5 host=localhost port=5432 sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Error connecting to database", err)
	}

	log.Println("Successfully connected to database")

	return db
}
