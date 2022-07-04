# copy paste detector

复制粘贴检查器

## 快速开始

clone 

```bash
git clone https://github.com/dengjiawen8955/copy-paste-detector.git
cd copy-paste-detector
```

配置文件 vim cpd.toml

```
minRepeatLine = 4  # 最小重复行
parseFolder   = "out"   # 解析的文件夹路径
Ignore        = [] # 过滤的正则表达式
Contain       = [] # 匹配的正则表达式文件, 优先级高于 Ignore
```

编译

```
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

```
$ ./copy-paste-detector 
@==========@
重复代码片段总数:1
@==========@1
重复代码片段:
b
c
d
3
@==========@1
重复代码片段文件总数2
1 out1.txt:2-5
2 out2.txt:6-9
```