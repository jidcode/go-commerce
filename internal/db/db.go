package db

import (
	"log"

	"github.com/jidcode/go-commerce/internal/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func ConnectDB(cfg config.Variable) *sqlx.DB {
	connString := cfg.DBUrl

	DB, err := sqlx.Connect("postgres", connString)
	if err != nil {
		log.Fatal("Error connecting to database:  %w", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Error pinging the database: %w", err)
	}

	defer DB.Close()


	log.Printf("Connection to database established successfully!")
	return DB
}