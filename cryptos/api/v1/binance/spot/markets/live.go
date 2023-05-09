package markets

import (
  "context"
  "net/http"
  "strconv"

  "github.com/go-chi/chi/v5"

  "taoniu.local/admin/cryptos/api"
  services "taoniu.local/admin/cryptos/grpc/services/binance/spot/markets"
  repositories "taoniu.local/admin/cryptos/repositories/binance/spot/markets"
)

type LiveInfo struct {
  ID        string  `json:"id"`
  Symbol    string  `json:"symbol"`
  Open      float32 `json:"open"`
  Price     float32 `json:"price"`
  High      float32 `json:"high"`
  Low       float32 `json:"low"`
  Volume    float64 `json:"volume"`
  Quota     float64 `json:"quota"`
  Timestamp int64   `json:"timestamp"`
}

type LiveHandler struct {
  Ctx        context.Context
  Response   *api.ResponseHandler
  Repository *repositories.LiveRepository
}

func NewLiveRouter() http.Handler {
  h := LiveHandler{
    Ctx: context.Background(),
  }
  h.Repository = &repositories.LiveRepository{}
  h.Repository.Service = &services.Live{
    Ctx: h.Ctx,
  }

  r := chi.NewRouter()
  r.Get("/", h.Pagenate)

  return r
}

func (h *LiveHandler) Pagenate(
  w http.ResponseWriter,
  r *http.Request,
) {
  h.Response = &api.ResponseHandler{
    Writer: w,
  }

  q := r.URL.Query()

  symbol := q.Get("symbol")

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

  pagenate, err := h.Repository.Pagenate(symbol, current, pageSize)
  if err != nil {
    h.Response.Error(http.StatusForbidden, 500, err.Error())
    return
  }
  data := make([]*LiveInfo, len(pagenate.Data))
  for i, liveInfo := range pagenate.Data {
    data[i] = &LiveInfo{
      Symbol:    liveInfo.Symbol,
      Open:      liveInfo.Open,
      Price:     liveInfo.Price,
      High:      liveInfo.High,
      Low:       liveInfo.Low,
      Volume:    liveInfo.Volume,
      Quota:     liveInfo.Quota,
      Timestamp: liveInfo.Timestamp.AsTime().Unix(),
    }
  }

  h.Response.Pagenate(data, pagenate.Total, current, pageSize)
}
