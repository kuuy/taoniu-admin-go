package fishers

import (
  "context"
  "github.com/go-chi/chi/v5"
  "net/http"
  "strconv"

  "taoniu.local/admin/cryptos/api"
  services "taoniu.local/admin/cryptos/grpc/services/binance/spot/analysis/tradings/fishers"
  repositories "taoniu.local/admin/cryptos/repositories/binance/spot/analysis/tradings/fishers"
)

type ChartHandler struct {
  Ctx        context.Context
  Response   *api.ResponseHandler
  Repository *repositories.ChartRepository
}

func NewChartRouter() http.Handler {
  h := ChartHandler{
    Ctx: context.Background(),
  }
  h.Repository = &repositories.ChartRepository{}
  h.Repository.Service = &services.Chart{
    Ctx: h.Ctx,
  }

  r := chi.NewRouter()
  r.Get("/series", h.Series)

  return r
}

func (h *ChartHandler) Series(
  w http.ResponseWriter,
  r *http.Request,
) {
  h.Response = &api.ResponseHandler{
    Writer: w,
  }

  var limit int
  if !r.URL.Query().Has("limit") {
    limit = 15
  } else {
    limit, _ = strconv.Atoi(r.URL.Query().Get("limit"))
  }
  if limit < 1 || limit > 100 {
    h.Response.Error(http.StatusForbidden, 1004, "limit not valid")
    return
  }

  series, err := h.Repository.Series(limit)
  if err != nil {
    h.Response.Error(http.StatusForbidden, 500, err.Error())
    return
  }

  h.Response.Json(series)
}
