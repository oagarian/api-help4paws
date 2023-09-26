package main

import (
	_ "modules_API/src/docs"
	midw "modules_API/src/middlewares"
	route "modules_API/src/routes"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)


func main() {
	app := echo.New()
	userGroup := app.Group("/user", middleware.RateLimiterWithConfig(middleware.RateLimiterConfig(midw.GetLimiter())))
	userGroup.GET("/get/", route.GetAssociatedsRoute)
	userGroup.GET("/get", route.GetAssociatedsRoute)
	userGroup.GET("/amount", route.GetAmountRoute)
	adminGroup := app.Group("/admin", middleware.BasicAuth(midw.AuthMiddleware))
	adminGroup.POST("/add", route.AddAssociatedRoute)
	adminGroup.PUT("/update", route.UpdateAssociatedRoute)
	adminGroup.DELETE("/delete", route.DeleteAssociatedRoute)

	
	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"Content-Type", "Authorization"}, 
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))
	
	app.GET("/docs/*", echoSwagger.WrapHandler)
	app.GET("/", route.MainRoute)
	app.Start(":8080")
}
