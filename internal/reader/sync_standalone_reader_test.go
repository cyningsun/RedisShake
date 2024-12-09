package reader

import (
	"bytes"
	"testing"

	"github.com/mcuadros/go-defaults"
	"github.com/spf13/viper"
)

func TestSyncReaderOptions(t *testing.T) {
	type args struct {
		config string
	}

	tests := []struct {
		name string
		args args
		want *SyncReaderOptions
	}{
		{
			name: "TestReaderOptions",
			args: args{
				config: ``,
			},
			want: &SyncReaderOptions{
				TCPReaderBufSize: 4096,
				TCPWriterBufSize: 4096,
				RDBReaderBufSize: 4096,
				RDBWriterBufSize: 33554432,
			},
		},
		{
			name: "TestReaderOptions",
			args: args{
				config: `
tcp_reader_buf_size = 4095 # set to default value of bufio.NewReader
tcp_writer_buf_size = 4097  # set to default value of bufio.NewWriter
rdb_reader_buf_size = 4098 # set to default value of bufio.NewReader
rdb_writer_buf_size = 33554433 # 32*1024*1024+1
`,
			},
			want: &SyncReaderOptions{
				TCPReaderBufSize: 4095,
				TCPWriterBufSize: 4097,
				RDBReaderBufSize: 4098,
				RDBWriterBufSize: 33554433,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := viper.New()
			v.SetConfigType("toml")
			v.ReadConfig(bytes.NewBufferString(tt.args.config))
			got := new(SyncReaderOptions)
			defaults.SetDefaults(got)
			if err := v.Unmarshal(got); err != nil {
				t.Errorf("ScanReaderOptions() error = %v", err)
				return
			}

			if got.TCPReaderBufSize != tt.want.TCPReaderBufSize ||
				got.TCPWriterBufSize != tt.want.TCPWriterBufSize ||
				got.RDBReaderBufSize != tt.want.RDBReaderBufSize ||
				got.RDBWriterBufSize != tt.want.RDBWriterBufSize {
				t.Errorf("ScanReaderOptions() = %v, want %v", got, tt.want)
			}
		})
	}
}
