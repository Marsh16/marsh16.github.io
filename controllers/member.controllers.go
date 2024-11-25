package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"archcalculator.github.io/models"
)

func ReadAllMember(c echo.Context) error {
	result, err := models.ReadAllMember()

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}

func DeleteMember(c echo.Context) error {
	id := c.FormValue("id")

	result, err := models.DeleteMember(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}

func EditMember(c echo.Context) error {
	id := c.FormValue("id")
	name := c.FormValue("name")
	phone_number := c.FormValue("phone_number")
	email := c.FormValue("email")
	birthday := c.FormValue("birthday")
	no_ktp := c.FormValue("no_ktp")

	result, err := models.EditMember(id, name, phone_number,email,birthday,no_ktp)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func CreateMember(c echo.Context) error {
	name := c.FormValue("name")
	phone_number := c.FormValue("phone_number")
	email := c.FormValue("email")
	birthday := c.FormValue("birthday")
	no_ktp := c.FormValue("no_ktp")
	result, err := models.CreateMember(name,phone_number,email,birthday,no_ktp)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}
