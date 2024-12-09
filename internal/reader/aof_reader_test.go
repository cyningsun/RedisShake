package reader

import (
	"reflect"
	"testing"
)

func TestNewAOFReader(t *testing.T) {
	type args struct {
		opts *AOFReaderOptions
	}
	tests := []struct {
		name string
		args args
		want *aofReader
	}{
		{
			name: "TestNewAOFReader",
			args: args{
				opts: &AOFReaderOptions{
					Filepath:      "/tmp",
					AOFTimestamp:  1,
					ReaderBufSize: 1,
				},
			},
			want: &aofReader{
				bufSize: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotInterface := NewAOFReader(tt.args.opts)
			got, err := (gotInterface).(*aofReader)
			if !err {
				t.Errorf("NewAOFReader() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got.bufSize, tt.want.bufSize) {
				t.Errorf("NewAOFReader() = %v, want %v", got, tt.want)
			}
		})
	}
}
