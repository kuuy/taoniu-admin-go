package commands

import (
  "fmt"
  "log"
  "net/http"
  "os"
  "taoniu.local/admin/account/api/v1"

  "github.com/go-chi/chi/v5"
  "github.com/urfave/cli/v2"
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
    r.Mount("/login", v1.NewLoginRouter())
    r.Mount("/logout", v1.NewLogoutRouter())
    r.Mount("/profile", v1.NewProfileRouter())
    r.Mount("/token", v1.NewTokenRouter())
  })

  http.ListenAndServe(
    fmt.Sprintf("127.0.0.1:%v", os.Getenv("ACCOUNT_API_PORT")),
    r,
  )

  return nil
}
