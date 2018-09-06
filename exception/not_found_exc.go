package exception

type NotFoundExc struct {
	message string
}

func NewNotFoundExc(message string) *NotFoundExc {
	return &NotFoundExc{
		message: message,
	}
}
func (e *NotFoundExc) Error() string {
	return e.message
}
