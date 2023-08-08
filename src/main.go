package main

import (
	midw "modules_API/src/middlewares"
	route "modules_API/src/routes"
	_ "modules_API/src/docs"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)


func main() {
	app := echo.New()
	userGroup := app.Group("/user", middleware.RateLimiterWithConfig(middleware.RateLimiterConfig(midw.GetLimiter())))
	userGroup.GET("/get/", route.GetAssociatedsRoute)
	userGroup.GET("/get", route.GetAssociatedsRoute)

	adminGroup := app.Group("/admin", middleware.BasicAuth(midw.AuthMiddleware))
	adminGroup.POST("/add", route.AddAssociatedRoute)
	adminGroup.PUT("/update", route.UpdateAssociatedRoute)
	adminGroup.DELETE("/delete", route.DeleteAssociatedRoute)
	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	app.GET("/docs/*", echoSwagger.WrapHandler)
	app.GET("/", route.MainRoute)
	app.Start(":8080")
}
