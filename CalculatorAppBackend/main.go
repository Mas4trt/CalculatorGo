package main

import (
	"fmt"
	"net/http"

	"github.com/Knetic/govaluate"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Calculation struct {
	ID         string `json:"id"`
	Expression string `json:"expression"`
	Result     string `json:"result"`
}

type CalculationRequest struct {
	Expression string `json:"expression"`
}

var calculations = []Calculation{}

func calculaExpression(expression string) (string, error) {
	expr, err := govaluate.NewEvaluableExpression(expression)
	if err != nil {
		return "", err
	}
	result, err := expr.Evaluate(nil)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", result), err
}

func getCalculators(c echo.Context) error {
	return c.JSON(http.StatusOK, calculations)
}

func postCalculators(c echo.Context) error {
	var req CalculationRequest
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid request")
	}

	result, err := calculaExpression(req.Expression)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid request")
	}

	calc := Calculation{
		ID:         uuid.NewString(),
		Expression: req.Expression,
		Result:     result,
	}
	calculations = append(calculations, calc)

	return c.JSON(http.StatusCreated, calc)

}

func main() {
	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/calculations", getCalculators)
	e.POST("/calculations", postCalculators)

	e.Start("localhost:8080")
}
