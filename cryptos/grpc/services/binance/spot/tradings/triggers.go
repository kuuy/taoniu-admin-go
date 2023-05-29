package tradings

import (
  "context"
  "fmt"
  "os"

  "google.golang.org/grpc"
  "google.golang.org/grpc/credentials/insecure"

  pb "taoniu.local/admin/cryptos/grpc/binance/spot/tradings/triggers"
)

type Triggers struct {
  Ctx            context.Context
  TriggersClient pb.TriggersClient
}

func (srv *Triggers) Client() pb.TriggersClient {
  if srv.TriggersClient == nil {
    conn, err := grpc.Dial(
      fmt.Sprintf("%v:%v", os.Getenv("CRYPTOS_GRPC_HOST"), os.Getenv("CRYPTOS_GRPC_PORT")),
      grpc.WithTransportCredentials(insecure.NewCredentials()),
    )
    if err != nil {
      panic(err.Error())
    }
    srv.TriggersClient = pb.NewTriggersClient(conn)
  }
  return srv.TriggersClient
}

func (srv *Triggers) Pagenate(symbol string, status []uint32, current int, pageSize int) (*pb.PagenateReply, error) {
  r, err := srv.Client().Pagenate(srv.Ctx, &pb.PagenateRequest{
    Symbol:   symbol,
    Status:   status,
    Current:  int32(current),
    PageSize: int32(pageSize),
  })
  if err != nil {
    return nil, err
  }
  return r, nil
}
