package config

import (
    "fmt"
    "log"
    "os"

    "github.com/jmoiron/sqlx"
    _ "github.com/lib/pq"
)

var DB *sqlx.DB

func Connect() {
    dbURL := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        getEnv("DB_HOST", "localhost"),
        getEnv("DB_PORT", "5432"),
        getEnv("DB_USER", "postgres"),
        getEnv("DB_PASSWORD", "postgres"),
        getEnv("DB_NAME", "library"))

    var err error
    DB, err = sqlx.Connect("postgres", dbURL)
    if err != nil {
        log.Fatalf("Failed to connect to DB: %v", err)
    }
}

func getEnv(key, defaultVal string) string {
    if val, exists := os.LookupEnv(key); exists {
        return val
    }
    return defaultVal
}