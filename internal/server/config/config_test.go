package config

import (
	"flag"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	tests := []struct {
		name       string
		wantConfig *Config
		wantErr    bool
		osEnv      map[string]string
		osFlag     map[string]string
	}{
		{
			name:       "test_a",
			wantConfig: &Config{Address: "localhost:8090", StorageType: ""},
			osEnv: map[string]string{
				"ADDRESS": "",
				"MEMORY":  "",
			},
			osFlag: map[string]string{
				"a": "localhost:8090",
			},
		},
		// {
		// 	name:       "test_b",
		// 	wantConfig: &Config{Address: "localhost:8090", StorageType: ""},
		// 	osEnv: map[string]string{
		// 		"ADDRESS": "localhost:8090",
		// 		"MEMORY":  "",
		// 	},
		// 	osFlag: map[string]string{},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, actualErr := NewConfig()
			if tt.wantErr {
				assert.NotNil(t, actualErr)
			} else {
				for k, v := range tt.osEnv {
					err := os.Setenv(k, v)
					assert.Nil(t, err)
				}
				for k, v := range tt.osFlag {
					err := flag.Set(k, v)
					assert.Nil(t, err)
				}
				assert.Equal(t, tt.wantConfig, actual)
			}
		})
	}
}
