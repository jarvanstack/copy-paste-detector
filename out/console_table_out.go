package out

import (
	"bytes"
	"copy-paste-detector/detector"
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

type ConsoleTableOut struct {
	dr *detector.Result
}

func (c *ConsoleTableOut) Output() {
	// 初始化 table
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"片段", "序号", "文件名:行数"})
	// 总结
	// fmt.Printf("重复代码片段总数:%d\n", len(c.dr.BlockMatrix))
	// 分开
	for i, repeatBlocks := range c.dr.BlockMatrix {
		block1 := fmt.Sprintf("%d", i+1)
		codeBuf := &bytes.Buffer{}
		node := repeatBlocks[0].Start
		for node != nil && node != repeatBlocks[0].End {
			codeBuf.WriteString(string(node.OriginContent) + "\n")
			node = node.Next
		}
		codeBuf.WriteString(string(repeatBlocks[0].End.OriginContent) + "\n")

		for fileNum, block := range repeatBlocks {
			fileName := block.Start.F.Folder + block.Start.F.FileName
			table.Append([]string{block1, fmt.Sprintf("%d", fileNum+1), fmt.Sprintf("%s:%d-%d", fileName, block.Start.LineNum, block.End.LineNum)})
		}
		table.Render()
		table.ClearRows()
		fmt.Print(codeBuf.String())
	}

}

func NewConsoleTableOut(dr *detector.Result) Out {
	return &ConsoleTableOut{
		dr: dr,
	}
}
