package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func main() {

}

// connectToDB is a simple connect to database function so we can demo using service containers in GitHub Actions.
func connectToDB() (bool, error) {
	dbHost, found := os.LookupEnv("DB_HOSTNAME")
	if !found {
		return false, fmt.Errorf("DB_HOSTNAME not set")
	}
	dbUsername, found := os.LookupEnv("DB_USERNAME")
	if !found {
		return false, fmt.Errorf("DB_USERNAME not set")
	}
	dbPassword, found := os.LookupEnv("DB_PASSWORD")
	if !found {
		return false, fmt.Errorf("DB_PASSWORD not set")
	}

	conn, err := pgx.Connect(context.Background(), fmt.Sprintf("postgres://%s:%s@%s:5432/postgres", dbUsername, dbPassword, dbHost))
	if err != nil {
		return false, fmt.Errorf("unable to connect to database: %v", err)
	}

	err = conn.Ping(context.Background())
	if err != nil {
		return false, fmt.Errorf("unable to ping database: %v", err)
	}

	return true, nil
}
