package reader

import (
	"reflect"
	"testing"
)

func TestNewRDBReader(t *testing.T) {
	type args struct {
		opts *RdbReaderOptions
	}
	tests := []struct {
		name string
		args args
		want *rdbReader
	}{
		{
			name: "TestNewRDBReader",
			args: args{
				opts: &RdbReaderOptions{
					Filepath:      "/tmp",
					ReaderBufSize: 1,
				},
			},
			want: &rdbReader{
				bufSize: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotInterface := NewRDBReader(tt.args.opts)
			got, err := (gotInterface).(*rdbReader)
			if !err {
				t.Errorf("NewRDBReader() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got.bufSize, tt.want.bufSize) {
				t.Errorf("NewRDBReader() = %v, want %v", got, tt.want)
			}
		})
	}
}
