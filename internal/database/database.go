package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file" // required for file:// protocol
	"github.com/spf13/viper"
)

var DB *sql.DB

func getDatabaseUrl() string {
	databaseUrl := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		viper.GetString("DB_USER"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_HOST"),
		viper.GetInt("DB_PORT"),
		viper.GetString("DB_NAME"),
	)
	return databaseUrl
}

func MigrateUp() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	m, err := migrate.New(fmt.Sprintf("file://%s/config/migrations", dir), getDatabaseUrl())
	if err != nil {
		log.Fatalln(err)
	}
	err = m.Up()
	if err != nil {
		log.Fatalln(err)
	}
}

func MigrateDown() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	m, err := migrate.New(fmt.Sprintf("file://%s/config/migrations", dir), getDatabaseUrl())
	if err != nil {
		log.Fatalln(err)
	}
	err = m.Steps(-1)
	if err != nil {
		log.Fatalln(err)
	}
}
