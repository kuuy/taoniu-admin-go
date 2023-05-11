package tradings

import (
  "context"
  "github.com/go-chi/chi/v5"
  "net/http"
  "strconv"

  "taoniu.local/admin/cryptos/api"
  services "taoniu.local/admin/cryptos/grpc/services/binance/spot/tradings"
  repositories "taoniu.local/admin/cryptos/repositories/binance/spot/tradings"
)

type TriggersInfo struct {
  ID           string  `json:"id"`
  Symbol       string  `json:"symbol"`
  BuyPrice     float32 `json:"buy_price"`
  BuyQuantity  float32 `json:"buy_quantity"`
  SellPrice    float32 `json:"sell_price"`
  SellQuantity float32 `json:"sell_quantity"`
  Status       int32   `json:"status"`
  CreatedAt    int64   `json:"created_at"`
  UpdatedAt    int64   `json:"updated_at"`
}

type TriggersHandler struct {
  Ctx        context.Context
  Response   *api.ResponseHandler
  Repository *repositories.TriggersRepository
}

func NewTriggersRouter() http.Handler {
  h := TriggersHandler{
    Ctx: context.Background(),
  }
  h.Repository = &repositories.TriggersRepository{}
  h.Repository.Service = &services.Triggers{
    Ctx: h.Ctx,
  }

  r := chi.NewRouter()
  r.Get("/", h.Pagenate)

  return r
}

func (h *TriggersHandler) Pagenate(
  w http.ResponseWriter,
  r *http.Request,
) {
  h.Response = &api.ResponseHandler{
    Writer: w,
  }

  q := r.URL.Query()

  symbol := q.Get("symbol")

  var status []uint32
  if q.Has("status[]") {
    for _, value := range q["status[]"] {
      number, _ := strconv.Atoi(value)
      status = append(status, uint32(number))
    }
  }

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

  pagenate, err := h.Repository.Pagenate(symbol, status, current, pageSize)
  if err != nil {
    h.Response.Error(http.StatusForbidden, 500, err.Error())
    return
  }
  data := make([]*TriggersInfo, len(pagenate.Data))
  for i, trigger := range pagenate.Data {
    data[i] = &TriggersInfo{
      ID:           trigger.Id,
      Symbol:       trigger.Symbol,
      BuyPrice:     trigger.BuyPrice,
      BuyQuantity:  trigger.BuyQuantity,
      SellPrice:    trigger.SellPrice,
      SellQuantity: trigger.SellQuantity,
      Status:       trigger.Status,
      CreatedAt:    trigger.CreatedAt.AsTime().Unix(),
      UpdatedAt:    trigger.UpdatedAt.AsTime().Unix(),
    }
  }

  h.Response.Pagenate(data, pagenate.Total, current, pageSize)
}
