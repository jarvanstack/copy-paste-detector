package detector

import "copy-paste-detector/parser"

type DefaultDetector struct {
	pr *parser.Result
}

func NewDefaultDetector(pr *parser.Result) Detector {
	return &DefaultDetector{
		pr: pr,
	}
}

func (d *DefaultDetector) Detect() *Result {
	return nil
}
