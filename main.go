package main

import (
	"copy-paste-detector/config"
	"copy-paste-detector/parser"
	"fmt"
)

func main() {
	conf := &config.Conf{
		MinRepeatLine: 0,
		ParseFolder:   "../copy-paste-detector",
	}
	d := parser.NewDefaultParser(conf)
	r := d.Parse()
	fmt.Printf("len(r.IDLines): %v\n", len(r.IDLines))
}
