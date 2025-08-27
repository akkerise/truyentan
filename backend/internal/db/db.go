package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connect initializes a PostgreSQL connection using GORM.
func Connect(host, port, user, password, name string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, name)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
