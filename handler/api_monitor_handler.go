package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/dewialvi/digital-channel-monitoring/repository"
)

type APIMonitorHandler struct {
	Repo *repository.APIMonitorRepository
}

func NewAPIMonitorHandler(
	repo *repository.APIMonitorRepository,
) *APIMonitorHandler {
	return &APIMonitorHandler{
		Repo: repo,
	}
}

func (h *APIMonitorHandler) GetAll(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	statusCode, _ := strconv.Atoi(c.QueryParam("status_code"))

	if page < 1 {
		page = 1
	}

	if limit < 1 || limit > 100 {
		limit = 20
	}

	filter := repository.APIMonitorFilter{
		Endpoint:   c.QueryParam("endpoint"),
		StatusCode: statusCode,
		Page:       page,
		Limit:      limit,
	}

	logs, total, err := h.Repo.FindAll(filter)

	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{
				"message": "Gagal mengambil data monitoring",
			},
		)
	}

	totalPages := (total + int64(limit) - 1) / int64(limit)

	return c.JSON(
		http.StatusOK,
		map[string]interface{}{
			"data": logs,
			"pagination": map[string]interface{}{
				"page":        page,
				"limit":       limit,
				"total_data":  total,
				"total_pages": totalPages,
			},
		},
	)
}

func (h *APIMonitorHandler) GetStats(c echo.Context) error {
	hoursParam := c.QueryParam("hours")

	hours, err := strconv.Atoi(hoursParam)

	if err != nil || hours <= 0 {
		hours = 24
	}

	since := time.Now().Add(
		-time.Duration(hours) * time.Hour,
	)

	stats, err := h.Repo.GetStatsSummary(since)

	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{
				"message": "Gagal mengambil statistik",
			},
		)
	}

	return c.JSON(
		http.StatusOK,
		map[string]interface{}{
			"period_hours": hours,
			"data":         stats,
		},
	)
}