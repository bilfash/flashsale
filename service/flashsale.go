package service

import (
	"encoding/json"

	"github.com/bilfash/flashsale/interfaces"
	"github.com/bilfash/flashsale/usecase"
)

func FlashSale(producer *kafka.Producer, stockSnapshotRepo *interfaces.StockSnapshotRepo, message usecase.StockSnapshot) error {
	for true {
		stockSnapshotObj, cas, err := stockSnapshotRepo.FindStockSnapshot(message.Key)
		if err != nil {
			return err
		}
		err = stockSnapshotObj.AdjustQty(message.AvailableQty)
		if err != nil {
			return err
		}
		_, err = stockSnapshotRepo.UpdateStockSnapshot(stockSnapshotObj, cas)
		if err != nil {
			continue
		}
		break
	}

	messageByte, _ := json.Marshal(message)
	producer.SendMessage("FlashSaleSuccess", messageByte)
	return nil
}
