package metrics

import (
	"fmt"
	"testing"

	"github.com/mcrgnt/yp1/internal/store/memory"
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
		{
			name: "test0",
			args: args{
				params: &PollMetricsParams{
					Storage: memory.NewMemoryStorage(),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := PollMetrics(tt.args.params); err != nil {
				fmt.Println(err)
			}
			assert.NotEqual(t, memory.MemoryStorage{}, tt.args.params.Storage)
		})
	}
}
