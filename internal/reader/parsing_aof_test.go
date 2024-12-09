package reader

import (
	"reflect"
	"testing"

	"RedisShake/internal/entry"
)

func TestNewAOFFileInfo(t *testing.T) {
	type args struct {
		aofFilePath string
		bufSize     int
		ch          chan *entry.Entry
	}
	tests := []struct {
		name string
		args args
		want *INFO
	}{
		{
			name: "TestNewAOFFileInfo",
			args: args{
				aofFilePath: "/tmp",
				bufSize:     1,
				ch:          nil,
			},
			want: &INFO{
				ReaderBufSize: 1,
				AOFDirName:    "/",
				AOFFileName:   "tmp",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAOFFileInfo(tt.args.aofFilePath, tt.args.bufSize, tt.args.ch); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAOFFileInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}
