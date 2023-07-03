package binance

import (
  "github.com/urfave/cli/v2"
  "taoniu.local/admin/cryptos/commands/binance/futures"
)

func NewFuturesCommand() *cli.Command {
  return &cli.Command{
    Name:  "futures",
    Usage: "",
    Subcommands: []*cli.Command{
      futures.NewPlansCommand(),
    },
  }
}
