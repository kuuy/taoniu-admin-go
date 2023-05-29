package tradings

import (
  pb "taoniu.local/admin/cryptos/grpc/binance/spot/tradings/scalping"
  services "taoniu.local/admin/cryptos/grpc/services/binance/spot/tradings"
)

type ScalpingRepository struct {
  Service *services.Scalping
}

func (r *ScalpingRepository) Pagenate(
  symbol string,
  status []uint32,
  current int,
  pageSize int,
) (*pb.PagenateReply, error) {
  return r.Service.Pagenate(symbol, status, current, pageSize)
}
