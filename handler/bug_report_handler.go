package handler

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	"github.com/dewialvi/digital-channel-monitoring/models"
	"github.com/dewialvi/digital-channel-monitoring/repository"
	"github.com/dewialvi/digital-channel-monitoring/service"
)

type BugReportHandler struct {
	Service  *service.BugReportService
	Validate *validator.Validate
}

func NewBugReportHandler(s *service.BugReportService) *BugReportHandler {
	return &BugReportHandler{
		Service:  s,
		Validate: validator.New(),
	}
}

type CreateBugReportRequest struct {
	Title            string `json:"title" validate:"required,min=5"`
	Description      string `json:"description" validate:"required"`
	Severity         string `json:"severity" validate:"required,oneof=critical high medium low"`
	Priority         string `json:"priority" validate:"required,oneof=P1 P2 P3 P4"`
	StepsToReproduce string `json:"steps_to_reproduce" validate:"required"`
}

type UpdateStatusRequest struct {
	Status string `json:"status" validate:"required"`
}

func (h *BugReportHandler) Create(c echo.Context) error {
	var req CreateBugReportRequest

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

	userID := c.Get("user_id").(uint)

	bug, err := h.Service.Create(service.CreateBugReportInput{
		ReportedBy:       userID,
		Title:            req.Title,
		Description:      req.Description,
		Severity:         models.Severity(req.Severity),
		Priority:         models.Priority(req.Priority),
		StepsToReproduce: req.StepsToReproduce,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Gagal membuat bug report",
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Bug report berhasil dibuat",
		"data":    bug,
	})
}

func (h *BugReportHandler) GetByID(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "ID tidak valid",
		})
	}

	bug, err := h.Service.GetByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": bug,
	})
}

func (h *BugReportHandler) GetAll(c echo.Context) error {
	page := 1
	limit := 10

	if pageParam := c.QueryParam("page"); pageParam != "" {
		if parsedPage, err := strconv.Atoi(pageParam); err == nil && parsedPage > 0 {
			page = parsedPage
		}
	}

	if limitParam := c.QueryParam("limit"); limitParam != "" {
		if parsedLimit, err := strconv.Atoi(limitParam); err == nil && parsedLimit > 0 {
			limit = parsedLimit
		}
	}

	filter := repository.BugReportFilter{
		Search:   c.QueryParam("search"),
		Severity: c.QueryParam("severity"),
		Priority: c.QueryParam("priority"),
		Status:   c.QueryParam("status"),
		Page:     page,
		Limit:    limit,
	}

	bugs, total, err := h.Service.GetAll(filter)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Gagal mengambil data",
		})
	}

	totalPages := (total + int64(filter.Limit) - 1) / int64(filter.Limit)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": bugs,
		"pagination": map[string]interface{}{
			"page":        filter.Page,
			"limit":       filter.Limit,
			"total_data":  total,
			"total_pages": totalPages,
		},
	})
}
func (h *BugReportHandler) UpdateStatus(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "ID tidak valid",
		})
	}

	var req UpdateStatusRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Format request tidak valid",
		})
	}

	if err := h.Validate.Struct(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{
			"message": "Status wajib diisi",
		})
	}

	bug, err := h.Service.UpdateStatus(
		uint(id),
		models.BugStatus(req.Status),
	)

	if err != nil {
		if err == service.ErrBugNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{
				"message": err.Error(),
			})
		}

		return c.JSON(http.StatusUnprocessableEntity, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Status bug berhasil diupdate",
		"data":    bug,
	})
}

func (h *BugReportHandler) Delete(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "ID tidak valid",
		})
	}

	err = h.Service.Delete(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Bug report berhasil dihapus",
	})
}
