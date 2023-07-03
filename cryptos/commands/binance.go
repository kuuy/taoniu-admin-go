package commands

import (
  "github.com/urfave/cli/v2"
  "taoniu.local/admin/cryptos/commands/binance"
)

func NewBinanceCommand() *cli.Command {
  return &cli.Command{
    Name:  "binance",
    Usage: "",
    Subcommands: []*cli.Command{
      binance.NewSpotCommand(),
      binance.NewFuturesCommand(),
    },
  }
}
