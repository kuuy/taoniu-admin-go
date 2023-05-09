package markets

import (
  "context"
  "log"

  "github.com/urfave/cli/v2"

  services "taoniu.local/admin/cryptos/grpc/services/binance/spot/markets"
  repositories "taoniu.local/admin/cryptos/repositories/binance/spot/markets"
)

type LiveHandler struct {
  Ctx        context.Context
  Repository *repositories.LiveRepository
}

func NewLiveCommand() *cli.Command {
  var h LiveHandler
  return &cli.Command{
    Name:  "live",
    Usage: "",
    Before: func(c *cli.Context) error {
      h = LiveHandler{
        Ctx: context.Background(),
      }
      h.Repository = &repositories.LiveRepository{}
      h.Repository.Service = &services.Live{
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

func (h *LiveHandler) Pagenate() error {
  log.Println("binance spot markets live pagenate...")
  symbol := ""
  r, err := h.Repository.Pagenate(symbol, 1, 15)
  if err != nil {
    return err
  }
  log.Println("result", r.Total, r.Data)

  return nil
}
