package db

import (
	"database/sql"
	"log/slog"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func DatabaseConnection() (*sql.DB, error) {
	err := godotenv.Load("C:/Users/joaog/DEV/coffee_shop/back/.env")
	if err != nil {
		slog.Error("error on load the .env")
		return nil, err
	}

	databaseCredentials := os.Getenv("DATABASE_CREDENTIALS")
	db, err := sql.Open("mysql", databaseCredentials)
	if err != nil {
		slog.Error("error on open a connection with database")
		return nil, err
	}

	return db, nil
}
