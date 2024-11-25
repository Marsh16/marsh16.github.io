package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"archcalculator.github.io/models"
)

func ReadAllCategory(c echo.Context) error {
	result, err := models.ReadAllCategory()

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}

func DeleteCategory(c echo.Context) error {
	id := c.FormValue("id")

	result, err := models.DeleteCategory(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}

func EditCategory(c echo.Context) error {
	id := c.FormValue("id")
	name := c.FormValue("name")

	result, err := models.EditCategory(id, name)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func CreateCategory(c echo.Context) error {
	name := c.FormValue("name")
	result, err := models.CreateCategory(name)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}
