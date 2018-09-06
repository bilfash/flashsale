package exception

type CasNotMatched struct {
	message string
}

func NewCasNotMatched() *CasNotMatched {
	return &CasNotMatched{
		message: "DML Error, possible causes include CAS mismatch or concurrent modification",
	}
}
func (e *CasNotMatched) Error() string {
	return e.message
}
