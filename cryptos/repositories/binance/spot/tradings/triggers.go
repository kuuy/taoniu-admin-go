package tradings

import (
  pb "taoniu.local/admin/cryptos/grpc/binance/spot/tradings/triggers"
  services "taoniu.local/admin/cryptos/grpc/services/binance/spot/tradings"
)

type TriggersRepository struct {
  Service *services.Triggers
}

func (r *TriggersRepository) Pagenate(
  symbol string,
  status []uint32,
  page int,
  pageSize int,
) (*pb.PagenateReply, error) {
  return r.Service.Pagenate(symbol, status, page, pageSize)
}
