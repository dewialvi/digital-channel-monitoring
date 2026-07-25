package middleware

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/dewialvi/digital-channel-monitoring/models"
	"github.com/dewialvi/digital-channel-monitoring/repository"
)

func APIMonitorMiddleware(
	repo *repository.APIMonitorRepository,
) echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {

		return func(c echo.Context) error {

			start := time.Now()

			// Jalankan handler/request asli
			err := next(c)

			// Hitung response time
			duration := time.Since(start).Milliseconds()

			// Simpan log monitoring secara asynchronous
			go func() {

				log := &models.APIMonitor{
					Endpoint:       c.Path(),
					Method:         c.Request().Method,
					StatusCode:     c.Response().Status,
					ResponseTimeMs: int(duration),
					CheckedAt:      start,
				}

				if err := repo.Create(log); err != nil {
					fmt.Println(
						"Failed to log API monitoring:",
						err,
					)
				}

			}()

			return err
		}
	}
}