package main

import (
	"log"
	"os"

	"github.com/truyentan/backend/internal/app"
	internaldb "github.com/truyentan/backend/internal/db"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// @title Truyentan API
// @version 1.0
// @description Truyentan service API
// @BasePath /api/v1
func main() {
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	if err := internaldb.Migrate(db); err != nil {
		log.Fatal(err)
	}

	srv := app.NewServer()
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
