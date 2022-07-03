package util

import (
	"os"
	"os/exec"
	"path/filepath"
)

func Pwd() string {
	file, _ := exec.LookPath(os.Args[0])
	// 获取包含可执行文件名称的路径
	path, _ := filepath.Abs(file)
	return path
}
