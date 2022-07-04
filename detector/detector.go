package detector

import (
	"copy-paste-detector/parser"
	"text/template/parse"
)

type Detector interface {
	Detect() *Result
}

// 代码块
type Block struct {
	Start *parse.ListNode
	End   *parser.LineNode
}

// 重复的代码块
type Repeat struct {
	blocks []*Block
}

// 重复代码块结果集
type Result struct {
	Repeats []*Repeat
}
