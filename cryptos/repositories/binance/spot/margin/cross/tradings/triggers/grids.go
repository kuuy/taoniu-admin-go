package triggers

import (
  pb "taoniu.local/admin/cryptos/grpc/binance/spot/margin/cross/tradings/triggers/grids"
  services "taoniu.local/admin/cryptos/grpc/services/binance/spot/margin/cross/tradings/triggers"
)

type GridsRepository struct {
  Service *services.Grids
}

func (r *GridsRepository) Pagenate(
  symbol string,
  status []uint32,
  current int,
  pageSize int,
) (*pb.PagenateReply, error) {
  return r.Service.Pagenate(symbol, status, current, pageSize)
}
