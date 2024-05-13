package metrics

import (
	"testing"

	"github.com/mcrgnt/yp1/internal/storage"
	"github.com/stretchr/testify/assert"
)

func TestPollMetrics(t *testing.T) {
	type args struct {
		params *PollMetricsParams
	}
	tests := []struct {
		args args
		name string
	}{
		// TODO: Add test cases.
		{
			name: "test0",
			args: args{
				params: &PollMetricsParams{
					Storage: storage.NewMemStorage(&storage.NewMemStorageParams{}),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PollMetrics(tt.args.params)
			assert.NotEqual(t, storage.Memory{}, tt.args.params.Storage)
		})
	}
}
