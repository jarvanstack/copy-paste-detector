package config

import (
	"reflect"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	tests := []struct {
		name string
		want *Conf
	}{
		{
			want: &Conf{
				MinRepeatLine: 5,
				ParseFolder:   "out",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LoadConfig().MinRepeatLine; !reflect.DeepEqual(got, tt.want.MinRepeatLine) {
				t.Errorf("LoadConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
