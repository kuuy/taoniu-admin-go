package indicators

import (
  "github.com/go-chi/chi/v5"
  "net/http"
  "taoniu.local/admin/cryptos/api/v1/binance/spot/indicators/daily"
)

func NewDailyRouter() http.Handler {
  r := chi.NewRouter()
  r.Mount("/ranking", daily.NewRankingRouter())
  return r
}
