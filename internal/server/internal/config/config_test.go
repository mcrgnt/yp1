package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	tests := []struct {
		wantConfig *Config
		osEnv      map[string]string
		name       string
		wantErr    bool
	}{
		{
			name:       "test_a",
			wantConfig: &Config{Address: "localhost:8080", StorageType: ""},
			osEnv: map[string]string{
				"ADDRESS": "",
				"MEMORY":  "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for k, v := range tt.osEnv {
				err := os.Setenv(k, v)
				assert.Nil(t, err)
			}
			actual, actualErr := NewConfig()
			if tt.wantErr {
				assert.NotNil(t, actualErr)
			} else {
				assert.Equal(t, tt.wantConfig, actual)
			}
		})
	}
}
