package fishers

import (
  services "taoniu.local/admin/cryptos/grpc/services/binance/spot/analysis/tradings/fishers/chart"
)

type ChartRepository struct {
  Service *services.Chart
}

func (r *ChartRepository) Series(limit int) ([]string, error) {
  return r.Service.Series(limit)
}
