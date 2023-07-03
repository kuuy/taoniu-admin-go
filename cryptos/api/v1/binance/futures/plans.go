package futures

import (
  "context"
  "github.com/go-chi/chi/v5"
  "net/http"
  "strconv"

  "taoniu.local/admin/cryptos/api"
  services "taoniu.local/admin/cryptos/grpc/services/binance/futures"
  repositories "taoniu.local/admin/cryptos/repositories/binance/futures"
)

type PlanInfo struct {
  ID        string  `json:"id"`
  Symbol    string  `json:"symbol"`
  Side      int     `json:"side"`
  Price     float32 `json:"price"`
  Quantity  float32 `json:"quantity"`
  Amount    float32 `json:"amount"`
  CreatedAt int64   `json:"created_at"`
}

type PlansHandler struct {
  Ctx        context.Context
  Response   *api.ResponseHandler
  Repository *repositories.PlansRepository
}

func NewPlansRouter() http.Handler {
  h := PlansHandler{
    Ctx: context.Background(),
  }
  h.Repository = &repositories.PlansRepository{}
  h.Repository.Service = &services.Plans{
    Ctx: h.Ctx,
  }

  r := chi.NewRouter()
  r.Get("/", h.Pagenate)

  return r
}

func (h *PlansHandler) Pagenate(
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
  data := make([]*PlanInfo, len(pagenate.Data))
  for i, plan := range pagenate.Data {
    data[i] = &PlanInfo{
      ID:        plan.Id,
      Symbol:    plan.Symbol,
      Side:      int(plan.Side),
      Price:     plan.Price,
      Quantity:  plan.Quantity,
      Amount:    plan.Amount,
      CreatedAt: plan.CreatedAt.AsTime().Unix(),
    }
  }

  h.Response.Pagenate(data, pagenate.Total, current, pageSize)
}
