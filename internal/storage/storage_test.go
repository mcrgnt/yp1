package storage

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMemStorage(t *testing.T) {
	type args struct {
		params *NewMemStorageParams
	}
	tests := []struct {
		wantMemStorage Storage
		args           args
		name           string
	}{
		{
			name: "test0",
			args: args{
				&NewMemStorageParams{},
			},
			wantMemStorage: &MemStorage{},
		},
		{
			name: "test1",
			args: args{
				&NewMemStorageParams{
					Type: "t",
				},
			},
			wantMemStorage: &MemStorage{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := NewStorage(tt.args.params)
			assert.Equal(t, reflect.TypeOf(tt.wantMemStorage), reflect.TypeOf(actual))
		})
	}
}
