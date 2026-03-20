package handlers

import (
	"errors"
	"log"
	"net/http"
	"time"

	"currencyAPI/internal/config"
	"currencyAPI/internal/usecase"

	"github.com/gin-gonic/gin"
)

// contains  handlers for the service API
type Handler struct {
	cfg config.Config
	svc *usecase.CurrencyService
}

// new constructs Handler with injected configuration and service.
func New(cfg config.Config, svc *usecase.CurrencyService) *Handler {
	return &Handler{cfg: cfg, svc: svc}
}

// register registers HTTP routes on the provided Gin engine.
func (h *Handler) Register(r *gin.Engine) {
	r.GET("/info", h.Info)
	r.GET("/info/currency", h.Currency)
}

// /info handler
func (h *Handler) Info(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"version": h.cfg.Version,
		"service": "currency",
		"author":  h.cfg.Author,
	})
}

// /info/currency handler with optional query parameters
func (h *Handler) Currency(c *gin.Context) {
	date := c.Query("date")
	currency := c.Query("currency")

	rates, err := h.svc.Rates(c.Request.Context(), date, currency)
	if err != nil {
		log.Printf("rates failed: date=%q currency=%q err=%v", date, currency, err)
		var parseErr *time.ParseError
		if errors.As(err, &parseErr) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed fetch data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"service": "currency",
		"data":    rates,
	})
}
