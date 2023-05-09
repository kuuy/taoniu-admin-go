package spot

import (
  "github.com/urfave/cli/v2"
  "taoniu.local/admin/cryptos/commands/binance/spot/markets"
)

func NewMarketsCommand() *cli.Command {
  return &cli.Command{
    Name:  "markets",
    Usage: "",
    Subcommands: []*cli.Command{
      markets.NewLiveCommand(),
    },
  }
}
