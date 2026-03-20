package usecase

import (
	"context"
	"strings"
	"time"

	"currencyAPI/internal/domain"
)

// describes a daily currency rates interface
type RateRepository interface {
	DailyRates(ctx context.Context, date time.Time) ([]domain.Rate, error)
}

// contains application logic for retrieving currency rates
type CurrencyService struct {
	repo RateRepository
}

// constructs CurrencyService with the provided repository
func NewCurrencyService(repo RateRepository) *CurrencyService {
	return &CurrencyService{
		repo: repo,
	}
}

// rates returns daily rates optionally filtered by currency code
func (s *CurrencyService) Rates(ctx context.Context, dateParam string, currency string) (map[string]float64, error) {
	date := time.Now()
	if dateParam != "" {
		parsed, err := time.Parse("2006-01-02", dateParam)
		if err != nil {
			return nil, err
		}
		date = parsed
	}

	rates, err := s.repo.DailyRates(ctx, date)
	if err != nil {
		return nil, err
	}

	currency = strings.ToUpper(strings.TrimSpace(currency))
	result := make(map[string]float64, len(rates))
	for _, r := range rates {
		if currency != "" && r.Code != currency {
			continue
		}
		result[r.Code] = r.Value
	}
	return result, nil
}
