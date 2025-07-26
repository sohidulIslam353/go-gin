// initializes Bun with CockroachDB
package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

var DB *bun.DB

func InitDB() {
	LoadConfig() // Load from config.yaml

	dbConf := AppConfig.Database
	dsn := fmt.Sprintf("postgresql://%s@%s:%d/%s?sslmode=disable",
		dbConf.User, dbConf.Host, dbConf.Port, dbConf.Name)

	sqlDB, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("❌ Failed to open DB: %v", err)
	}

	DB = bun.NewDB(sqlDB, pgdialect.New())

	if err = sqlDB.Ping(); err != nil {
		log.Fatalf("❌ Failed to ping DB: %v", err)
	}

	log.Println("✅ Connected to CockroachDB")
}
