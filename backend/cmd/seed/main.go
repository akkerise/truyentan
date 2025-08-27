package main

import (
	"log"
	"os"

	internaldb "github.com/truyentan/backend/internal/db"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	if err := internaldb.Migrate(db); err != nil {
		log.Fatal(err)
	}
	if err := internaldb.Seed(db); err != nil {
		log.Fatal(err)
	}
}
