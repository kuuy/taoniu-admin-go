package binance

import (
  "github.com/urfave/cli/v2"
  "taoniu.local/admin/cryptos/commands/binance/spot"
)

func NewSpotCommand() *cli.Command {
  return &cli.Command{
    Name:  "spot",
    Usage: "",
    Subcommands: []*cli.Command{
      spot.NewMarketsCommand(),
      spot.NewAnalysisCommand(),
      spot.NewTradingsCommand(),
      spot.NewMarginCommand(),
    },
  }
}
