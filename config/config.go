package config

// 解析传入的配置
type Conf struct {
	MinRepeatLine int      // 最小重复行
	ParseFolder   string   // 解析的文件夹路径
	Ignore        []string // 过滤的正则表达式
	Contain       []string // 匹配的正则表达式文件, 优先级高于 Ignore
}
