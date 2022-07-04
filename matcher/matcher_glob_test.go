package matcher

import (
	"reflect"
	"testing"
)

func TestGlobMatcher_Match(t *testing.T) {
	type fields struct {
		contains []string
	}
	type args struct {
		str string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			fields: fields{
				contains: []string{
					"*.go",
				},
			},
			args: args{
				str: "matcher.go",
			},
			want: true,
		},
		{
			fields: fields{
				contains: []string{
					"*.go",
				},
			},
			args: args{
				str: "matcher.123",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGlobMatcher(tt.fields.contains)
			if got := g.Match(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GlobMatcher.Match() = %v, want %v", got, tt.want)
			}
		})
	}
}
