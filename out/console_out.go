package out

import (
	"copy-paste-detector/detector"
	"fmt"
)

type ConsoleOut struct {
	dr *detector.Result
}

func (c *ConsoleOut) Output() {
	// 总结
	fmt.Printf("%s\n", "@==========@")
	fmt.Printf("重复代码片段总数:%d\n", len(c.dr.BlockMatrix))
	// 分开
	for i, repeatBlocks := range c.dr.BlockMatrix {
		fmt.Printf("@==========@%d\n", i+1)
		fmt.Printf("%s\n", "重复代码片段:")
		node := repeatBlocks[0].Start
		for node != nil && node != repeatBlocks[0].End {
			fmt.Println(string(node.OriginContent))
			node = node.Next
		}
		fmt.Println(string(repeatBlocks[0].End.OriginContent))

		fmt.Printf("@==========@%d\n", i+1)
		fmt.Printf("重复代码片段文件总数%d\n", len(repeatBlocks))
		for fileNum, block := range repeatBlocks {
			fileName := block.Start.F.Folder + block.Start.F.FileName
			fmt.Printf("%d %s:%d-%d\n", fileNum+1, fileName, block.Start.LineNum, block.End.LineNum)
		}
	}
}

func NewConsoleOut(dr *detector.Result) Out {
	return &ConsoleOut{
		dr: dr,
	}
}
