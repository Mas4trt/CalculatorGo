package handlers

import (
	calculationservice "calculator-app/internal/calculationService"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CalculationHandler struct {
	service calculationservice.CalculationService
}

func NewCalculationHandler(s calculationservice.CalculationService) *CalculationHandler {
	return &CalculationHandler{service: s}
}

func (h *CalculationHandler) GetCalculators(c echo.Context) error {
	calculations, err := h.service.GetAllCalculation()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not get calculations"})
	}

	return c.JSON(http.StatusOK, calculations)
}

func (h *CalculationHandler) PostCalculators(c echo.Context) error {
	var req calculationservice.CalculationRequest
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid request")
	}

	calc, err := h.service.CreateCalculation(req.Expression)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error":   "Could not create request",
			"details": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, calc)
}

func (h *CalculationHandler) UpdateCalculators(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "ID is required",
		})
	}

	var req calculationservice.CalculationRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error":   "Invalid request format",
			"details": err.Error(),
		})
	}

	updatedCalc, err := h.service.UpdateCalculation(id, req.Expression)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error":   "Could not update request",
			"details": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, updatedCalc)
}

func (h *CalculationHandler) DeleteCalculators(c echo.Context) error {
	id := c.Param("id")

	if err := h.service.DeleteCalculation(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not delete calculations"})
	}

	return c.NoContent(http.StatusNoContent)
}
