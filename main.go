package main

import (
	"copy-paste-detector/config"
	"copy-paste-detector/detector"
	"copy-paste-detector/out"
	"copy-paste-detector/parser"
)

func main() {
	conf := config.LoadConfig()
	parse := parser.NewDefaultParser(conf)
	pr := parse.Parse()
	d := detector.NewDefaultDetector(conf, pr)
	dr := d.Detect()
	o := out.NewConsoleOut(dr)
	o.Output()
}
