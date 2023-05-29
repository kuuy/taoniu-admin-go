package daily

import (
  pb "taoniu.local/admin/cryptos/grpc/binance/spot/indicators/daily/ranking"
  services "taoniu.local/admin/cryptos/grpc/services/binance/spot/indicators/daily"
)

type RankingRepository struct {
  Service *services.Ranking
}

func (r *RankingRepository) Pagenate(
  symbol string,
  fields []string,
  sortField string,
  sortType int32,
  current int,
  pageSize int,
) (*pb.PagenateReply, error) {
  return r.Service.Pagenate(symbol, fields, sortField, sortType, current, pageSize)
}
