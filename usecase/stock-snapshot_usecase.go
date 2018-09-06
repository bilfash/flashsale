package usecase

import (
	"fmt"

	"github.com/bilfash/flashsale/domain"
	"github.com/bilfash/flashsale/exception"
)

type StockSnapshot struct {
	domain.StockSnapshot
}

func (s *StockSnapshot) AdjustQty(qty int) error {
	if s.AvailableQty+qty <= 0 {
		return exception.NewNotEnoughQtyExc(fmt.Sprint("%s qty not enough", s.Sku))
	}
	return nil
}
