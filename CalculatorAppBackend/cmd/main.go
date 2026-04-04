package main

import (
	calculationservice "calculator-app/internal/calculationService"
	"calculator-app/internal/db"
	"calculator-app/internal/handlers"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database, err := db.InitDB()

	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	e := echo.New()

	calcRepo := calculationservice.NewCalculationRepository(database)
	calcService := calculationservice.NewCalculationService(calcRepo)
	calcHandlers := handlers.NewCalculationHandler(calcService)

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/calculations", calcHandlers.GetCalculators)
	e.POST("/calculations", calcHandlers.PostCalculators)
	e.PATCH("/calculations/:id", calcHandlers.UpdateCalculators)
	e.DELETE("/calculations/:id", calcHandlers.DeleteCalculators)

	e.Start("localhost:8080")
}
