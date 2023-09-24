package files

import (
	"os"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGet(t *testing.T) {
	LICENSE, _ := os.ReadFile("../../LICENSE")

	type args struct {
		urlOrPath string
	}
	tests := []struct {
		name    string
		args    args
		wantBuf []byte
	}{
		{
			name: "1",
			args: args{
				urlOrPath: "https://raw.githubusercontent.com/zapsaang/conf_factory/master/LICENSE",
			},
			wantBuf: LICENSE,
		},
		{
			name: "2",
			args: args{
				urlOrPath: "../../LICENSE",
			},
			wantBuf: LICENSE,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotBuf := Get(tt.args.urlOrPath); !reflect.DeepEqual(gotBuf, tt.wantBuf) {
				t.Errorf("Get() = %v, want %v", gotBuf, tt.wantBuf)
			}
		})
	}
}

func TestGetAll(t *testing.T) {
	LICENSE, _ := os.ReadFile("../../LICENSE")

	README, _ := os.ReadFile("../../README.md")

	type args struct {
		urlOrPaths []string
	}
	tests := []struct {
		name string
		args args
		want [][]byte
	}{
		{
			name: "1",
			args: args{
				[]string{
					"https://raw.githubusercontent.com/zapsaang/conf_factory/master/LICENSE",
					"../../README.md",
					"https://raw.githubusercontent.com/zapsaang/conf_factory/master/README.md",
					"../../LICENSE",
					"https://raw.githubusercontent.com/zapsaang/conf_factory/master/LICENSE",
					"../../README.md",
					"https://raw.githubusercontent.com/zapsaang/conf_factory/master/README.md",
					"../../LICENSE",
					"https://raw.githubusercontent.com/zapsaang/conf_factory/master/README.md",
					"../../LICENSE",
					"https://raw.githubusercontent.com/zapsaang/conf_factory/master/LICENSE",
					"../../README.md",
					"../../LICENSE",
					"https://raw.githubusercontent.com/zapsaang/conf_factory/master/LICENSE",
					"../../README.md",
					"https://raw.githubusercontent.com/zapsaang/conf_factory/master/README.md",
					"../../LICENSE",
				},
			},
			want: [][]byte{
				LICENSE,
				README,
				README,
				LICENSE,
				LICENSE,
				README,
				README,
				LICENSE,
				README,
				LICENSE,
				LICENSE,
				README,
				LICENSE,
				LICENSE,
				README,
				README,
				LICENSE,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAll(tt.args.urlOrPaths); !reflect.DeepEqual(got, tt.want) {
				t.Error(cmp.Diff(got, tt.want))
				for i := range got {
					if !reflect.DeepEqual(got[i], tt.want[i]) {
						t.Errorf("GetAll() = %v, want %v", string(got[i]), string(tt.want[i]))
					}
				}
			}
		})
	}
}
