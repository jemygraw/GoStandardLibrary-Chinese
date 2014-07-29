#path

import "path"

##简介

包path实现了操作斜线分隔型路径的实用程序.


##概览

包path实现了操作斜线分隔型路径的实用程序.


###变量

```go
var ErrBadPattern = errors.New("syntax error in pattern")
// 表示匹配模型无法识别
```


###func Base
```go
func Base(path string) string
```
Base返回路径的最后一个元素.在提取最后一个元素之前已经去除了末尾斜线.如果path是空的,Base返回".".如果path只包含斜线,那么Base返回"/".


###func Clean
```go
func Clean(path string) string
```
Clean返回经过纯粹的词法处理后的跟path等效的最短路径.它迭代应用以下规则直到不能再进行更多处理:

1. 使用一个斜线替代多个斜线
2. 去除每个 . 路径名元素(当前目录)
3. 去除每个 .. 路径名元素(父目录),除非在 .. 之前没有其它元素.
4. 去除根目录开始的 .. 元素:就是,在路径开始部门替换"/.."为"/"

只有根目录"/"时返回的路径才会以斜线结尾.

如果处理后的结果是一个空字符串,Clean返回字符串".".

可以再参考Rob Pike的文档[Lexical File Names in Plan 9 or Getting Dot-Dot Right](http://plan9.bell-labs.com/sys/doc/lexnames.html)

###func Dir
```go
func Dir(path string) string
```
Dir返回path除了最后一个元素的其它部分,通常就是path的目录部分.使用Split去掉最后一个元素后,path会再进行Clean处理和去除末尾的斜线.如果path是空的,Dir返回".".  
如果path在非斜线字符后部分全部是斜线,那么Dir返回一个斜线.在其它任何情况下,返回的路径不会以斜线结尾.


###func Ext
```go
func Ext(path string) string
```
Ext返回路径中使用的文件扩展名.path以斜线进行分隔,最后一个元素的最后一个点之后的部分就是扩展;如果路径中没有点,扩展名就是空的.


###func IsAbs
```go
func IsAbs(path string) bool
```
IsAbs判断路径是否是绝对路径.


###func Join
```go
func Join(elem ...string) string
```
Join把任意数量的路径元素拼接成一个单一路径,根据需要加上分隔的斜线.结果是经过Clean处理;需要特别注意,全部空字符串会被忽略.


###func Match
```go
func Match(pattern, name string) (matched bool, err error)
```
Match判断name是否匹配shell文件名模板.模板语法是:

	pattern:
		{ term }
	term:
		'*'         matches any sequence of non-/ characters
		'?'         matches any single non-/ character
		'[' [ '^' ] { character-range } ']'
		            character class (must be non-empty)
		c           matches character c (c != '*', '?', '\\', '[')
		'\\' c      matches character c

	character-range:
		c           matches character c (c != '\\', '-', ']')
		'\\' c      matches character c
		lo '-' hi   matches character c for lo <= c <= hi
```

Match要求模板匹配名字的整个部分,不仅仅是其中一个子串.唯一有可能返回的错误是`ErrBadPattern`,当模板是错误的.


###func Split
```go
func Split(path string) (dir, file string)
```
Split根据最后一个斜线对路径进行切割,分成路径部分和文件名部分.如果是不带斜线的path参数,Split返回一个空的目录部分,整个参数path作为文件名.返回的值有这样的属性 path = dir + file.


