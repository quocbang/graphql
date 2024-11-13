package main

import (
	"database/sql"
	"log"

	migrate "github.com/rubenv/sql-migrate"
)

func main() {
	migrations := migrate.FileMigrationSource{
		Dir: "./migrations",
	}

	n, err := migrate.Exec(&sql.DB{}, "postgres", migrations, migrate.Up)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Applied %d migration!\n", n)
}
