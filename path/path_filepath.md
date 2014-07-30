#path/filepath

import "path/filepath"

##简介

filepath包实现了操作文件名路径的实用程序.

##概览

filepath包实现了操作文件名路径的实用程序,在一定程度上跟目标操作系统定义的文件目录相兼容.


###常量

```go
const (
    Separator     = os.PathSeparator
    ListSeparator = os.PathListSeparator
)
```

###变量

```go
var ErrBadPattern = errors.New("syntax error in pattern")
// 表示匹配模型无法识别

var SkipDir = errors.New("skip this directory")
// SkipDir作为在WalkFuncs中一个返回值,表示在当次调用中目录名将被忽略.
// 它不会被其它任何函数返回作为一个错误值.
```


###func Abs
```go
func Abs(path string) (string, error)
```
Abs返回path的绝对路径.如果path不是绝对路径,它会跟当前目录进行拼接成为绝对路径.对于给定的一个文件,其绝对路径不保证是唯一的.


###func Base
```go
func Base(path string) string
```
Base返回路径的最后一个元素.在提取最后一个元素之前已经去除了路径分隔符.如果path是空的,Base返回".".如果path只包含斜分隔符,那么Base返回一个路径分隔符.
    

###func Clean
```go
func Clean(path string) string
```
Clean返回经过纯粹的词法处理后的跟path等效的最短路径.它迭代应用以下规则直到不能再进行更多处理:

1. 使用一个分隔符替代多个分隔符.
2. 去除每个 . 路径名元素(当前目录)
3. 去除每个 .. 路径名元素(父目录),除非在 .. 之前没有其它元素.
4. 假设分隔符是"/",去除根目录开始的 .. 元素:就是,在路径开始部分替换"/.."为"/".

只有根目录时返回的路径才会以斜线结尾,比如说Unix上的"/"或者Windows上"C:\\".

如果处理后的结果是一个空字符串,Clean返回字符串".".

可以再参考Rob Pike的文档[Lexical File Names in Plan 9 or Getting Dot-Dot Right](http://plan9.bell-labs.com/sys/doc/lexnames.html)


###func Dir
```go
func Dir(path string) string
```
Dir返回path除了最后一个元素的其它部分,通常就是path的目录部分.使用Split去掉最后一个元素后,path会再进行Clean处理和去除末尾的斜线.如果path是空的,Dir返回".".  
如果path全部是由分隔符组成,那么Dir返回一个分隔符.除非是根目录,返回的路径不会以分隔符结尾.    


###func EvalSymlinks
```go
func EvalSymlinks(path string) (string, error)
```
EvalSymlinks访问符号链接,返回链接指向的路径.如果路径是相对的,那么结果就是相对于当前目录,除非其中一个组成部分是绝对符号链接.


###func Ext
```go
func Ext(path string) string
```
Ext返回路径中使用的文件扩展名.path以分隔符进行分隔,最后一个元素的最后一个点之后的部分就是扩展;如果路径中没有点,扩展名就是空的.


###func FromSlash
```go
func FromSlash(path string) string
```
FromSlash把path中的每一个斜线('/)替换为分隔符.多个斜线替换成多个分隔符.



###func Glob
```go
func Glob(pattern string) (matches []string, err error)
```
Glob返回所有匹配模板的文件名,如果没有任何匹配得上的文件名则返回nil.模板的语法跟Match的一样.模板可以描述层次结构的名称,例如 /usr/*/bin/ed (假设分隔符就是 '/').


###func HasPrefix
```go
func HasPrefix(p, prefix string) bool
```
HasPrefix是为了历史兼容性而保留,不要使用这个函数.


###func IsAbs
```go
func IsAbs(path string) bool
```
IsAbs判断路径是否是绝对路径.


###func Join
```go
func Join(elem ...string) string
```
Join把任意数量的路径元素拼接成一个单一路径,根据需要加上分隔符.结果是经过Clean处理;需要特别注意,全部空字符串会被忽略.


###func Match
```go
func Match(pattern, name string) (matched bool, err error)
```
Match判断name是否匹配shell文件名模板.模板语法是:

	pattern:
		{ term }
	term:
		'*'         matches any sequence of non-Separator characters
		'?'         matches any single non-Separator character
		'[' [ '^' ] { character-range } ']'
		            character class (must be non-empty)
		c           matches character c (c != '*', '?', '\\', '[')
		'\\' c      matches character c

	character-range:
		c           matches character c (c != '\\', '-', ']')
		'\\' c      matches character c
		lo '-' hi   matches character c for lo <= c <= hi


Match要求模板匹配名字的整个部分,不仅仅是其中一个子串.唯一有可能返回的错误是`ErrBadPattern`,当模板是错误的.

在Windows上,转义是禁用的. "\\\\"会本认为是路径分隔符. 


###func Rel
```go
func Rel(basepath, targpath string) (string, error)
```
Rel返回一个相对路径,词法上等同于targpath相对于basepath的路径.意味着Join(basepath, Rel(basepath, targpath))相当于targpath本身.当调用成功时,返回的路径总是相对于basepath,即使basepath和targpath没有任何共同元素.当targpath无法计算相对于basepath的目录,或者完成计算需要获取当前的工作目录这两种情况时,会返回一个错误.


###func Split
```go
func Split(path string) (dir, file string)
```
Split根据最后一个分隔符对路径进行切割,分成路径部分和文件名部分.如果是不带分隔符的path参数,Split返回一个空的目录部分,整个参数path作为文件名.返回的值有这样的属性 path = dir + file.


###func SplitList
```go
func SplitList(path string) []string
```
SplitList对一系列路径组成的path进行分割,分隔符是特定系统相关的ListSeparator,ListSeparator通常会在环境变量PATH或者GOPATH中看到.
跟strings.Split不一样,当参数是一个空字符串,SplitList返回一个空数组.


###func ToSlash
```go
func ToSlash(path string) string
```
ToSlash把path中的每一个分隔符替换为斜线('/').多个分隔符替换为多个斜线.


###func VolumeName
```go
func VolumeName(path string) (v string)
```
VolumeName返回起始的卷标名.在windows下给出"C:\foo\bar",返回"C:",给出 "\\\\host\share\foo" 返回 "\\\\host\share".在其它系统返回 "".


###func Walk
```go
func Walk(root string, walkFn WalkFunc) error
```
Walk遍历root为根的文件树,对于遍历到的每一个文件或者目录,包括root本身调用walkFn函数.在访问文件和目录的过程中发生的错误会先由walkFn进行过滤.所有文件按照词法顺序进行遍历,这保证了输出的稳定性,但是又意味着对于非常大的目录Walk的效率会很差.Walk不会跟随符号链接进行遍历.


###func WalkFunc
```go
type WalkFunc func(path string, info os.FileInfo, err error) error
```
WalkFunc是Walk访问每一个文件或者目录时调用的函数的类型.调用Walk的参数作为参数path的前缀部分;这意味着,如果调用Walk遍历目录"dir",目录里面包含了一个文件"a",将会使用参数"dir/a"调用遍历函数.参数info为path提供的os.FileInfo类型信息.

如果在遍历path指定的文件或者目录时出现了问题,传递进来的error会描述具体的问题,函数可以决定如何处理这个错误 (Walk不会继续进入那个目录).如果一个错误返回了,处理会停止.唯一例外的是,如果path是一个目录且函数返回了一个特殊值`SkipDir`,目录的所有内容会被忽略,后续处理会继续进行遍历下一个文件.