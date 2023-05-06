package fishers

import (
  "context"
  "log"
  services "taoniu.local/admin/cryptos/grpc/services/binance/spot/tradings/fishers"
  repositories "taoniu.local/admin/cryptos/repositories/binance/spot/tradings/fishers"

  "github.com/urfave/cli/v2"
)

type GridsHandler struct {
  Ctx        context.Context
  Repository *repositories.GridsRepository
}

func NewGridsCommand() *cli.Command {
  var h GridsHandler
  return &cli.Command{
    Name:  "grids",
    Usage: "",
    Before: func(c *cli.Context) error {
      h = GridsHandler{
        Ctx: context.Background(),
      }
      h.Repository = &repositories.GridsRepository{}
      h.Repository.Service = &services.Grids{
        Ctx: h.Ctx,
      }
      return nil
    },
    Subcommands: []*cli.Command{
      {
        Name:  "pagenate",
        Usage: "",
        Action: func(c *cli.Context) error {
          if err := h.Pagenate(); err != nil {
            return cli.Exit(err.Error(), 1)
          }
          return nil
        },
      },
    },
  }
}

func (h *GridsHandler) Pagenate() error {
  log.Println("binance spot analysis tradings fishers grids series...")

  r, err := h.Repository.Pagenate(1, 15)
  if err != nil {
    return err
  }
  log.Println("result", r.Total, r.Data)

  return nil
}
