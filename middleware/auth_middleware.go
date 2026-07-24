package middleware

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"

	"github.com/dewialvi/digital-channel-monitoring/utils"
)

func JWTMiddleware(secret string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			authHeader := c.Request().Header.Get("Authorization")

			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"message": "Token tidak ditemukan, silakan login",
				})
			}

			parts := strings.Split(authHeader, " ")

			if len(parts) != 2 || parts[0] != "Bearer" {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"message": "Format token tidak valid",
				})
			}

			claims, err := utils.ParseJWT(parts[1], secret)

			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"message": "Token tidak valid atau sudah kedaluwarsa",
				})
			}

			c.Set("user_id", claims.UserID)
			c.Set("email", claims.Email)
			c.Set("role", claims.Role)

			return next(c)
		}
	}
}

func RoleMiddleware(allowedRoles ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			role, ok := c.Get("role").(string)

			if !ok {
				return c.JSON(http.StatusForbidden, map[string]string{
					"message": "Akses ditolak",
				})
			}

			for _, allowedRole := range allowedRoles {
				if role == allowedRole {
					return next(c)
				}
			}

			return c.JSON(http.StatusForbidden, map[string]string{
				"message": "Anda tidak memiliki akses untuk aksi ini",
			})
		}
	}
}