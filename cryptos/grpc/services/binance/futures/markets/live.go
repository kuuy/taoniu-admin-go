package markets

import (
  "context"
  "fmt"
  "os"

  "google.golang.org/grpc"
  "google.golang.org/grpc/credentials/insecure"

  pb "taoniu.local/admin/cryptos/grpc/binance/futures/markets/live"
)

type Live struct {
  Ctx        context.Context
  LiveClient pb.LiveClient
}

func (srv *Live) Client() pb.LiveClient {
  if srv.LiveClient == nil {
    conn, err := grpc.Dial(
      fmt.Sprintf("%v:%v", os.Getenv("CRYPTOS_GRPC_HOST"), os.Getenv("CRYPTOS_GRPC_PORT")),
      grpc.WithTransportCredentials(insecure.NewCredentials()),
    )
    if err != nil {
      panic(err.Error())
    }
    srv.LiveClient = pb.NewLiveClient(conn)
  }
  return srv.LiveClient
}

func (srv *Live) Pagenate(symbol string, current int, pageSize int) (*pb.PagenateReply, error) {
  r, err := srv.Client().Pagenate(srv.Ctx, &pb.PagenateRequest{
    Symbol:   symbol,
    Current:  int32(current),
    PageSize: int32(pageSize),
  })
  if err != nil {
    return nil, err
  }
  return r, nil
}
