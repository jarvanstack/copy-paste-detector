package matcher

// 字符串匹配器
type Marcher interface {
	Match(str string) bool
}
