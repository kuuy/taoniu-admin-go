package v1

import (
  "net/http"

  "github.com/go-chi/chi/v5"

  "taoniu.local/admin/cryptos/api"
  "taoniu.local/admin/cryptos/api/v1/binance"
)

func NewBinanceRouter() http.Handler {
  r := chi.NewRouter()
  r.Use(api.Authenticator)
  r.Mount("/spot", binance.NewSpotRouter())
  r.Mount("/futures", binance.NewFuturesRouter())

  return r
}
