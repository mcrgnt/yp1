package storage

type StorageParams struct {
	Value  any
	String string
	Type   string
	Name   string
}

type Storage interface {
	MetricSet(*StorageParams) error
	MetricReset(*StorageParams) error
	GetMetricString(*StorageParams) error
	GetMetricAll() string
}

type NewMemStorageParams struct {
	Type string
}

func NewStorage(params *NewMemStorageParams) (storage Storage) {
	switch params.Type {
	case "memory":
		storage = NewMemStorage()
	default:
		storage = NewMemStorage()
	}
	return
}
