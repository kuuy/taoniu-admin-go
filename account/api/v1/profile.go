package v1

import (
  "github.com/go-chi/chi/v5"
  "net/http"
  "taoniu.local/admin/account/api"
)

type ProfileHandler struct {
  Response *api.ResponseHandler
}

func NewProfileRouter() http.Handler {
  h := ProfileHandler{}

  r := chi.NewRouter()
  r.Use(api.Authenticator)
  r.Get("/", h.Show)

  return r
}

func (h *ProfileHandler) Show(
  w http.ResponseWriter,
  r *http.Request,
) {
  h.Response = &api.ResponseHandler{
    Writer: w,
  }

  h.Response.Json(map[string]interface{}{
    "roles":        []string{"admin"},
    "introduction": "I am a super administrator",
    "avatar":       "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
    "name":         "Super Admin",
  })
}
