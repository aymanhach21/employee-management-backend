package controllers

import (
	"employee-management/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetEmployees godoc
// @Summary List all employees
// @Description Get a list of all employees in the database
// @Tags employees
// @Accept json
// @Produce json
// @Success 200 {array} models.Employee
// @Router /employees [get]


func GetEmployees(c *gin.Context) {
	var employees []models.Employee
	if err := models.DB.Find(&employees).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, employees)
}


// CreateEmployee godoc
// @Summary Add a new employee
// @Description Create a new employee with the provided details
// @Tags employees
// @Accept json
// @Produce json
// @Param employee body models.Employee true "Employee details"
// @Success 201 {object} models.Employee
// @Router /employees [post]

func CreateEmployee(c *gin.Context) {
	var employee models.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.DB.Create(&employee).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, employee)
}

// GetEmployeeByID godoc
// @Summary Get an employee by ID
// @Description Retrieve a specific employee by their ID
// @Tags employees
// @Accept json
// @Produce json
// @Param id path int true "Employee ID"
// @Success 200 {object} models.Employee
// @Router /employees/{id} [get]

func GetEmployeeByID(c *gin.Context) {
	id := c.Param("id")
	var employee models.Employee
	if err := models.DB.First(&employee, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}
	c.JSON(http.StatusOK, employee)
}

// UpdateEmployee godoc
// @Summary Update an employee
// @Description Update an employee's details by their ID
// @Tags employees
// @Accept json
// @Produce json
// @Param id path int true "Employee ID"
// @Param employee body models.Employee true "Updated employee details"
// @Success 200 {object} models.Employee
// @Router /employees/{id} [put]

func UpdateEmployee(c *gin.Context) {
	id := c.Param("id")
	var employee models.Employee
	// Récupérer l'employé à mettre à jour
	if err := models.DB.First(&employee, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}
	// Mettre à jour les champs
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.DB.Save(&employee).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, employee)
}
// DeleteEmployee godoc
// @Summary Delete an employee
// @Description Remove an employee by their ID
// @Tags employees
// @Accept json
// @Produce json
// @Param id path int true "Employee ID"
// @Success 200 {string} string "Employee deleted successfully"
// @Router /employees/{id} [delete]

func DeleteEmployee(c *gin.Context) {
	id := c.Param("id")
	var employee models.Employee
	// Récupérer l'employé à supprimer
	if err := models.DB.First(&employee, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}
	if err := models.DB.Delete(&employee).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Employee deleted successfully"})
}



// SearchEmployees godoc
// @Summary Search for employees
// @Description Search employees by first name, department, or position, with optional pagination
// @Tags employees
// @Accept json
// @Produce json
// @Param firstName query string false "First name of the employee"
// @Param department query string false "Department of the employee"
// @Param position query string false "Position of the employee"
// @Param page query int false "Page number for pagination"
// @Param limit query int false "Number of results per page"
// @Success 200 {array} models.Employee
// @Router /employees/search [get]

func SearchEmployees(c *gin.Context) {
	var employees []models.Employee

	// Récupérer les paramètres de recherche
	firstName := c.Query("firstName")
	department := c.Query("department")
	position := c.Query("position")

	// Récupérer les paramètres de pagination
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")

	// Convertir les paramètres de pagination en entiers
	pageNum, _ := strconv.Atoi(page)
	limitNum, _ := strconv.Atoi(limit)
	offset := (pageNum - 1) * limitNum

	// Construire la requête dynamique avec pagination
	query := models.DB.Offset(offset).Limit(limitNum)
	if firstName != "" {
		query = query.Where("first_name ILIKE ?", "%"+firstName+"%")
	}
	if department != "" {
		query = query.Where("department ILIKE ?", "%"+department+"%")
	}
	if position != "" {
		query = query.Where("position ILIKE ?", "%"+position+"%")
	}

	// Exécuter la requête
	if err := query.Find(&employees).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, employees)
}
