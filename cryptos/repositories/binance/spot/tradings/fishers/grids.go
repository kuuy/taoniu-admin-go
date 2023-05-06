package fishers

import (
  pb "taoniu.local/admin/cryptos/grpc/binance/spot/tradings/fishers/grids"
  services "taoniu.local/admin/cryptos/grpc/services/binance/spot/tradings/fishers"
)

type GridsRepository struct {
  Service *services.Grids
}

func (r *GridsRepository) Pagenate(page int, pageSize int) (*pb.PagenateReply, error) {
  return r.Service.Pagenate(page, pageSize)
}