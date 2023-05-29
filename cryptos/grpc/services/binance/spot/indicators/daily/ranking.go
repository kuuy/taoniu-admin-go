package daily

import (
  "context"
  "fmt"
  "os"

  "google.golang.org/grpc"
  "google.golang.org/grpc/credentials/insecure"

  pb "taoniu.local/admin/cryptos/grpc/binance/spot/indicators/daily/ranking"
)

type Ranking struct {
  Ctx           context.Context
  RankingClient pb.RankingClient
}

func (srv *Ranking) Client() pb.RankingClient {
  if srv.RankingClient == nil {
    conn, err := grpc.Dial(
      fmt.Sprintf("%v:%v", os.Getenv("CRYPTOS_GRPC_HOST"), os.Getenv("CRYPTOS_GRPC_PORT")),
      grpc.WithTransportCredentials(insecure.NewCredentials()),
    )
    if err != nil {
      panic(err.Error())
    }
    srv.RankingClient = pb.NewRankingClient(conn)
  }
  return srv.RankingClient
}

func (srv *Ranking) Pagenate(
  symbol string,
  fields []string,
  sortField string,
  sortType int32,
  current int,
  pageSize int,
) (*pb.PagenateReply, error) {
  r, err := srv.Client().Pagenate(srv.Ctx, &pb.PagenateRequest{
    Symbol:    symbol,
    Fields:    fields,
    SortField: sortField,
    SortType:  sortType,
    Current:   int32(current),
    PageSize:  int32(pageSize),
  })
  if err != nil {
    return nil, err
  }
  return r, nil
}
