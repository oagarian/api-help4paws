package midw

import (
	"net/http"
	"time"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ConfigRateLimit middleware.RateLimiterConfig
func GetLimiter() ConfigRateLimit {
	return ConfigRateLimit{
		Store: middleware.NewRateLimiterMemoryStoreWithConfig(
			middleware.RateLimiterMemoryStoreConfig{Rate: 10, Burst: 20, ExpiresIn: 1 * time.Minute},
		),
		IdentifierExtractor: func(context echo.Context) (string, error) {
			ip := context.RealIP()
			return ip, nil
		},
		ErrorHandler: func(context echo.Context, err error) error {
			return context.JSON(http.StatusForbidden, nil)
		},
		DenyHandler: func(context echo.Context, identifier string,	err error) error {
        return context.JSON(http.StatusTooManyRequests, nil)
    	},
	}
}