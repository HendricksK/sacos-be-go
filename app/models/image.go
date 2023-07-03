package models 

import (
	// "log"
	"time"
	"runtime"
	// "fmt"
	// "net/http"
	// "context"
	// "github.com/labstack/echo/v4"
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
	
	rows, err := db.Query(selectQuery + " WHERE " + image_model + ".entity = '" + entity + "' AND " + image_model + ".id = " + entity_id)
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