package fishers

import (
  "context"
  "github.com/go-chi/chi/v5"
  "net/http"
  "strconv"

  "taoniu.local/admin/cryptos/api"
  services "taoniu.local/admin/cryptos/grpc/services/binance/spot/margin/isolated/tradings/fishers"
  repositories "taoniu.local/admin/cryptos/repositories/binance/spot/margin/isolated/tradings/fishers"
)

type GridInfo struct {
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

type GridsHandler struct {
  Ctx        context.Context
  Response   *api.ResponseHandler
  Repository *repositories.GridsRepository
}

func NewGridsRouter() http.Handler {
  h := GridsHandler{
    Ctx: context.Background(),
  }
  h.Repository = &repositories.GridsRepository{}
  h.Repository.Service = &services.Grids{
    Ctx: h.Ctx,
  }

  r := chi.NewRouter()
  r.Get("/", h.Pagenate)

  return r
}

func (h *GridsHandler) Pagenate(
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
  data := make([]*GridInfo, len(pagenate.Data))
  for i, grid := range pagenate.Data {
    data[i] = &GridInfo{
      ID:           grid.Id,
      Symbol:       grid.Symbol,
      BuyPrice:     grid.BuyPrice,
      BuyQuantity:  grid.BuyQuantity,
      SellPrice:    grid.SellPrice,
      SellQuantity: grid.SellQuantity,
      Status:       grid.Status,
      CreatedAt:    grid.CreatedAt.AsTime().Unix(),
      UpdatedAt:    grid.UpdatedAt.AsTime().Unix(),
    }
  }

  h.Response.Pagenate(data, pagenate.Total, current, pageSize)
}
