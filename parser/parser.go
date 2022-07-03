// 解析器
package parser

// Parser 解析器
type Parser interface {
	Parse() *Result
}

// 解析传入的配置
type Conf struct {
	MinRepeatLine int      // 最小重复行
	ParseFolder   string   // 解析的文件夹路径
	Ignore        []string // 过滤的正则表达式列表
}

type Result struct {
	Files       []*File
	IDLines     map[int]*LineNode
	RepeatLines map[string][]int
}

type File struct {
	Folder   string
	FileName string
	Head     *LineNode
}

type LineNode struct {
	ID               int       // 全局递增的序列号,用于降重
	LineNum          int       // 文件内部行号
	ProcessedContent string    // 处理后的文件内容
	OriginContent    []byte    // 原文件内容
	Prev             *LineNode // 上一行
	Next             *LineNode // 下一行
	F                *File     // 所属文件
}
