package spot

import (
  "github.com/urfave/cli/v2"
  "taoniu.local/admin/cryptos/commands/binance/spot/margin"
)

func NewMarginCommand() *cli.Command {
  return &cli.Command{
    Name:  "margin",
    Usage: "",
    Subcommands: []*cli.Command{
      margin.NewIsolatedCommand(),
    },
  }
}
