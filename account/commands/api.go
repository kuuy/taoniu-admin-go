package commands

import (
  "fmt"
  "log"
  "net/http"
  "os"

  "github.com/go-chi/chi/v5"
  "github.com/urfave/cli/v2"

  "taoniu.local/admin/account/api/v1/admin"
)

type ApiHandler struct{}

func NewApiCommand() *cli.Command {
  var h ApiHandler
  return &cli.Command{
    Name:  "api",
    Usage: "",
    Before: func(c *cli.Context) error {
      h = ApiHandler{}
      return nil
    },
    Action: func(c *cli.Context) error {
      if err := h.run(); err != nil {
        return cli.Exit(err.Error(), 1)
      }
      return nil
    },
  }
}

func (h *ApiHandler) run() error {
  log.Println("api running...")

  r := chi.NewRouter()
  r.Route("/v1", func(r chi.Router) {
    r.Mount("/login", admin.NewLoginRouter())
    r.Mount("/logout", admin.NewLogoutRouter())
    r.Mount("/token", admin.NewTokenRouter())
  })

  http.ListenAndServe(
    fmt.Sprintf("127.0.0.1:%v", os.Getenv("ACCOUNT_API_PORT")),
    r,
  )

  return nil
}
