package database

import (
	"fmt"
	"log"

	"dnd-sheet-api/internal/config"
	"dnd-sheet-api/internal/domain/character"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect(configuration *config.Config) *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		configuration.DBUser,
		configuration.DBPassword,
		configuration.DBHost,
		configuration.DBPort,
		configuration.DBName,
	)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	if err := database.AutoMigrate(&character.Character{}); err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}

	log.Println("database connection established")

	return database
}
