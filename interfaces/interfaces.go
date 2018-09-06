package interfaces

type DbHandler interface {
	GetByKeyAndCas(key string) (map[string]interface{}, uint64, error)
	Replace(data map[string]interface{}, cas uint64) (uint64, error)
}
