package aof

import (
	"reflect"
	"testing"

	"RedisShake/internal/entry"
)

var mockEntryChan = make(chan *entry.Entry, 1)

func TestNewLoader(t *testing.T) {
	type args struct {
		filePath string
		bufSize  int
		ch       chan *entry.Entry
	}
	tests := []struct {
		name string
		args args
		want *Loader
	}{
		{
			name: "TestNewLoader",
			args: args{
				filePath: "test",
				bufSize:  1,
				ch:       mockEntryChan,
			},
			want: &Loader{
				filePath: "test",
				ch:       mockEntryChan,
				bufSize:  1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLoader(tt.args.filePath, tt.args.bufSize, tt.args.ch); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLoader() = %v, want %v", got, tt.want)
			}
		})
	}
}
