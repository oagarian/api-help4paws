package main

import (
	route "modules_API/src/routes"
	midw "modules_API/src/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	
	app.GET("/", route.MainRoute)
	app.Start(":8080")
}
