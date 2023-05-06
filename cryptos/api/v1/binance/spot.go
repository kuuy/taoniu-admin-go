package binance

import (
  "github.com/go-chi/chi/v5"
  "net/http"
  "taoniu.local/admin/cryptos/api/v1/binance/spot"
)

func NewSpotRouter() http.Handler {
  r := chi.NewRouter()
  r.Mount("/analysis", spot.NewAnalysisRouter())
  r.Mount("/tradings", spot.NewTradingsRouter())
  return r
}
