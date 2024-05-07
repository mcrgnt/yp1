package storage

import (
	"testing"
)

func TestStorageParams_ValidateType(t *testing.T) {
	type fields struct {
		Type  string
		Name  string
		Value any
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "test0",
			fields: fields{
				Type:  "gauge",
				Name:  "test0",
				Value: "4",
			},
			wantErr: false,
		},
		{
			name: "test1",
			fields: fields{
				Type:  "counter",
				Name:  "test1",
				Value: "4",
			},
			wantErr: false,
		},
		{
			name: "test2",
			fields: fields{
				Type:  "t",
				Name:  "test2",
				Value: "4",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &StorageParams{
				Type:  tt.fields.Type,
				Name:  tt.fields.Name,
				Value: tt.fields.Value,
			}
			if err := tr.ValidateType(); (err != nil) != tt.wantErr {
				t.Errorf("StorageParams.ValidateType() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStorageParams_ValidateName(t *testing.T) {
	type fields struct {
		Type  string
		Name  string
		Value any
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "test0",
			fields: fields{
				Type:  "gauge",
				Name:  "test0",
				Value: "4",
			},
			wantErr: false,
		},
		{
			name: "test1",
			fields: fields{
				Type:  "gauge",
				Name:  "",
				Value: float64(4),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &StorageParams{
				Type:  tt.fields.Type,
				Name:  tt.fields.Name,
				Value: tt.fields.Value,
			}
			if err := tr.ValidateName(); (err != nil) != tt.wantErr {
				t.Errorf("StorageParams.ValidateName() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStorageParams_ValidateValue(t *testing.T) {
	type fields struct {
		Type  string
		Name  string
		Value any
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "test0",
			fields: fields{
				Type:  "gauge",
				Name:  "test0",
				Value: "4",
			},
			wantErr: false,
		},
		{
			name: "test1",
			fields: fields{
				Type:  "gauge",
				Name:  "test1",
				Value: "4.1",
			},
			wantErr: false,
		},
		{
			name: "test1+",
			fields: fields{
				Type:  "gauge",
				Name:  "test1+",
				Value: "-4.1",
			},
			wantErr: false,
		},
		{
			name: "test1++",
			fields: fields{
				Type:  "gauge",
				Name:  "test1++",
				Value: "1.7976931348623157e+309",
			},
			wantErr: true,
		},
		{
			name: "test2",
			fields: fields{
				Type:  "gauge",
				Name:  "test2",
				Value: "",
			},
			wantErr: true,
		},
		{
			name: "test3",
			fields: fields{
				Type:  "counter",
				Name:  "test3",
				Value: "4",
			},
			wantErr: false,
		},
		{
			name: "test3+",
			fields: fields{
				Type:  "counter",
				Name:  "test3+",
				Value: "-4",
			},
			wantErr: false,
		},
		{
			name: "test3++",
			fields: fields{
				Type:  "counter",
				Name:  "test3++",
				Value: "-402934802830498290384029384092830498203984209384092830984209384092830948209384028348972394857",
			},
			wantErr: true,
		},
		{
			name: "test4",
			fields: fields{
				Type:  "counter",
				Name:  "test4",
				Value: "4.4",
			},
			wantErr: true,
		},
		{
			name: "test5",
			fields: fields{
				Type:  "counter",
				Name:  "test5",
				Value: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &StorageParams{
				Type:  tt.fields.Type,
				Name:  tt.fields.Name,
				Value: tt.fields.Value,
			}
			if err := tr.ValidateValue(); (err != nil) != tt.wantErr {
				t.Errorf("StorageParams.ValidateValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
