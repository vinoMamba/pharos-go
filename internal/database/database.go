package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file" // required for file:// protocol
)

var DB *sql.DB

const (
	host     = "192.168.110.75"
	port     = 5432
	user     = "pharos"
	password = "123456"
	dbName   = "pharos_dev"
)

func MigrateUp() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(dir)
	m, err := migrate.New(
		fmt.Sprintf("file://%s/config/migrations", dir),
		fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user, password, host, port, dbName),
	)
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
	m, err := migrate.New(
		fmt.Sprintf("file://%s/config/migrations", dir),
		fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user, password, host, port, dbName),
	)
	if err != nil {
		log.Fatalln(err)
	}
	err = m.Steps(-1)
	if err != nil {
		log.Fatalln(err)
	}
}
