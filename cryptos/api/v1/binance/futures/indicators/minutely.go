package indicators

import (
  "github.com/go-chi/chi/v5"
  "net/http"
  "taoniu.local/admin/cryptos/api/v1/binance/futures/indicators/minutely"
)

func NewMinutelyRouter() http.Handler {
  r := chi.NewRouter()
  r.Mount("/ranking", minutely.NewRankingRouter())
  return r
}
