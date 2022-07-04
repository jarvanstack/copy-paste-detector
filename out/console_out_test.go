package out

import (
	"copy-paste-detector/config"
	"copy-paste-detector/detector"
	"copy-paste-detector/parser"
	"testing"
)

func TestConsoleOut_Output(t *testing.T) {
	conf := &config.Conf{
		MinRepeatLine: 4,
		ParseFolder:   "../out",
	}
	p := parser.NewDefaultParser(conf)
	pr := p.Parse()
	d := detector.NewDefaultDetector(conf, pr)
	dr := d.Detect()
	o := NewConsoleOut(dr)
	o.Output()
}
