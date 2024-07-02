package store

import (
	"github.com/mcrgnt/yp1/internal/store/memory"
	"github.com/mcrgnt/yp1/internal/store/models"
)

type NewStorageParams struct {
	Type string
}

func NewStorage(params *NewStorageParams) (storage models.Storage) {
	switch params.Type {
	case "memory":
		storage = memory.NewMemoryStorage()
	default:
		storage = memory.NewMemoryStorage()
	}
	return storage
}
