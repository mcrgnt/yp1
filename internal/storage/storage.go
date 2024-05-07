package storage

type MemStorage interface {
	Update(params *StorageParams)
	Get(params *StorageParams)
}

type NewMemStorageParams struct {
	Type string
}

func NewMemStorage(params *NewMemStorageParams) (memStorage MemStorage) {
	switch params.Type {
	default:
		memStorage = NewMemory()
	}
	return
}
