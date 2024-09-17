package controllers

import (
	// "log"
	"net/http"
	"github.com/labstack/echo/v4"
	// "strconv"
	// "strings"
	models "github.com/HendricksK/sacosbego/app/models"
)

func GetImagesViaEntity(c echo.Context) error {
	entity := c.Param("entity")

	response := models.GetImagesViaEntity(entity)
	httpStatusCode := http.StatusOK

	if len(response) == 0 {
		httpStatusCode = http.StatusNotFound
	} 

	return c.JSON(httpStatusCode, response)
}

func GetImagesViaTag(c echo.Context) error {
	// For now, not enough data to stress about pagination
	tags := c.Param("tags")

	response := models.GetImagesViaTags(tags)
	httpStatusCode := http.StatusOK

	if len(response) == 0 {
		httpStatusCode = http.StatusNotFound

	} 

	return c.JSON(httpStatusCode, response)
}

func CreateImage(c echo.Context) error {
	httpStatusCode := models.CreateImage(c)
	return c.JSON(httpStatusCode, 0)
}

func PatchImage(c echo.Context) error {
	httpStatusCode := models.CreateImage(c)
	return c.JSON(httpStatusCode, 0)
}
