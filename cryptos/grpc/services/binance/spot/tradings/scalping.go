package tradings

import (
  "context"
  "fmt"
  "os"

  "google.golang.org/grpc"
  "google.golang.org/grpc/credentials/insecure"

  pb "taoniu.local/admin/cryptos/grpc/binance/spot/tradings/scalping"
)

type Scalping struct {
  Ctx            context.Context
  ScalpingClient pb.ScalpingClient
}

func (srv *Scalping) Client() pb.ScalpingClient {
  if srv.ScalpingClient == nil {
    conn, err := grpc.Dial(
      fmt.Sprintf("%v:%v", os.Getenv("CRYPTOS_GRPC_HOST"), os.Getenv("CRYPTOS_GRPC_PORT")),
      grpc.WithTransportCredentials(insecure.NewCredentials()),
    )
    if err != nil {
      panic(err.Error())
    }
    srv.ScalpingClient = pb.NewScalpingClient(conn)
  }
  return srv.ScalpingClient
}

func (srv *Scalping) Pagenate(symbol string, status []uint32, page int, pageSize int) (*pb.PagenateReply, error) {
  r, err := srv.Client().Pagenate(srv.Ctx, &pb.PagenateRequest{
    Symbol:   symbol,
    Status:   status,
    Page:     int32(page),
    PageSize: int32(pageSize),
  })
  if err != nil {
    return nil, err
  }
  return r, nil
}
