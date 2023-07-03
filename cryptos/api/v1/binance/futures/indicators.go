package futures

import (
  "github.com/go-chi/chi/v5"
  "net/http"
  "taoniu.local/admin/cryptos/api/v1/binance/futures/indicators"
)

func NewIndicatorsRouter() http.Handler {
  r := chi.NewRouter()
  r.Mount("/daily", indicators.NewDailyRouter())
  r.Mount("/minutely", indicators.NewMinutelyRouter())
  return r
}
