package database

import (
	"database/sql"
	"fmt"
	"log"

	"gin-bun-cockroach/config"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

var DB *bun.DB

func InitDatabase() {
	config.LoadConfig()

	dbConfig := config.AppConfig.Database

	dsn := fmt.Sprintf("postgresql://%s@%s:%d/%s?sslmode=disable",
		dbConfig.User,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Name,
	)

	// ✅ Use sql.OpenDB with pgdriver.NewConnector
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	// ✅ Pass *sql.DB to bun.NewDB
	DB = bun.NewDB(sqldb, pgdialect.New())

	log.Println("✅ Connected to CockroachDB")
}
