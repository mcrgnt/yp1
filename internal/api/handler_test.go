package api

// func TestDefaultHandler_handleUpdate(t *testing.T) {
// 	type fields struct {
// 		storage storage.MemStorage
// 		ctx     *logger.Logger
// 	}
// 	type args struct {
// 		r *http.Request
// 	}
// 	tests := []struct {
// 		name             string
// 		fields           fields
// 		args             args
// 		wantStatusHeader int
// 		wantErr          bool
// 	}{
// 		{
// 			name: "test0",
// 			fields: fields{
// 				storage: storage.NewMemStorage(&storage.NewMemStorageParams{}),
// 				ctx:     logger.NewLogger(&logger.LoggerInitParams{}),
// 			},
// 			args: args{
// 				r: httptest.NewRequest(http.MethodGet, "/test0", nil),
// 			},
// 			wantStatusHeader: http.StatusBadRequest,
// 			wantErr:          true,
// 		},
// 		{
// 			name: "test1",
// 			fields: fields{
// 				storage: storage.NewMemStorage(&storage.NewMemStorageParams{}),
// 				ctx:     logger.NewLogger(&logger.LoggerInitParams{}),
// 			},
// 			args: args{
// 				r: httptest.NewRequest(http.MethodPost, "/test1", nil),
// 			},
// 			wantStatusHeader: http.StatusNotFound,
// 			wantErr:          true,
// 		},
// 		{
// 			name: "test2",
// 			fields: fields{
// 				storage: storage.NewMemStorage(&storage.NewMemStorageParams{}),
// 				ctx:     logger.NewLogger(&logger.LoggerInitParams{}),
// 			},
// 			args: args{
// 				r: httptest.NewRequest(http.MethodPost, "/update", nil),
// 			},
// 			wantStatusHeader: http.StatusNotFound,
// 			wantErr:          true,
// 		},
// 		{
// 			name: "test3",
// 			fields: fields{
// 				storage: storage.NewMemStorage(&storage.NewMemStorageParams{}),
// 				ctx:     logger.NewLogger(&logger.LoggerInitParams{}),
// 			},
// 			args: args{
// 				r: httptest.NewRequest(http.MethodPost, "/update/gauge//6.5", nil),
// 			},
// 			wantStatusHeader: http.StatusNotFound,
// 			wantErr:          true,
// 		},
// 		{
// 			name: "test4",
// 			fields: fields{
// 				storage: storage.NewMemStorage(&storage.NewMemStorageParams{}),
// 				ctx:     logger.NewLogger(&logger.LoggerInitParams{}),
// 			},
// 			args: args{
// 				r: httptest.NewRequest(http.MethodPost, "/update/t/Name/6.5", nil),
// 			},
// 			wantStatusHeader: http.StatusBadRequest,
// 			wantErr:          true,
// 		},
// 		{
// 			name: "test5",
// 			fields: fields{
// 				storage: storage.NewMemStorage(&storage.NewMemStorageParams{}),
// 				ctx:     logger.NewLogger(&logger.LoggerInitParams{}),
// 			},
// 			args: args{
// 				r: httptest.NewRequest(http.MethodPost, "/update/gauge/Name/6.5", nil),
// 			},
// 			wantStatusHeader: http.StatusOK,
// 			wantErr:          false,
// 		},
// 		{
// 			name: "test6",
// 			fields: fields{
// 				storage: storage.NewMemStorage(&storage.NewMemStorageParams{}),
// 				ctx:     logger.NewLogger(&logger.LoggerInitParams{}),
// 			},
// 			args: args{
// 				r: httptest.NewRequest(http.MethodPost, "/update/counter/Name/6.5", nil),
// 			},
// 			wantStatusHeader: http.StatusBadRequest,
// 			wantErr:          true,
// 		},
// 		{
// 			name: "test7",
// 			fields: fields{
// 				storage: storage.NewMemStorage(&storage.NewMemStorageParams{}),
// 				ctx:     logger.NewLogger(&logger.LoggerInitParams{}),
// 			},
// 			args: args{
// 				r: httptest.NewRequest(http.MethodPost, "/update/gauge/Name/1.7976931348623157e+309", nil),
// 			},
// 			wantStatusHeader: http.StatusBadRequest,
// 			wantErr:          true,
// 		},
// 		{
// 			name: "test8",
// 			fields: fields{
// 				storage: storage.NewMemStorage(&storage.NewMemStorageParams{}),
// 				ctx:     logger.NewLogger(&logger.LoggerInitParams{}),
// 			},
// 			args: args{
// 				r: httptest.NewRequest(http.MethodPost, "/update/counter/Name/1", nil),
// 			},
// 			wantStatusHeader: http.StatusOK,
// 			wantErr:          false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tr := &DefaultHandler{
// 				storage: tt.fields.storage,
// 				ctx:     tt.fields.ctx,
// 			}
// 			gotStatusHeader, err := tr.handleUpdate(tt.args.r, tr.parsePath(tt.args.r))
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("DefaultHandler.handleUpdate() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if gotStatusHeader != tt.wantStatusHeader {
// 				t.Errorf("DefaultHandler.handleUpdate() = %v, want %v", gotStatusHeader, tt.wantStatusHeader)
// 			}
// 		})
// 	}
// }

// func TestDefaultHandler_parsePath(t *testing.T) {
// 	type fields struct {
// 		storage storage.MemStorage
// 		ctx     *logger.Logger
// 	}
// 	type args struct {
// 		r *http.Request
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 		want   []string
// 	}{
// 		// TODO: Add test cases.
// 		{
// 			name: "test0",
// 			fields: fields{
// 				storage: storage.NewMemStorage(&storage.NewMemStorageParams{}),
// 				ctx:     logger.NewLogger(&logger.LoggerInitParams{}),
// 			},
// 			args: args{
// 				r: httptest.NewRequest(http.MethodPost, "/update/gauge/Name/1.7976931348623157e+309", nil),
// 			},
// 			want: []string{
// 				"update",
// 				"gauge",
// 				"Name",
// 				"1.7976931348623157e+309",
// 			},
// 		},
// 		{
// 			name: "test1",
// 			fields: fields{
// 				storage: storage.NewMemStorage(&storage.NewMemStorageParams{}),
// 				ctx:     logger.NewLogger(&logger.LoggerInitParams{}),
// 			},
// 			args: args{
// 				r: httptest.NewRequest(http.MethodPost, "/update/", nil),
// 			},
// 			want: []string{
// 				"update",
// 			},
// 		},
// 		{
// 			name: "test2",
// 			fields: fields{
// 				storage: storage.NewMemStorage(&storage.NewMemStorageParams{}),
// 				ctx:     logger.NewLogger(&logger.LoggerInitParams{}),
// 			},
// 			args: args{
// 				r: httptest.NewRequest(http.MethodPost, "/", nil),
// 			},
// 			want: []string{
// 				"",
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tr := &DefaultHandler{
// 				storage: tt.fields.storage,
// 				ctx:     tt.fields.ctx,
// 			}
// 			if got := tr.parsePath(tt.args.r); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("DefaultHandler.parsePath() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
