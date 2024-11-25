package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"archcalculator.github.io/models"
)

func ReadBookByCategoryId(c echo.Context) error {
	id := c.FormValue("id")
	result, err := models.ReadBookByCategoryId(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}

func ReadCategoryByBookId(c echo.Context) error {
	id := c.FormValue("id")
	result, err := models.ReadCategoryByBookId(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}

func DeleteBookCategory(c echo.Context) error {
	id := c.FormValue("id")

	result, err := models.DeleteCategory(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}

func CreateBookCategory(c echo.Context) error {
	book_id := c.FormValue("book_id")
	category_id := c.FormValue("category_id")
	result, err := models.CreateBookCategory(book_id,category_id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}
