package pref

import (
	"reflect"
	"testing"
)

func TestUnmarshalPref(t *testing.T) {
	type args struct {
		buf string
	}
	tests := []struct {
		name    string
		args    args
		want    *Preference
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				buf: `
				version = 1
				[server]
				host = "0.0.0.0"
				port = 8888
				log_level = "info"
				[[templates]]
				platform = "platform1"
				source = "source1"
				`,
			},
			want: &Preference{
				Version: 1,
				ServerConfig: Server{
					Host:     "0.0.0.0",
					Port:     8888,
					LogLevel: "info",
				},
				Templates: []Template{
					{
						Platform: "platform1",
						Source:   "source1",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "2",
			args: args{
				buf: `
				invalid toml
				`,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UnmarshalPref([]byte(tt.args.buf))
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalPref() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnmarshalPref() = %v, want %v", got, tt.want)
			}
		})
	}
}
