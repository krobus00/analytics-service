package infrastructure

import (
	"net/http"
	"time"

	"github.com/krobus00/analytics-service/internal/config"
	"github.com/krobus00/analytics-service/internal/model"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
)

func makeLogEntry(c echo.Context) *logrus.Entry {
	if c == nil {
		return logrus.WithFields(logrus.Fields{
			"at": time.Now().Format("2006-01-02 15:04:05"),
		})
	}

	return logrus.WithFields(logrus.Fields{
		"at":     time.Now().Format("2006-01-02 15:04:05"),
		"method": c.Request().Method,
		"uri":    c.Request().URL.String(),
		"ip":     c.RealIP(),
	})
}

func middlewareLogging(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		makeLogEntry(c).Info("incoming request")
		return next(c)
	}
}

func NewHTTPServer() *echo.Echo {
	e := echo.New()

	e.Use(middlewareLogging)

	rateLimitConfig := middleware.RateLimiterConfig{
		Skipper: middleware.DefaultSkipper,
		Store: middleware.NewRateLimiterMemoryStoreWithConfig(
			middleware.RateLimiterMemoryStoreConfig{Rate: 10, Burst: 30, ExpiresIn: 1 * time.Hour},
		),
		IdentifierExtractor: func(ctx echo.Context) (string, error) {
			id := ctx.RealIP()
			return id, nil
		},
		ErrorHandler: func(context echo.Context, err error) error {
			res := model.NewResponse().WithMessage("forbidden").WithStatusCode(http.StatusForbidden)
			return context.JSON(res.StatusCode, res)
		},
		DenyHandler: func(context echo.Context, identifier string, err error) error {
			res := model.NewResponse().WithMessage("too many request").WithStatusCode(http.StatusTooManyRequests)
			return context.JSON(res.StatusCode, res)
		},
	}

	e.Use(middleware.RateLimiterWithConfig(rateLimitConfig))

	recoverConfig := middleware.RecoverConfig{
		StackSize: 1 << 10,
	}

	switch config.LogLevel() {
	case "error":
		recoverConfig.LogLevel = log.ERROR
	case "warn":
		recoverConfig.LogLevel = log.WARN
	default:
		recoverConfig.LogLevel = log.INFO
	}

	e.Use(middleware.RecoverWithConfig(recoverConfig))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
	}))

	return e
}
