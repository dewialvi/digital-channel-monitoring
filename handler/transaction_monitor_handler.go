package handler

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"github.com/dewialvi/digital-channel-monitoring/models"
	"github.com/dewialvi/digital-channel-monitoring/repository"
)

type TransactionMonitorHandler struct {
	Repo     *repository.TransactionMonitorRepository
	Validate *validator.Validate
}

func NewTransactionMonitorHandler(
	repo *repository.TransactionMonitorRepository,
) *TransactionMonitorHandler {
	return &TransactionMonitorHandler{
		Repo:     repo,
		Validate: validator.New(),
	}
}

type CreateTransactionRequest struct {
	MSISDN          string  `json:"msisdn" validate:"required"`
	TransactionType string  `json:"transaction_type" validate:"required"`
	Amount          float64 `json:"amount" validate:"required,gt=0"`
	Status          string  `json:"status" validate:"required,oneof=success failed pending"`
}

func (h *TransactionMonitorHandler) Create(c echo.Context) error {
	var req CreateTransactionRequest

	// Bind request JSON ke struct
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Format request tidak valid",
		})
	}

	// Validasi request
	if err := h.Validate.Struct(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{
			"message": "Validasi gagal: " + err.Error(),
		})
	}

	// Membuat data transaction monitor
	trx := &models.TransactionMonitor{
		TransactionID:   uuid.New().String(),
		MSISDN:          req.MSISDN,
		TransactionType: req.TransactionType,
		Amount:          req.Amount,
		Status:          req.Status,
	}

	// Simpan ke database
	if err := h.Repo.Create(trx); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Gagal mencatat transaksi",
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Transaksi berhasil dicatat",
		"data":    trx,
	})
}

func (h *TransactionMonitorHandler) GetAll(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	// Default pagination
	if page < 1 {
		page = 1
	}

	if limit < 1 || limit > 100 {
		limit = 20
	}

	filter := repository.TransactionFilter{
		MSISDN: c.QueryParam("msisdn"),
		Status: c.QueryParam("status"),
		Page:   page,
		Limit:  limit,
	}

	trx, total, err := h.Repo.FindAll(filter)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Gagal mengambil data",
		})
	}

	totalPages := (total + int64(limit) - 1) / int64(limit)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": trx,
		"pagination": map[string]interface{}{
			"page":        page,
			"limit":       limit,
			"total_data":  total,
			"total_pages": totalPages,
		},
	})
}