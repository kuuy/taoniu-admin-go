package tradings

import (
  "github.com/go-chi/chi/v5"
  "net/http"
  "taoniu.local/admin/cryptos/api/v1/binance/spot/margin/isolated/tradings/fishers"
)

func NewFishersRouter() http.Handler {
  r := chi.NewRouter()
  r.Mount("/grids", fishers.NewGridsRouter())
  return r
}
