package binance

import (
  "github.com/go-chi/chi/v5"
  "net/http"
  "taoniu.local/admin/cryptos/api/v1/binance/futures"
)

func NewFuturesRouter() http.Handler {
  r := chi.NewRouter()
  r.Mount("/markets", futures.NewMarketsRouter())
  r.Mount("/indicators", futures.NewIndicatorsRouter())
  r.Mount("/plans", futures.NewPlansRouter())
  r.Mount("/tradings", futures.NewTradingsRouter())
  return r
}
