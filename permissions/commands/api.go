package commands

import (
	"log"
	"net/http"

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
	})

	http.ListenAndServe("127.0.0.1:3600", r)

	return nil
}
