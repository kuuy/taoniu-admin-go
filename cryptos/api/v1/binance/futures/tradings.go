package futures

import (
  "github.com/go-chi/chi/v5"
  "net/http"
  "taoniu.local/admin/cryptos/api/v1/binance/futures/tradings"
)

func NewTradingsRouter() http.Handler {
  r := chi.NewRouter()
  r.Mount("/triggers", tradings.NewTriggersRouter())
  return r
}
