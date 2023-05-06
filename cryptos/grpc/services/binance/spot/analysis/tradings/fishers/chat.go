package fishers

import (
  "context"
  "fmt"
  "os"

  "google.golang.org/grpc"
  "google.golang.org/grpc/credentials/insecure"

  pb "taoniu.local/admin/cryptos/grpc/binance/spot/analysis/tradings/fishers/chart"
)

type Chart struct {
  Ctx         context.Context
  ChartClient pb.ChartClient
}

func (srv *Chart) Client() pb.ChartClient {
  if srv.ChartClient == nil {
    conn, err := grpc.Dial(
      fmt.Sprintf("%v:%v", os.Getenv("CRYPTOS_GRPC_HOST"), os.Getenv("CRYPTOS_GRPC_PORT")),
      grpc.WithTransportCredentials(insecure.NewCredentials()),
    )
    if err != nil {
      panic(err.Error())
    }
    srv.ChartClient = pb.NewChartClient(conn)
  }
  return srv.ChartClient
}

func (srv *Chart) Series(limit int) ([]string, error) {
  r, err := srv.Client().Series(srv.Ctx, &pb.SeriesRequest{Limit: int32(limit)})
  if err != nil {
    return nil, err
  }
  return r.Series, nil
}
