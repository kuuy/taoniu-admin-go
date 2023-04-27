package commands

import (
  "context"
  "crypto/md5"
  "encoding/hex"
  "gorm.io/gorm"
  "log"

  "github.com/go-redis/redis/v8"
  "github.com/urfave/cli/v2"

  "taoniu.admin.local/account/common"
  repositories "taoniu.admin.local/account/repositories"
)

type AdminsHandler struct {
  Db         *gorm.DB
  Rdb        *redis.Client
  Ctx        context.Context
  Repository *repositories.AdminsRepository
}

func NewAdminsCommand() *cli.Command {
  var h AdminsHandler
  return &cli.Command{
    Name:  "admins",
    Usage: "",
    Before: func(c *cli.Context) error {
      h = AdminsHandler{
        Rdb: common.NewRedis(),
        Ctx: context.Background(),
      }
      h.Repository = &repositories.AdminsRepository{
        Db: common.NewDB(),
      }
      return nil
    },
    Subcommands: []*cli.Command{
      {
        Name:  "create",
        Usage: "",
        Action: func(c *cli.Context) error {
          email := c.Args().Get(0)
          password := c.Args().Get(1)
          if email == "" {
            log.Fatal("email can not be empty")
            return nil
          }
          if password == "" {
            log.Fatal("password can not be empty")
            return nil
          }
          if err := h.create(email, password); err != nil {
            return cli.Exit(err.Error(), 1)
          }
          return nil
        },
      },
    },
  }
}

func (h *AdminsHandler) create(email string, password string) error {
  log.Println("admins create...")

  hash := md5.Sum([]byte(password))
  password = hex.EncodeToString(hash[:])

  return h.Repository.Create(email, password)
}
