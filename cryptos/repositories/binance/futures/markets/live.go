package markets

import (
  pb "taoniu.local/admin/cryptos/grpc/binance/futures/markets/live"
  services "taoniu.local/admin/cryptos/grpc/services/binance/futures/markets"
)

type LiveRepository struct {
  Service *services.Live
}

func (r *LiveRepository) Pagenate(
  symbol string,
  current int,
  pageSize int,
) (*pb.PagenateReply, error) {
  return r.Service.Pagenate(symbol, current, pageSize)
}
