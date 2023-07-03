package margin

import (
  "github.com/go-chi/chi/v5"
  "net/http"
  "taoniu.local/admin/cryptos/api/v1/binance/spot/margin/cross"
)

func NewCrossRouter() http.Handler {
  r := chi.NewRouter()
  r.Mount("/tradings", cross.NewTradingsRouter())
  return r
}
