package main

import (
    "employee-management/controllers" // Contrôleurs des endpoints
    "employee-management/models"      // Modèles et connexion DB
    "employee-management/docs"        // Documentation générée par Swagger

    "fmt"                              // Pour les messages dans la console
    "github.com/gin-gonic/gin"         // Gin framework
    ginSwagger "github.com/swaggo/gin-swagger" // Swagger pour Gin
    swaggerFiles "github.com/swaggo/files"    // Nouveau package pour swaggerFiles
)

func main() {
    // Initialiser la base de données
    models.ConnectDatabase()

    fmt.Println("Database connection and migration successful!")

    // Charger la documentation Swagger
    docs.SwaggerInfo.Title = "Employee Management API"
    docs.SwaggerInfo.Description = "API for managing employees"
    docs.SwaggerInfo.Version = "1.0"
    docs.SwaggerInfo.Host = "localhost:8080"
    docs.SwaggerInfo.BasePath = "/"
    docs.SwaggerInfo.Schemes = []string{"http"}

    // Initialiser Gin
    r := gin.Default()

    // Ajouter les routes Swagger
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    // Définir les routes API
    r.GET("/employees", controllers.GetEmployees)             // Lister tous les employés
    r.POST("/employees", controllers.CreateEmployee)          // Ajouter un employé
    r.GET("/employees/:id", controllers.GetEmployeeByID)      // Récupérer un employé par ID
    r.PUT("/employees/:id", controllers.UpdateEmployee)       // Mettre à jour un employé
    r.DELETE("/employees/:id", controllers.DeleteEmployee)    // Supprimer un employé
    r.GET("/employees/search", controllers.SearchEmployees)   // Rechercher et filtrer les employés

    // Lancer le serveur
    r.Run(":8080")
}
