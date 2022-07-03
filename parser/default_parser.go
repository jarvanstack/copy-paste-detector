package parser

import (
	"bufio"
	"copy-paste-detector/util"
	"io"
	"io/ioutil"
	"os"
	"regexp"
)

// DefaultParser 默认解析器
type DefaultParser struct {
	conf   *Conf
	result *Result
}

func New(conf *Conf) Parser {
	return &DefaultParser{
		conf: conf,
		result: &Result{
			IDLines:     make(map[int]*LineNode),
			RepeatLines: make(map[string][]int),
		},
	}
}

// 获取所有符合条件的文件对象
//   层序遍历
func (d *DefaultParser) getFiles() []*File {
	// 文件夹
	folders := []string{
		d.conf.ParseFolder,
	}

	match, err := regexp.Compile(d.conf.Contain)
	if err != nil {
		panic(err)
	}
	contains, err := regexp.Compile(d.conf.Contains)
	if err != nil {
		panic(err)
	}

	// 遍历所有文件夹
	for len(folders) != 0 {
		// 第二层
		var secondFolders []string
		for _, folder := range folders {
			fis, err := ioutil.ReadDir(folder)
			if err != nil {
				panic(err)
			}
			for _, fi := range fis {
				// 添加文件夹
				if fi.IsDir() {
					secondFolders = append(secondFolders, folder+"/"+fi.Name())
					continue
				}
				file := &File{
					Folder:   folder,
					FileName: fi.Name(),
				}
				// 如果不包含就跳过
				path := []byte(file.Folder + "/" + file.FileName)
				if (d.conf.Contain != "" && !match.Match(path)) ||
					(d.conf.Contains != "" && contains.Match(path)) {
					continue
				}
				// 添加文件到结果
				d.result.Files = append(d.result.Files, file)
			}
		}
		folders = secondFolders
	}
	return d.result.Files
}

func (d *DefaultParser) parseResult() {
	// 获取获取结果集
	id := 1 // 自增序列号
	for _, f := range d.result.Files {
		f2, err := os.Open(f.Folder + "/" + f.FileName)
		defer f2.Close()
		if err != nil {
			panic(err)
		}
		reader := bufio.NewReader(f2)
		lineNum := 1
		var cur *LineNode
		for {
			lineBytes, _, err := reader.ReadLine()
			if err == io.EOF {
				break
			}
			pc := util.TrimAllSpace(string(lineBytes))
			line := &LineNode{
				ID:               id,
				LineNum:          lineNum,
				ProcessedContent: pc,
				OriginContent:    lineBytes,
				F:                f,
			}
			id++
			lineNum++
			// 链表
			if cur == nil {
				cur = line
				f.Head = cur
			} else {
				cur.Next = line
				line.Prev = cur
				cur = cur.Next
			}

			// 序列号map
			d.result.IDLines[line.ID] = line

			// 重复行
			d.result.RepeatLines[line.ProcessedContent] = append(d.result.RepeatLines[line.ProcessedContent], line.ID)
		}
	}
}

func (d *DefaultParser) Parse() *Result {
	// 获得文件对象
	d.getFiles()

	// 通过文件对象解析结果
	d.parseResult()

	return d.result
}
