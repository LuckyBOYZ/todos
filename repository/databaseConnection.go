package repository

import (
	"database/sql"
	"fmt"
	"github.com/LuckyBOYZ/todos/configuration"
	_ "github.com/lib/pq"
	"log"
	"os/exec"
	"strings"
	"time"
)

func NewDatabaseConnection() (*sql.DB, error) {
	return createDbConnection()
}

func createDbConnection() (*sql.DB, error) {
	username := configuration.GetString("username")
	password := configuration.GetString("password")
	database := configuration.GetString("database")
	stringConn := fmt.Sprintf("postgres://%s:%s@localhost:5432/%s?sslmode=disable", username, password, database)
	db, err := sql.Open("postgres", stringConn)
	if err != nil {
		return nil, fmt.Errorf("error occurred while connecting to database: %s", err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("postgres database is not running. trying to run it in docker")
		err := runDockerContainer()
		if err != nil {
			return nil, err
		}
		err = waitForPostgres()
		if err != nil {
			return nil, err
		}
	}
	db.SetMaxOpenConns(2)
	db.SetMaxIdleConns(1)
	db.SetConnMaxLifetime(1 * time.Minute)
	return db, nil
}

func runDockerContainer() error {
	cmd := exec.Command("docker", "compose", "up", "-d")
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to run docker compose: %w", err)
	}
	fmt.Println("docker compose is running")
	return nil
}

func waitForPostgres() error {
	dbContainerName := configuration.GetString("dbContainerName")
	timeoutChan := time.After(10 * time.Second)
	ticker := time.Tick(1 * time.Second)
	attempt := 1
	for {
		select {
		case <-timeoutChan:
			return fmt.Errorf("timeout reached: could not confirm that PostgreSQL is ready")
		case <-ticker:
			cmd := exec.Command("docker", "exec", dbContainerName, "pg_isready")
			output, err := cmd.Output()
			if err == nil && strings.Contains(string(output), "accepting connections") {
				var suffix string
				switch attempt {
				case 1:
					suffix = "st"
				case 2:
					suffix = "nd"
				case 3:
					suffix = "rd"
				default:
					suffix = "th"
				}
				fmt.Printf("PostgreSQL is ready at %d%s attempt\n", attempt, suffix)
				return nil
			}
			log.Printf("Attempt %d: Waiting for PostgreSQL to be ready...\n", attempt)
		}
	}
}
