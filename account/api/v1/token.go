package v1

import (
  "encoding/json"
  "github.com/go-chi/chi/v5"
  "net/http"
  "taoniu.local/admin/account/api"
  repositories "taoniu.local/admin/account/repositories/admin"
)

type TokenHandler struct {
  Response        *api.ResponseHandler
  TokenRepository *repositories.TokenRepository
}

type AccessToken struct {
  AccessToken string `json:"access_token"`
}

type RefreshParams struct {
  RefreshToken string `json:"refresh_token"`
}

func NewTokenRouter() http.Handler {
  h := TokenHandler{}

  r := chi.NewRouter()
  r.Post("/refresh", h.Refresh)

  return r
}

func (h *TokenHandler) Token() *repositories.TokenRepository {
  if h.TokenRepository == nil {
    h.TokenRepository = &repositories.TokenRepository{}
  }
  return h.TokenRepository
}

func (h *TokenHandler) Refresh(
  w http.ResponseWriter,
  r *http.Request,
) {
  h.Response = &api.ResponseHandler{
    Writer: w,
  }

  var params RefreshParams
  err := json.NewDecoder(r.Body).Decode(&params)
  if err != nil {
    h.Response.Error(http.StatusForbidden, 1004, "request not valid")
    return
  }

  uid, err := h.Token().Uid(params.RefreshToken)
  if err != nil {
    if uid != "" {
      h.Response.Error(http.StatusForbidden, 401, err.Error())
    } else {
      h.Response.Error(http.StatusForbidden, 403, "token not valid")
    }
    return
  }

  accessToken, err := h.Token().AccessToken(uid)
  if err != nil {
    h.Response.Error(http.StatusInternalServerError, 500, "server error")
    return
  }

  token := &AccessToken{
    AccessToken: accessToken,
  }

  h.Response.Json(token)
}
