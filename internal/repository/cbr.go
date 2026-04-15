package repository

import (
	"bytes"
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"currencyAPI/internal/domain"

	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

// provides access to external rate sources
type Repository struct {
	httpClient *http.Client
	baseURL    string
}

// creates a repository that fetches daily currency rates from API
func NewRepository(httpClient *http.Client) *Repository {
	return &Repository{
		httpClient: httpClient,
		baseURL:    "http://www.cbr.ru/scripts/XML_daily.asp",
	}
}

// minimal XML model for the API response
type ValCurs struct {
	Valutes []Valute `xml:"Valute"`
}

// XML model for a single currency entry in CBR response
type Valute struct {
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

// fetches and parses rates for the given date
func (r *Repository) DailyRates(ctx context.Context, date time.Time) ([]domain.Rate, error) {
	url := r.baseURL + "?date_req=" + date.Format("02/01/2006")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0")

	resp, err := r.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error make request: %w", err)
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			fmt.Printf("failed to close response body: %v\n", err)
		}
	}()

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error read response body: %w", err)
	}

	decoded, err := decodeWindows1251(raw)
	if err != nil {
		return nil, fmt.Errorf("error decode windows-1251: %w", err)
	}

	decoded = strings.ReplaceAll(decoded, `encoding="windows-1251"`, "")

	var curs ValCurs
	if err := xml.Unmarshal([]byte(decoded), &curs); err != nil {
		return nil, err
	}

	rates := make([]domain.Rate, 0, len(curs.Valutes))
	for _, v := range curs.Valutes {
		valueStr := strings.ReplaceAll(strings.TrimSpace(v.Value), ",", ".")
		value, _ := strconv.ParseFloat(valueStr, 64)
		code := strings.TrimSpace(v.CharCode)
		if code == "" {
			continue
		}
		rates = append(rates, domain.Rate{Code: code, Value: value})
	}

	return rates, nil
}

func decodeWindows1251(data []byte) (string, error) {
	reader := transform.NewReader(bytes.NewReader(data), charmap.Windows1251.NewDecoder())
	decoded, err := io.ReadAll(reader)
	if err != nil {
		return "", fmt.Errorf("decode windows-1251: %w", err)
	}
	return string(decoded), nil
}
