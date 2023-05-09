package markets

import (
  pb "taoniu.local/admin/cryptos/grpc/binance/spot/markets/live"
  services "taoniu.local/admin/cryptos/grpc/services/binance/spot/markets"
)

type LiveRepository struct {
  Service *services.Live
}

func (r *LiveRepository) Pagenate(
  symbol string,
  page int,
  pageSize int,
) (*pb.PagenateReply, error) {
  return r.Service.Pagenate(symbol, page, pageSize)
}
