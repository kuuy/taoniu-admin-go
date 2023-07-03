package tradings

import (
  "github.com/go-chi/chi/v5"
  "net/http"
  "taoniu.local/admin/cryptos/api/v1/binance/spot/margin/cross/tradings/triggers"
)

func NewTriggersRouter() http.Handler {
  r := chi.NewRouter()
  r.Mount("/grids", triggers.NewGridsRouter())
  return r
}
