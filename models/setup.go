package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB // Variable globale pour la connexion à la base de données

// ConnectDatabase initialise la connexion à la base de données
func ConnectDatabase() {
	dsn := "host=localhost user=postgres password=1234 dbname=employee_management port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	// Affecter la connexion à la variable globale
	DB = database

	// Migrer les modèles (crée les tables si elles n'existent pas)
	err = database.AutoMigrate(&Employee{})
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}
}
