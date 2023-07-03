package minutely

import (
  "context"
  "github.com/go-chi/chi/v5"
  "net/http"
  "strconv"
  "strings"

  "taoniu.local/admin/cryptos/api"
  services "taoniu.local/admin/cryptos/grpc/services/binance/futures/indicators/minutely"
  repositories "taoniu.local/admin/cryptos/repositories/binance/futures/indicators/minutely"
)

type RankingHandler struct {
  Ctx        context.Context
  Response   *api.ResponseHandler
  Repository *repositories.RankingRepository
}

func NewRankingRouter() http.Handler {
  h := RankingHandler{
    Ctx: context.Background(),
  }
  h.Repository = &repositories.RankingRepository{}
  h.Repository.Service = &services.Ranking{
    Ctx: h.Ctx,
  }

  r := chi.NewRouter()
  r.Get("/", h.Pagenate)

  return r
}

func (h *RankingHandler) Pagenate(
  w http.ResponseWriter,
  r *http.Request,
) {
  h.Response = &api.ResponseHandler{
    Writer: w,
  }

  q := r.URL.Query()

  symbol := q.Get("symbol")

  if q.Get("fields") == "" {
    h.Response.Error(http.StatusForbidden, 1004, "fields is empty")
    return
  }

  if q.Get("sort") == "" {
    h.Response.Error(http.StatusForbidden, 1004, "sort is empty")
    return
  }

  fields := strings.Split(q.Get("fields"), ",")

  sort := strings.Split(q.Get("sort"), ",")
  sortField := sort[0]
  sortType, _ := strconv.Atoi(sort[1])

  var current int
  if !q.Has("current") {
    current = 1
  }
  current, _ = strconv.Atoi(q.Get("current"))
  if current < 1 {
    h.Response.Error(http.StatusForbidden, 1004, "current not valid")
    return
  }

  var pageSize int
  if !q.Has("page_size") {
    pageSize = 50
  } else {
    pageSize, _ = strconv.Atoi(q.Get("page_size"))
  }
  if pageSize < 1 || pageSize > 100 {
    h.Response.Error(http.StatusForbidden, 1004, "page size not valid")
    return
  }

  pagenate, err := h.Repository.Pagenate(symbol, fields, sortField, int32(sortType), current, pageSize)
  if err != nil {
    h.Response.Error(http.StatusForbidden, 500, err.Error())
    return
  }

  h.Response.Pagenate(pagenate.Data, pagenate.Total, current, pageSize)
}
