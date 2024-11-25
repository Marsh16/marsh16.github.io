package controllers

import (
	"io"
	// "archcalculator.github.io/helpers"
	"archcalculator.github.io/models"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func ReadBookByBookId(c echo.Context) error {
	id := c.FormValue("id")

	result, err := models.ReadBookByBookId(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}

func ReadAllBook(c echo.Context) error {
	result, err := models.ReadAllBook()

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}

func ReadBookByMemberId(c echo.Context) error {
	id := c.FormValue("id")
	result, err := models.ReadBookByMemberId(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}

func DeleteBook(c echo.Context) error {
	id := c.FormValue("id")

	result, err := models.DeleteBook(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}

func EditBook(c echo.Context) error {
	id := c.FormValue("id")
	title := c.FormValue("title")
	synopsis := c.FormValue("synopsis")
	cover_image, err := c.FormFile("cover_image")
	author := c.FormValue("author")
	publish_date := c.FormValue("publish_date")
	member_id := c.FormValue("member_id")

	if err != nil {
		return c.JSON(http.StatusBadRequest, &models.Response{
			Message: "Invalid data! The data type must be images!",
		})
	}

	pathImage := "./images/" + cover_image.Filename

	if err := saveUploadedFile(cover_image, pathImage); err != nil {
		return c.JSON(http.StatusInternalServerError, &models.Response{
			Message: "An internal server error occurred when saving the image. Please try again in a few moments!",
		})
	}

	baseURL := "http://marsh16.github.io/"
	pictureURL := baseURL + "/images/" + cover_image.Filename

	result, err := models.EditBook(id, title, synopsis, pictureURL, author, publish_date, member_id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func CreateBook(c echo.Context) error {
	title := c.FormValue("title")
	synopsis := c.FormValue("synopsis")
	cover_image, err := c.FormFile("cover_image")
	author := c.FormValue("author")
	publish_date := c.FormValue("publish_date")
	member_id := c.FormValue("member_id")

	if err != nil {
		return c.JSON(http.StatusBadRequest, &models.Response{
			Message: "Invalid data! The data type must be images!",
		})
	}

	pathImage := "./images/" + cover_image.Filename

	if err := saveUploadedFile(cover_image, pathImage); err != nil {
		return c.JSON(http.StatusInternalServerError, &models.Response{
			Message: "An internal server error occurred when saving the image. Please try again in a few moments!",
		})
	}
	baseURL := "http://marsh16.github.io/"
	pictureURL := baseURL + "/images/" + cover_image.Filename

	result, err := models.CreateBook(title, synopsis, pictureURL, author, publish_date, member_id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}

func saveUploadedFile(file *multipart.FileHeader, path string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(path)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return err
	}
	return nil
}
