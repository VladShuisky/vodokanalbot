package utils

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/VladShuisky/vodokanalbot/database"
)

func CheckDbConnect() string {
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	fmt.Println(connStr)

	db, err := database.NewPostgres(connStr)

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	var version string

	err = db.Pool.QueryRow(context.Background(), "SELECT version()").Scan(&version)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("PostgreSQL version:", version)
	return version
}