package detector

import (
	"copy-paste-detector/config"
	"copy-paste-detector/parser"
	"reflect"
	"testing"
)

func buildLineNode(strs []string) []*parser.LineNode {
	var resp []*parser.LineNode
	var head *parser.LineNode
	var cur *parser.LineNode
	for _, str := range strs {
		if head == nil {
			head = &parser.LineNode{ProcessedContent: str}
			cur = head
		} else {
			node := &parser.LineNode{ProcessedContent: str}
			cur.Next = node
			node.Prev = cur
			cur = cur.Next
		}
		resp = append(resp, cur)
	}
	return resp
}

func Test_buildLineNode(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			args: args{
				strs: []string{
					"1",
					"2",
					"3",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nodes := buildLineNode(tt.args.strs)
			head := nodes[0]
			for _, str := range tt.args.strs {
				if head.ProcessedContent != str {
					t.Errorf("node = %v, want %v", head, str)
				}
				head = head.Next
			}
		})
	}
}

func Test_getRepeatBlock(t *testing.T) {
	node1s := buildLineNode([]string{
		"1",
		"2",
		"3",
	})
	node2s := buildLineNode([]string{
		"1",
		"2",
		"4",
	})
	type args struct {
		head1         *parser.LineNode
		head2         *parser.LineNode
		minRepeatLine int
	}
	tests := []struct {
		name string
		args args
		want []*Block
	}{
		{
			args: args{
				head1:         node1s[0],
				head2:         node2s[0],
				minRepeatLine: 2,
			},
			want: []*Block{
				{
					Start: node1s[0],
					End:   node1s[1],
				},
				{
					Start: node2s[0],
					End:   node2s[1],
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getRepeatBlock(tt.args.head1, tt.args.head2, tt.args.minRepeatLine); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getRepeatBlock() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultDetector_Detect(t *testing.T) {
	file1 := buildLineNode([]string{
		"1",
		"2",
		"3",
	})
	file2 := buildLineNode([]string{
		"1",
		"2",
		"4",
	})
	file3 := buildLineNode([]string{
		"5",
		"1",
		"2",
	})
	file1[0].ID = 0
	file1[1].ID = 1
	file1[2].ID = 2
	file2[0].ID = 3
	file2[1].ID = 4
	file2[2].ID = 5
	file3[0].ID = 6
	file3[1].ID = 7
	file3[2].ID = 8
	type fields struct {
		conf *config.Conf
		pr   *parser.Result
	}
	tests := []struct {
		name   string
		fields fields
		want   *Result
	}{
		{
			fields: fields{
				conf: &config.Conf{
					MinRepeatLine: 2,
				},
				pr: &parser.Result{
					Files: []*parser.File{
						{
							Head: file1[0],
						},
						{
							Head: file2[0],
						},
						{
							Head: file3[0],
						},
					},
					IDLines: map[int]*parser.LineNode{
						0: file1[0],
						1: file1[1],
						2: file1[2],
						3: file2[0],
						4: file2[1],
						5: file2[2],
						6: file3[0],
						7: file3[1],
						8: file3[2],
					},
					RepeatLines: map[string][]int{
						"1": {0, 3, 7},
						"2": {1, 4, 8},
						"3": {2},
						"4": {5},
						"5": {6},
					},
				},
			},
			want: &Result{
				[][]*Block{
					{
						{
							Start: file1[0],
							End:   file1[1],
						},
						{
							Start: file2[0],
							End:   file2[1],
						},
						{
							Start: file3[1],
							End:   file3[2],
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DefaultDetector{
				conf: tt.fields.conf,
				pr:   tt.fields.pr,
			}
			if got := d.Detect(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefaultDetector.Detect() = %v, want %v", got, tt.want)
			}
		})
	}
}
