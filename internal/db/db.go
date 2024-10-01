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

	var err error

	DB, err = sqlx.Connect("postgres", connString)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("Error pinging the database: %v", err)
	}

	log.Printf("Connection to database established successfully!")
	return DB
}
