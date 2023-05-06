package spot

import (
  "github.com/go-chi/chi/v5"
  "net/http"
  "taoniu.local/admin/cryptos/api/v1/binance/spot/analysis"
)

func NewAnalysisRouter() http.Handler {
  r := chi.NewRouter()
  r.Mount("/tradings", analysis.NewTradingsRouter())
  return r
}
