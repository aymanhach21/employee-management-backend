package models

import "gorm.io/gorm"

// Employee représente un employé dans la base de données.
type Employee struct {
    gorm.Model         // Ajoute automatiquement les champs ID, CreatedAt, UpdatedAt, DeletedAt
    FirstName  string `gorm:"not null"`
    LastName   string `gorm:"not null"`
    Email      string `gorm:"unique;not null"`
    Phone      string `gorm:"not null"`
    Position   string `gorm:"not null"`
    Department string `gorm:"not null"`
    HireDate   string `gorm:"not null"`
}
