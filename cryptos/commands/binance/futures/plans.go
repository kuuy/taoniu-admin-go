package futures

import (
  "context"
  "log"

  "github.com/urfave/cli/v2"

  services "taoniu.local/admin/cryptos/grpc/services/binance/futures"
  repositories "taoniu.local/admin/cryptos/repositories/binance/futures"
)

type PlansHandler struct {
  Ctx        context.Context
  Repository *repositories.PlansRepository
}

func NewPlansCommand() *cli.Command {
  var h PlansHandler
  return &cli.Command{
    Name:  "plans",
    Usage: "",
    Before: func(c *cli.Context) error {
      h = PlansHandler{
        Ctx: context.Background(),
      }
      h.Repository = &repositories.PlansRepository{}
      h.Repository.Service = &services.Plans{
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

func (h *PlansHandler) Pagenate() error {
  log.Println("binance futures tradings fishers grids pagenate...")
  symbol := ""
  r, err := h.Repository.Pagenate(symbol, 1, 15)
  if err != nil {
    return err
  }
  log.Println("result", r.Total, r.Data)

  return nil
}
