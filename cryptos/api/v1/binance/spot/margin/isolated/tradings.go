package isolated

import (
  "github.com/go-chi/chi/v5"
  "net/http"
  "taoniu.local/admin/cryptos/api/v1/binance/spot/margin/isolated/tradings"
)

func NewTradingsRouter() http.Handler {
  r := chi.NewRouter()
  r.Mount("/fishers", tradings.NewFishersRouter())
  return r
}
