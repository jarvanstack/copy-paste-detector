package detector

import (
	"copy-paste-detector/config"
	"copy-paste-detector/parser"
)

type DefaultDetector struct {
	conf *config.Conf
	pr   *parser.Result
}

func NewDefaultDetector(conf *config.Conf, pr *parser.Result) Detector {
	return &DefaultDetector{
		conf: conf,
		pr:   pr,
	}
}

func (d *DefaultDetector) Detect() *Result {
	resp := &Result{}

	for _, file := range d.pr.Files {
		node := file.Head
		for node != nil {
			if len(d.pr.RepeatLines[node.ProcessedContent]) > 1 {
				var repeatBlocks []*Block
				for _, lineID := range d.pr.RepeatLines[node.ProcessedContent] {
					// 序列号小于等于当前节点或者遍历过就跳过
					if lineID <= node.ID || d.pr.IDLines[lineID].Visited {
						continue
					}
					// 获取 2 个节点重复的代码块
					blocks := getRepeatBlock(node, d.pr.IDLines[lineID], d.conf.MinRepeatLine)
					// 返回空跳过
					if len(blocks) == 0 {
						continue
					}
					if len(repeatBlocks) == 0 {
						// 没有初始化就将原始值压入
						repeatBlocks = append(repeatBlocks, blocks...)
					} else {
						// 初始化过只需要压入新的值
						repeatBlocks = append(repeatBlocks, blocks[1])
					}
				}
				if len(repeatBlocks) != 0 {
					resp.BlockMatrix = append(resp.BlockMatrix, repeatBlocks)
					// 设置visit
					visit(repeatBlocks)
				}
			}
			node = node.Next
		}
	}

	return resp
}

func visit(repeatBlocks []*Block) {
	for _, block := range repeatBlocks {
		start := block.Start
		for start != block.End {
			start.Visited = true
			start = start.Next
		}
	}
}

// 获取2个节点重复的代码块
//   没有符合条件的重复代码块返回 nil
//   符合条件返回数组 blocks[0] 是 head1
func getRepeatBlock(head1, head2 *parser.LineNode, minRepeatLine int) []*Block {
	if (head1 == nil || head2 == nil) || (head1.ProcessedContent != head2.ProcessedContent) {
		return nil
	}

	cur1, cur2 := head1, head2
	size := 1

	for (cur1.Next != nil && cur2.Next != nil) && (cur1.Next.ProcessedContent == cur2.Next.ProcessedContent) {
		size++
		cur1 = cur1.Next
		cur2 = cur2.Next
	}

	if size >= minRepeatLine {
		return []*Block{
			{
				Start: head1,
				End:   cur1,
			},
			{
				Start: head2,
				End:   cur2,
			},
		}
	}

	return nil
}
