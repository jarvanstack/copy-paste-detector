package detector

import (
	"copy-paste-detector/parser"
)

type Detector interface {
	Detect() *Result
}

// 代码块
type Block struct {
	Start *parser.LineNode
	End   *parser.LineNode
}

// 重复代码块结果集
type Result struct {
	BlockMatrix [][]*Block
}
