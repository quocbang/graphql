package main

import (
	"log"

	"github.com/quocbang/graphql/adapters/infra/repositories"
	"github.com/quocbang/graphql/pkg/config"
	migrate "github.com/rubenv/sql-migrate"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := repositories.NewDBConnection(repositories.DBConfig{
		Host:         cfg.DB.Host,
		Port:         cfg.DB.Port,
		UserName:     cfg.DB.User,
		Password:     cfg.DB.Password,
		DatabaseName: cfg.DB.Name,
	})
	if err != nil {
		log.Fatal(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	migrations := migrate.FileMigrationSource{
		Dir: "./migrations",
	}

	n, err := migrate.Exec(sqlDB, "postgres", migrations, migrate.Up)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Applied %d migration!\n", n)
}
