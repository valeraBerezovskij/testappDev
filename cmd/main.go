package main

import (
	"log"

	"example.com/taskservice/internal/domain"
	"example.com/taskservice/internal/server"
	"example.com/taskservice/pkg/database"
)

func main() {
	db := database.NewPostgresDB()

	if err := db.AutoMigrate(&domain.Task{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	r := server.NewServer(db)

	r.Run(":8080")
}
