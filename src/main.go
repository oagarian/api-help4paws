package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)


func main() {
	app := echo.New()
	userGroup := app.Group("/user", middleware.RateLimiterWithConfig(middleware.RateLimiterConfig(GetLimiter())))
	userGroup.GET("/get", GetAssociateds)

	adminGroup := app.Group("/admin", middleware.BasicAuth(AuthMiddleware))
	adminGroup.POST("/add", AddAssociated)
	
	app.GET("/", MainRoute)
	app.Start(":8080")
}
