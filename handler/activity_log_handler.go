package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/dewialvi/digital-channel-monitoring/service"
)

type ActivityLogHandler struct {
	Service *service.ActivityLogService
}

func NewActivityLogHandler(
	service *service.ActivityLogService,
) *ActivityLogHandler {
	return &ActivityLogHandler{
		Service: service,
	}
}

func (h *ActivityLogHandler) GetAll(c echo.Context) error {

	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	if page < 1 {
		page = 1
	}

	if limit < 1 || limit > 100 {
		limit = 20
	}

	logs, total, err := h.Service.GetAll(page, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Gagal mengambil activity log",
		})
	}

	totalPages := (total + int64(limit) - 1) / int64(limit)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": logs,
		"pagination": map[string]interface{}{
			"page":        page,
			"limit":       limit,
			"total_data":  total,
			"total_pages": totalPages,
		},
	})
}