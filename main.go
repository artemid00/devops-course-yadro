/*
  ,-.       _,---._ __  / \
 /  )    .-'       `./ /   \
(  (   ,'            `/    /|
 \  `-"             \'\   / |
  `.              ,  \ \ /  |
   /`.          ,'-`----Y   |
  (            ;        |   '
  |  ,-.    ,-'         |  /
  |  | (   |            | /
  )  |  \  `.___________|/
  `--'   `--'

  currencyAPI service
  version 1.1.0
  author a.bezpyatko

  Central Bank of Russia API

  has 2 endpoints:
  GET /info - returns service info
  GET /info/currency?date=YYYY-MM-DD&currency=CODE  - returns exchange rates
  for specified date and currency code
*/

package main

import (
	"log"
	"net/http"
	"time"

	"currencyAPI/internal/config"
	"currencyAPI/internal/handlers"
	"currencyAPI/internal/repository"
	"currencyAPI/internal/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	log.Printf("starting currencyAPI version=%s author=%s port=%s", cfg.Version, cfg.Author, cfg.Port)

	gin.DefaultWriter = log.Writer()
	gin.DefaultErrorWriter = log.Writer()
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	httpClient := &http.Client{Timeout: 10 * time.Second}
	repo := repository.NewRepository(httpClient)
	svc := usecase.NewCurrencyService(repo)
	h := handlers.New(cfg, svc)
	h.Register(r)

	if err := r.Run(":" + cfg.Port); err != nil {
		panic(err)
	}
}
