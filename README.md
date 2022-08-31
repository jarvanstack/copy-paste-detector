# copy paste detector

复制粘贴检查器

![基础架构](https://markdown-1304103443.cos.ap-guangzhou.myqcloud.com/2022-02-0420220831104153.png)

## 快速开始

clone

```bash
git clone https://github.com/dengjiawen8955/copy-paste-detector.git
cd copy-paste-detector
```

配置文件 vim cpd.toml

```toml
minRepeatLine = 4  # 最小重复行
parseFolder   = "out"   # 解析的文件夹路径
Ignore        = [] # 过滤的正则表达式
Contain       = [] # 匹配的正则表达式文件, 优先级高于 Ignore
```

编译

```bash
go build copy-paste-detector .
```

可选: 指定配置文件

```bash
$ ./copy-paste-detector -h
Usage of ./copy-paste-detector:
      --config string   config file (default "cpd.toml")
pflag: help requested
```

run

```bash
$ ./copy-paste-detector 
+------+------+--------------+
| 片段 | 序号 | 文件名:行数  |
+------+------+--------------+
|    1 |    1 | out1.txt:2-5 |
|    1 |    2 | out2.txt:6-9 |
+------+------+--------------+
b
c
d
3
```
