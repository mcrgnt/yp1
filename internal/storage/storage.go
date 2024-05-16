package storage

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
		storage = NewMemory()
	default:
		storage = NewMemory()
	}
	return
}
