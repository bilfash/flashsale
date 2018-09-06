package exception

type NotEnoughQtyExc struct {
	message string
}

func NewNotEnoughQtyExc(message string) *NotEnoughQtyExc {
	return &NotEnoughQtyExc{
		message: message,
	}
}
func (e *NotEnoughQtyExc) Error() string {
	return e.message
}
