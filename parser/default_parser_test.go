package parser

import (
	"copy-paste-detector/config"
	"fmt"
	"testing"
)

var _defaultConf *config.Conf

func TestMain(m *testing.M) {
	_defaultConf = &config.Conf{
		MinRepeatLine: 0,
		ParseFolder:   "../../copy-paste-detector",
	}
	m.Run()
}

func TestDefaultParser_getFiles(t *testing.T) {
	type fields struct {
		conf *config.Conf
	}
	tests := []struct {
		name   string
		fields fields
		want   []*File
	}{
		{
			fields: fields{
				conf: _defaultConf,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := NewDefaultParser(tt.fields.conf).(*DefaultParser)
			got := d.getFiles()
			for _, file := range got {
				fmt.Printf("file=%#v\n", file.Folder+file.FileName)
			}
		})
	}
}

func TestDefaultParser_parseResult(t *testing.T) {
	type fields struct {
		conf   *config.Conf
		result *Result
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			fields: fields{
				conf: _defaultConf,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := NewDefaultParser(tt.fields.conf).(*DefaultParser)
			d.getFiles()
			d.parseResult()
			for k := range d.result.RepeatLines {
				fmt.Printf("%s\n", k)
			}
			fmt.Printf("%s\n", "=====ids=======")
			for k := range d.result.IDLines {
				fmt.Printf("%d\n", k)
			}
			fmt.Printf("%s\n", "=======lines======")
			fmt.Printf("len(d.result.IDLines): %v\n", len(d.result.IDLines))
			fmt.Printf("len(d.result.RepeatLines): %v\n", len(d.result.RepeatLines))
		})
	}
}

func TestDefaultParser_Parse(t *testing.T) {
	type fields struct {
		conf   *config.Conf
		result *Result
	}
	tests := []struct {
		name   string
		fields fields
		want   *Result
	}{
		{
			fields: fields{
				conf: _defaultConf,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := NewDefaultParser(tt.fields.conf)
			r := d.Parse()
			fmt.Printf("len(r.IDLines): %v\n", len(r.IDLines))
		})
	}
}
