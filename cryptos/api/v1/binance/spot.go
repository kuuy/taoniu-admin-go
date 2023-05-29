package binance

import (
  "github.com/go-chi/chi/v5"
  "net/http"
  "taoniu.local/admin/cryptos/api/v1/binance/spot"
)

func NewSpotRouter() http.Handler {
  r := chi.NewRouter()
  r.Mount("/markets", spot.NewMarketsRouter())
  r.Mount("/analysis", spot.NewAnalysisRouter())
  r.Mount("/indicators", spot.NewIndicatorsRouter())
  r.Mount("/tradings", spot.NewTradingsRouter())
  r.Mount("/margin", spot.NewMarginRouter())
  return r
}
