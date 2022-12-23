package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type InitDB struct {
	DB   *sql.DB
	Port string
}

func InitDatabase() *InitDB {
	godotenv.Load(".env")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	var (
		host     = "localhost"
		dbPort   = 5432
		user     = os.Getenv("POSTGRES_USER")
		password = os.Getenv("POSTGRES_PASSWORD")
		dbname   = os.Getenv("POSTGRES_DB")
	)

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, dbPort, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return &InitDB{
		DB:   db,
		Port: port,
	}
}
