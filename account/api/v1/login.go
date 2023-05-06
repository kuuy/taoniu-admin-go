package v1

import (
  "encoding/json"
  "net/http"

  "github.com/go-chi/chi/v5"
  "gorm.io/gorm"

  "taoniu.local/admin/account/api"
  "taoniu.local/admin/account/common"
  repositories "taoniu.local/admin/account/repositories/admin"
)

type LoginHandler struct {
  Db              *gorm.DB
  Response        *api.ResponseHandler
  UserRepository  *repositories.UsersRepository
  TokenRepository *repositories.TokenRepository
}

type Token struct {
  AccessToken  string `json:"access_token"`
  RefreshToken string `json:"refresh_token"`
}

type LoginParams struct {
  Email    string `json:"email"`
  Password string `json:"password"`
}

func NewLoginRouter() http.Handler {
  h := LoginHandler{
    Db: common.NewDB(),
  }

  r := chi.NewRouter()
  r.Post("/", h.Do)

  return r
}

func (h *LoginHandler) Users() *repositories.UsersRepository {
  if h.UserRepository == nil {
    h.UserRepository = &repositories.UsersRepository{
      Db: h.Db,
    }
  }
  return h.UserRepository
}

func (h *LoginHandler) Token() *repositories.TokenRepository {
  if h.TokenRepository == nil {
    h.TokenRepository = &repositories.TokenRepository{}
  }
  return h.TokenRepository
}

func (h *LoginHandler) Do(
  w http.ResponseWriter,
  r *http.Request,
) {
  h.Response = &api.ResponseHandler{
    Writer: w,
  }

  var params LoginParams
  err := json.NewDecoder(r.Body).Decode(&params)
  if err != nil {
    h.Response.Error(http.StatusForbidden, 1004, "request not valid")
    return
  }

  user := h.Users().Get(params.Email)
  if user == nil {
    h.Response.Error(http.StatusForbidden, 1000, "email or password not exists")
    return
  }
  if !common.VerifyPassword(params.Password, user.Salt, user.Password) {
    h.Response.Error(http.StatusForbidden, 1000, "email or password not exists")
    return
  }

  accessToken, err := h.Token().AccessToken(user.ID)
  if err != nil {
    h.Response.Error(http.StatusInternalServerError, 500, "server error")
    return
  }
  refreshToken, err := h.Token().RefreshToken(user.ID)
  if err != nil {
    h.Response.Error(http.StatusInternalServerError, 500, "server error")
    return
  }

  token := &Token{
    AccessToken:  accessToken,
    RefreshToken: refreshToken,
  }

  h.Response.Json(token)
}
