package local

type LocalSpike struct {
	LocalInStock     int64 // 本地库存
	LocalSalesVolume int64 // 本地销量
}

// LocalDeductStock 本地扣减库存
func (ls *LocalSpike) LocalDeductStock() bool {
	ls.LocalSalesVolume = ls.LocalSalesVolume + 1
	return ls.LocalSalesVolume <= ls.LocalInStock
}
