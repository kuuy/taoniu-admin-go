package v1

import (
  "net/http"

  "github.com/go-chi/chi/v5"
  "gorm.io/gorm"

  "taoniu.local/admin/permissions/api"
  "taoniu.local/admin/permissions/common"
)

type RolesHandler struct {
  Db       *gorm.DB
  Response *api.ResponseHandler
}

func NewRolesRouter() http.Handler {
  h := RolesHandler{
    Db: common.NewDB(),
  }

  r := chi.NewRouter()
  r.Post("/", h.Do)

  return r
}

func (h *RolesHandler) Do(
  w http.ResponseWriter,
  r *http.Request,
) {
  h.Response = &api.ResponseHandler{
    Writer: w,
  }
}
