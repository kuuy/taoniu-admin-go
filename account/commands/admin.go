package commands

import (
	"github.com/urfave/cli/v2"
	"taoniu.admin.local/account/commands/admin"
)

func NewAdminCommand() *cli.Command {
	return &cli.Command{
		Name:  "admin",
		Usage: "",
		Subcommands: []*cli.Command{
			admin.NewUsersCommand(),
		},
	}
}
