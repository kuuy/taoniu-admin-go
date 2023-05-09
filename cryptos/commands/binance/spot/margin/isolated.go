package margin

import (
  "github.com/urfave/cli/v2"
  "taoniu.local/admin/cryptos/commands/binance/spot/margin/isolated"
)

func NewIsolatedCommand() *cli.Command {
  return &cli.Command{
    Name:  "isolated",
    Usage: "",
    Subcommands: []*cli.Command{
      isolated.NewTradingsCommand(),
    },
  }
}
