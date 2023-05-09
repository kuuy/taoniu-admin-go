package spot

import (
  "github.com/go-chi/chi/v5"
  "net/http"
  "taoniu.local/admin/cryptos/api/v1/binance/spot/markets"
)

func NewMarketsRouter() http.Handler {
  r := chi.NewRouter()
  r.Mount("/live", markets.NewLiveRouter())
  return r
}
