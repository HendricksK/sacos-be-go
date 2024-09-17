package models 

import (
	// "log"
	"time"
	"runtime"
	"fmt"
	"strings"
	"net/http"
	// "context"
	"github.com/labstack/echo/v4"
	// // "encoding/json"
	database "github.com/HendricksK/sacosbego/app/database"
	extensions "github.com/HendricksK/sacosbego/app/extensions"

)

type Image struct {
	Id 			*int 		`json:"id"`
	Entity      *string		`json:"entity"` 
	EntityId	*int 		`json:"entity_id"`
	Url 		*string 	`json:"url"`
	Tags		*string		`json:"tags"`		
}

type ImageAggregate struct {
	ImageId			*int 		`json:"image_id"`
	Tags 			*string 	`json:"tags"` 
	CreatedAt 		*time.Time  `json:"created_at"`
	UpdatedAt 		*time.Time  `json:"updated_at"`
}

var image_model = "image"
var image_model_aggregate = "image_aggregate"

func GetImages(entity string, entity_id string) []Image {
	var images []Image

	db := database.Open()

	// CREATE CONN
	fields := []string{
		image_model + ".id",
		image_model + ".entity",
		image_model + ".entity_id",
		image_model + ".url",
		image_model_aggregate + ".tags"}

	var selectQuery = BuildSelectQueryWithAggregate(fields, image_model, image_model_aggregate)
	
	rows, err := db.Query(selectQuery + ` WHERE ` + image_model + `.entity = '` + entity + `' AND ` + image_model + `.id = ` + entity_id)
	if err != nil {
		_, filename, line, _ := runtime.Caller(1)
		extensions.Log(err.Error(), filename, line)
		return images
	}
	defer rows.Close()

	for rows.Next() {
		var image Image
		
		err = rows.Scan(
			&image.Id, 
			&image.Entity, 
			&image.EntityId, 
			&image.Url,  
			&image.Tags)

		if err != nil {
			_, filename, line, _ := runtime.Caller(1)
			extensions.Log(err.Error(), filename, line)
			panic(err)
		}

		images = append(images, image)
	}

	database.Close(db)

	// CLOSE CONN
	return images
}

func GetImagesViaTags(tags string) []Image {
	var images []Image

	db := database.Open()

	// CREATE CONN
	fields := []string{
		image_model + ".id",
		image_model + ".entity",
		image_model + ".entity_id",
		image_model + ".url",
		image_model_aggregate + ".tags"}

	var selectQuery = BuildSelectQueryWithAggregate(fields, image_model, image_model_aggregate)

	// SELECT * FROM image WHERE image.id IN(
		// SELECT image_id FROM image_aggregate ia2 WHERE JSON_CONTAINS(JSON_EXTRACT(tags, '$.tags'), '["williams"]'))
	tagSlice := strings.Split(tags, ",")

	for _, tag := range tagSlice {

		rows, err := db.Query(selectQuery + ` WHERE ` + image_model + `.id IN(SELECT image_id FROM image_aggregate WHERE JSON_CONTAINS(JSON_EXTRACT(tags, '$.tags'), '["` + tag + `"]'))`)
		if err != nil {
			_, filename, line, _ := runtime.Caller(1)
			extensions.Log(err.Error(), filename, line)
			return images
		}
		defer rows.Close()

		for rows.Next() {
			var image Image
			
			err = rows.Scan(
				&image.Id, 
				&image.Entity, 
				&image.EntityId, 
				&image.Url,  
				&image.Tags)

			if err != nil {
				_, filename, line, _ := runtime.Caller(1)
				extensions.Log(err.Error(), filename, line)
				panic(err)
			}

			images = append(images, image)
		}
	}

	database.Close(db)

	// CLOSE CONN
	return images
}

func GetImagesViaEntity(entity string) []Image {
	var images []Image

	db := database.Open()

	// CREATE CONN
	fields := []string{
		image_model + ".id",
		image_model + ".entity",
		image_model + ".entity_id",
		image_model + ".url",
		image_model_aggregate + ".tags"}

	var selectQuery = BuildSelectQueryWithAggregate(fields, image_model, image_model_aggregate)

	fmt.Println(selectQuery)
	
	rows, err := db.Query(selectQuery + ` WHERE ` + image_model + `.entity = '` + entity + `'`)
	if err != nil {
		_, filename, line, _ := runtime.Caller(1)
		extensions.Log(err.Error(), filename, line)
		return images
	}
	defer rows.Close()

	for rows.Next() {
		var image Image
		
		err = rows.Scan(
			&image.Id, 
			&image.Entity, 
			&image.EntityId, 
			&image.Url,  
			&image.Tags)

		if err != nil {
			_, filename, line, _ := runtime.Caller(1)
			extensions.Log(err.Error(), filename, line)
			panic(err)
		}

		images = append(images, image)
	}

	database.Close(db)

	// CLOSE CONN
	return images
}

func CreateImage(c echo.Context) int {
	return http.StatusForbidden
}

func UpdateImage(c echo.Context) int {
	return http.StatusForbidden
}

func DeleteImage(c echo.Context) int {
	return http.StatusForbidden
}