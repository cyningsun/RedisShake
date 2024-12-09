package rdb

import (
	"context"
	"reflect"
	"testing"

	"RedisShake/internal/entry"
)

// BenchmarkParseRDB is a benchmark for ParseRDB
// The baseline is "20	 350030327 ns/op	213804114 B/op	 1900715 allocs/op"
func BenchmarkParseRDB(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	tempChan := make(chan *entry.Entry, 1024)
	updateFunc := func(offset int64) {
	}
	b.N = 20

	for i := 0; i < b.N; i++ {
		loader := NewLoader("rdb", 4096, updateFunc, "./dump.rdb", tempChan)
		go func() {
			for temp := range tempChan {
				print(temp.CmdName)
			}
		}()
		loader.ParseRDB(context.Background())
	}
}

var mockEntryChan = make(chan *entry.Entry, 1)

func TestNewLoader(t *testing.T) {
	type args struct {
		name       string
		bufSize    int
		updateFunc func(int64)
		filPath    string
		ch         chan *entry.Entry
	}
	tests := []struct {
		name string
		args args
		want *Loader
	}{
		{
			name: "TestNewLoader",
			args: args{
				name:       "test",
				bufSize:    1,
				updateFunc: nil,
				filPath:    "fillpath",
				ch:         mockEntryChan,
			},
			want: &Loader{
				name:       "test",
				bufSize:    1,
				updateFunc: nil,
				filPath:    "fillpath",
				ch:         mockEntryChan,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLoader(tt.args.name, tt.args.bufSize, tt.args.updateFunc, tt.args.filPath, tt.args.ch); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLoader() = %v, want %v", got, tt.want)
			}
		})
	}
}
