package handler

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	"github.com/dewialvi/digital-channel-monitoring/models"
	"github.com/dewialvi/digital-channel-monitoring/service"
)

type AuthHandler struct {
	AuthService *service.AuthService
	Validate    *validator.Validate
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{
		AuthService: authService,
		Validate:    validator.New(),
	}
}

type RegisterRequest struct {
	Name     string `json:"name" validate:"required,min=3"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Role     string `json:"role" validate:"required,oneof=admin staff"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (h *AuthHandler) Register(c echo.Context) error {
	var req RegisterRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Format request tidak valid",
		})
	}

	if err := h.Validate.Struct(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{
			"message": "Validasi gagal: " + err.Error(),
		})
	}

	user, err := h.AuthService.Register(
		req.Name,
		req.Email,
		req.Password,
		models.Role(req.Role),
	)

	if err != nil {
		return c.JSON(http.StatusConflict, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Registrasi berhasil",
		"data": map[string]interface{}{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
			"role":  user.Role,
		},
	})
}

func (h *AuthHandler) Login(c echo.Context) error {
	var req LoginRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Format request tidak valid",
		})
	}

	if err := h.Validate.Struct(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{
			"message": "Email dan password wajib diisi",
		})
	}

	token, user, err := h.AuthService.Login(
		req.Email,
		req.Password,
	)

	if err != nil {
		if err == service.ErrAccountInactive {
			return c.JSON(http.StatusForbidden, map[string]string{
				"message": err.Error(),
			})
		}

		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Login berhasil",
		"data": map[string]interface{}{
			"token": token,
			"user": map[string]interface{}{
				"id":    user.ID,
				"name":  user.Name,
				"email": user.Email,
				"role":  user.Role,
			},
		},
	})
}

func (h *AuthHandler) Logout(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Logout berhasil",
	})
}