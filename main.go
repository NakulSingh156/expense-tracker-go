package main

import (
	"expense-tracker/api"
	"log"

	"github.com/gin-gonic/gin"
	
	// NEW: Swagger imports
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "expense-tracker/docs" // This will be auto-generated in the next step
)

// @title Expense Tracker & Settlement API
// @version 1.0
// @description This is a high-performance debt settlement API using a Greedy Algorithm.
// @host localhost:8080
// @BasePath /
func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	// Your existing API routes
	router.POST("/expense", api.AddExpense)
	router.GET("/balances", api.GetBalances)
	router.GET("/settle", api.SettleDebts)

	// NEW: The Swagger UI route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Println("Server running on http://localhost:8080")
	log.Println("Swagger UI available at http://localhost:8080/swagger/index.html")
	
	router.Run(":8080")
}
