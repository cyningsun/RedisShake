package reader

import (
	"bytes"
	"testing"

	"github.com/mcuadros/go-defaults"
	"github.com/spf13/viper"
)

func TestScanReaderOptions(t *testing.T) {
	type args struct {
		config string
	}

	tests := []struct {
		name string
		args args
		want *ScanReaderOptions
	}{
		{
			name: "TestReaderOptions",
			args: args{
				config: ``,
			},
			want: &ScanReaderOptions{
				TCPReaderBufSize: 4096,
				TCPWriterBufSize: 4096,
			},
		},
		{
			name: "TestReaderOptions",
			args: args{
				config: `
tcp_reader_buf_size = 4095 # set to default value of bufio.NewReader
tcp_writer_buf_size = 4097  # set to default value of bufio.NewWriter
`,
			},
			want: &ScanReaderOptions{
				TCPReaderBufSize: 4095,
				TCPWriterBufSize: 4097,
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
				got.TCPWriterBufSize != tt.want.TCPWriterBufSize {
				t.Errorf("ScanReaderOptions() = %v, want %v", got, tt.want)
			}
		})
	}
}
