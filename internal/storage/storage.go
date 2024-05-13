package storage

type MemStorage interface {
	Update(params *StorageParams)
	GetByName(params *StorageParams)
	GetByType(params *StorageParams) error
	GetAll() string
}

type NewMemStorageParams struct {
	Type string
}

func NewMemStorage(params *NewMemStorageParams) (memStorage MemStorage) {
	switch params.Type {
	case "memory":
		memStorage = NewMemory()
	default:
		memStorage = NewMemory()
	}
	return
}
