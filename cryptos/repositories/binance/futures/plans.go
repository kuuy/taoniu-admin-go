package futures

import (
  pb "taoniu.local/admin/cryptos/grpc/binance/futures/plans"
  services "taoniu.local/admin/cryptos/grpc/services/binance/futures"
)

type PlansRepository struct {
  Service *services.Plans
}

func (r *PlansRepository) Pagenate(
  symbol string,
  current int,
  pageSize int,
) (*pb.PagenateReply, error) {
  return r.Service.Pagenate(symbol, current, pageSize)
}
