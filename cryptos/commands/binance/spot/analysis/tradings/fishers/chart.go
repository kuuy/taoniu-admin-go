package fishers

import (
  "context"
  "log"
  services "taoniu.local/admin/cryptos/grpc/services/binance/spot/analysis/tradings/fishers"

  "github.com/urfave/cli/v2"

  repositories "taoniu.local/admin/cryptos/repositories/binance/spot/analysis/tradings/fishers"
)

type ChartHandler struct {
  Ctx        context.Context
  Repository *repositories.ChartRepository
}

func NewChartCommand() *cli.Command {
  var h ChartHandler
  return &cli.Command{
    Name:  "chart",
    Usage: "",
    Before: func(c *cli.Context) error {
      h = ChartHandler{
        Ctx: context.Background(),
      }
      h.Repository = &repositories.ChartRepository{}
      h.Repository.Service = &services.Chart{
        Ctx: h.Ctx,
      }
      return nil
    },
    Subcommands: []*cli.Command{
      {
        Name:  "series",
        Usage: "",
        Action: func(c *cli.Context) error {
          if err := h.Series(); err != nil {
            return cli.Exit(err.Error(), 1)
          }
          return nil
        },
      },
    },
  }
}

func (h *ChartHandler) Series() error {
  log.Println("binance spot analysis tradings fishers chart series...")

  series, err := h.Repository.Series(10)
  if err != nil {
    return err
  }
  log.Println("series", series)

  return nil
}
