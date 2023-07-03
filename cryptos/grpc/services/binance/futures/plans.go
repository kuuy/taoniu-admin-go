package futures

import (
  "context"
  "fmt"
  "os"

  "google.golang.org/grpc"
  "google.golang.org/grpc/credentials/insecure"

  pb "taoniu.local/admin/cryptos/grpc/binance/futures/plans"
)

type Plans struct {
  Ctx         context.Context
  PlansClient pb.PlansClient
}

func (srv *Plans) Client() pb.PlansClient {
  if srv.PlansClient == nil {
    conn, err := grpc.Dial(
      fmt.Sprintf("%v:%v", os.Getenv("CRYPTOS_GRPC_HOST"), os.Getenv("CRYPTOS_GRPC_PORT")),
      grpc.WithTransportCredentials(insecure.NewCredentials()),
    )
    if err != nil {
      panic(err.Error())
    }
    srv.PlansClient = pb.NewPlansClient(conn)
  }
  return srv.PlansClient
}

func (srv *Plans) Pagenate(symbol string, current int, pageSize int) (*pb.PagenateReply, error) {
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
