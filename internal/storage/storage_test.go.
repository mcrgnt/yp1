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
		wantMemStorage MemStorage
		args           args
		name           string
	}{
		{
			name: "test0",
			args: args{
				&NewMemStorageParams{},
			},
			wantMemStorage: &Memory{},
		},
		{
			name: "test1",
			args: args{
				&NewMemStorageParams{
					Type: "t",
				},
			},
			wantMemStorage: &Memory{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := NewMemStorage(tt.args.params)
			assert.Equal(t, reflect.TypeOf(tt.wantMemStorage), reflect.TypeOf(actual))
		})
	}
}
