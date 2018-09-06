package interfaces

import (
	"encoding/json"
	"fmt"

	"github.com/bilfash/flashsale/exception"
	"github.com/bilfash/flashsale/usecase"
)

type StockSnapshotRepo struct {
	DbHandler DbHandler
}

func NewStockSnapshotRepo(dbHandlers DbHandler) *StockSnapshotRepo {
	snapshotRepo := new(StockSnapshotRepo)
	snapshotRepo.DbHandler = dbHandlers
	return snapshotRepo
}

func (repo *StockSnapshotRepo) FindStockSnapshot(key string) (*usecase.StockSnapshot, uint64, error) {
	row, cas, err := repo.DbHandler.GetByKeyAndCas(key)
	if err != nil {
		return nil, 0, exception.NewNotFoundExc(fmt.Sprint("%s not found", key))
	}
	var res usecase.StockSnapshot
	temp, _ := json.Marshal(row)
	json.Unmarshal(temp, &res)
	return &res, cas, nil
}

func (repo *StockSnapshotRepo) UpdateStockSnapshot(stockSnapshot *usecase.StockSnapshot, cas uint64) (uint64, error) {
	var data map[string]interface{}
	temp, _ := json.Marshal(stockSnapshot)
	_ = json.Unmarshal(temp, &data)
	if cas, err := repo.DbHandler.Replace(data, cas); err != nil {
		return 0, exception.NewCasNotMatched()
	} else {
		return cas, nil
	}
}
