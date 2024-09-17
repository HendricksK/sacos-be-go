package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"

	controllers "github.com/HendricksK/sacosbego/app/controllers"
	// auth "github.com/HendricksK/sacosbego/app/auth"
)

func Routes() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())


	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://sacos.localhost","http://127.0.0.1:8080", "102.182.190.217"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodOptions},
	}))
	

	// HEALTH
	e.GET("/ping", controllers.Ping)
	e.GET("/healthz", controllers.Healthz)

	// ARTICLES
	e.GET("/article/:id", controllers.GetArticle)
	e.GET("/articles", controllers.GetArticles)
	e.POST("/article", controllers.CreateArticle)
	e.PATCH("/article", controllers.PatchArticle)
	e.DELETE("/article", controllers.DeleteArticle)

	// RIDER
	e.GET("/rider/:id", controllers.GetRider)
	e.GET("/riders", controllers.GetRiders)
	e.POST("/rider", controllers.CreateRider)
	e.PATCH("/rider", controllers.PatchRider)
	e.DELETE("/rider", controllers.DeleteRider)

	// IMAGE 
	e.GET("/image/tags/:tags", controllers.GetImagesViaTag)
	e.GET("/image/entity/:entity", controllers.GetImagesViaEntity)
	e.POST("/image", controllers.CreateImage)

	// PAGE
	e.GET("page/:id", controllers.GetPage)

	// need to get port from os here.
	e.Logger.Fatal(e.Start("localhost:9002"))
}
