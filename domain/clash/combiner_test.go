package clash

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestStringListCombiner(t *testing.T) {
	type args struct {
		p  interface{}
		ap interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "1",
			args: args{
				p: []string{
					"11",
					"22",
				},
				ap: []string{
					"33",
					"44",
				},
			},
			want: func() interface{} {
				return []string{
					"11",
					"22",
					"33",
					"44",
				}
			}(),
		},
		{
			name: "2",
			args: args{
				p: []string{
					"11",
					"22",
				},
				ap: []string{},
			},
			want: func() interface{} {
				return []string{
					"11",
					"22",
				}
			}(),
		},
		{
			name: "3",
			args: args{
				p: []string{},
				ap: []string{
					"33",
					"44",
				},
			},
			want: func() interface{} {
				return []string{
					"33",
					"44",
				}
			}(),
		},
		{
			name: "4",
			args: args{
				p: []string{
					"11",
					"22",
				},
				ap: nil,
			},
			want: func() interface{} {
				return []string{
					"11",
					"22",
				}
			}(),
		},
		{
			name: "5",
			args: args{
				p: nil,
				ap: []string{
					"33",
					"44",
				},
			},
			want: func() interface{} {
				return []string{
					"33",
					"44",
				}
			}(),
		},
		{
			name: "6",
			args: args{
				p:  nil,
				ap: nil,
			},
			want: []interface{}{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringListCombiner(tt.args.p, tt.args.ap); !cmp.Equal(got, tt.want) {
				t.Errorf(cmp.Diff(got, tt.want))
			}
		})
	}
}

func BenchmarkStringListCombiner(b *testing.B) {
	p := []string{
		"11",
		"22",
	}
	ap := []string{
		"33",
		"44",
	}
	for n := 0; n < b.N; n++ {
		StringListCombiner(p, ap)
	}
}
