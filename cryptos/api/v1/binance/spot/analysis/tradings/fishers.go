package tradings

import (
  "github.com/go-chi/chi/v5"
  "net/http"
  "taoniu.local/admin/cryptos/api/v1/binance/spot/analysis/tradings/fishers"
)

func NewFishersRouter() http.Handler {
  r := chi.NewRouter()
  r.Mount("/chart", fishers.NewChartRouter())
  return r
}
