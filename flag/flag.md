# flag 包

import "flag"

---

##简介

##概览
flag包实现了命令行的参数解析。

使用flag.String()、Bool()、Int()等来定义flags。
以下声明了一个整形的flag。 `-flagname`存储在一个int型指针`ip`中。
```go
import "flag"
var ip = flag.Int("flagname", 1234, "help message for flagname")
```
>然后你通过`*ip`来获得输入的参数。

当然你可以这样，使用Var()函数将flag绑定到一个变量上面。
```go
var flagvar int
func init() {
	flag.IntVar(&flagvar, "flagname", 1234, "help message for flagname")
}
```
或者，你可以创建自定义的flags，它满足指针接收，然后将它们一起进行flag解析：
```go
flag.Var(&flagVal, "name", "help message for flagname")
```
对于这样的flags,默认值就是变量的初始值。
定义了所有的flags，然后调用`flag.Parse()`进行命令行参数解析。

>执行flag.Parse()是必须的操作。

Flags可以直接使用。如果你用flags本身，他们就都是指针，如果你将他们绑定到变量，他们都是值。

```go
fmt.Println("ip has value ", *ip)
fmt.Println("flagvar has value ", flagvar)
```
解析完成以后，flag之后的参数就成了flag.Args()切片或者单独的flag.Args(i)。参数的索引从0直到flag.NArg()-1。

命令行语法：
```go
-flag
-flag=x
-flag x  // 只支持非bool类型
```
一个或者两个减号都可以，是等效的。最后一种形式对于bool flags是禁止的，因为如果一个文件名为`0`、`false`，`cmd -x *`的意思会改变。

你必须使用`-flag=false`形式来关闭bool flag。在第一个non-flag参数之前或者在终止符`'--'`停止Flag解析。
整形flags接收1234, 0664, 0x1234 ，而且可以是负值。bool flags可以是：`1, 0, t, f, T, F, true, false, TRUE, FALSE, True, False`。

>布尔类型的参数防止解析时的二义性，应该使用等号的方式指定。

Duration flags接收任何符合time.ParseDuration的有效的输入参数。

顶层的函数控制默认的命令行flags集合。FlagSet的方法模拟了顶层函数定义了命令行 允许一个定义flags独立的集合，例如在一个命令行接口来实现子命令
。FlagSet的方法模拟了顶层函数定义了命令行flag集合。

##变量
```go
var CommandLine = NewFlagSet(os.Args[0], ExitOnError)
```
命令行是一个默认的命令行flags集合，从os.Args解析而来。顶层函数比如BoolVar、Arg和on都包含了命令行方法。
```go
var ErrHelp = errors.New("flag: help requested")
```
当flag-help被触发，但是没有flag定义，返回ErrHelp。

```go
var Usage = func() {
    fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
    PrintDefaults()
}
```
Usage函数为用户输出所有定义了的命令行参数和帮助信息`usage message`。一般，当命令行参数解析出错时，该函数会被调用。。这个函数是一个变量，可以指向一个自定义的函数。

###func Arg
```go
func Arg(i int) string
```
返回第`i`个命令行参数。Arg(0)是当flags被处理后的第一个保留参数。

###func Args
```go
func Args() []string
```
返回命令行中的non-flag参数。

###func Bool
```go
func Bool(name string, value bool, usage string) *bool
```

###func BoolVar
```go
func BoolVar(p *bool, name string, value bool, usage string)
```

###func Duration
```go
func Duration(name string, value time.Duration, usage string) *time.Duration
```

###func DurationVar
```go
func DurationVar(p *time.Duration, name string, value time.Duration, usage string)
```

###func Float64
```go
func Float64(name string, value float64, usage string) *float64
```

###func Float64Var
```go
func Float64Var(p *float64, name string, value float64, usage string)
```

###func Int
```go
func Int(name string, value int, usage string) *int
```
定义了一个int类型的flag并指定了一个具体的flag名字`name`，默认值是`value`，使用说明是`usage`，也就是给用户的提示信息。返回值是一个int指针，它所存储的地址用来保存flag的值。

###func Int64
```go
func Int64(name string, value int64, usage string) *int64
```

###func Int64Var
```go
func Int64Var(p *int64, name string, value int64, usage string)
```

###func IntVar
```go
func IntVar(p *int, name string, value int, usage string)
```
定义了一个int类型的flag并指定了一个具体的flag名字`name`，默认值是`value`，使用说明是`usage`，也就是给用户的提示信息。指针`p`指向的变量用来存储flag的值。

>IntVar是绑定函数，适用的情形是已经预先声明了存储flag信息的变量。IntVar一般可以在包的Init()函数中执行。

###func NArg
```go
func NArg() int
```
当flags被处理后，返回non-flag参数的个数。

###func NFlag
```go
func NFlag() int
```
返回命令行中已经被设置的flags个数。

>获得FlagSet中actual长度（即被设置了的参数个数）。

###func Parse
```go
func Parse()
```
从os.Args[1:]解析命令行参数。必须在所有的flag都已经定义好并且在flags信息到达程序之前调用。

>内部通过调用CommandLine.Parse(os.Args[1:])实现。

>go标准库中，经常这么做：
定义了一个类型，提供了很多方法；为了方便使用，会在在全局变量中实例化一个该类型的实例，这样便可以直接使用该实例调用方法。

>在flag包中，进行了进一步封装：将FlagSet的方法都重新定义了一遍，也就是提供了一系列函数，而函数中只是简单的调用已经实例化好了的FlagSet实例CommandLine的 的方法，这样FlagSet实例便不需要export。这样，使用者是这么调用：flag.Parse()而不是flag.NewFlagSet(x,y).Parse()。

###func Parsed
```go
func Parsed() bool
```
当命令行flags已经被解析时返回true。

###func PrintDefaults
```go
func PrintDefaults()
```
向标准错误输出打印所有已经定义的命令行参数的默认值。

###func Set
```go
func Set(name, value string) error
```
给命名为`name`的命令行参数设置值`value`。

###func String
```go
func String(name string, value string, usage string) *string
```
定义了一个string类型的flag并指定了一个具体的flag名字`name`，默认值是`value`，使用说明是`usage`，也就是给用户的提示信息。返回值是一个string指针，它所存储的地址用来保存flag的值。

###func StringVar
```go
func StringVar(p *string, name string, value string, usage string)
```

###func Uint
```go
func Uint(name string, value uint, usage string) *uint
```

###func Uint64
```go
func Uint64(name string, value uint64, usage string) *uint64
```

###func Uint64Var
```go
func Uint64Var(p *uint64, name string, value uint64, usage string)
```

###func UintVar
```go
func UintVar(p *uint, name string, value uint, usage string)
```

###func Var
```go
func Var(value Value, name string, usage string)
```
定义了一个flag，并指定了具体的`name`和`usage`。flag的类型和值用第一个参数来指出，它实例化了用户自定义的类型。例如，调用者可以创建一个flag，它通过给切片Value方法将用逗号分隔的字符串转换为一个string切片。特别的，集合应该将用逗号分隔的字符串转换为一个string切片。

###func Visit
```go
func Visit(fn func(*Flag))
```
按照字典顺序访问命令行参数，并对每个执行`fn`函数。只访问那些已经设置的flags。

###func VisitAll
```go
func VisitAll(fn func(*Flag))
```
按照字典顺序访问命令行参数，并对每个执行`fn`函数。访问所有的flags，无论是否已经设置。

###type ErrorHandling
```go
type ErrorHandling int
```
定义了如何处理flag解析错误。
```go
const (
    ContinueOnError ErrorHandling = iota
    ExitOnError
    PanicOnError
)
```
>三个常量在源码的FlagSet方法Parse()中使用了。

###type Flag
```go
type Flag struct {
    Name     string // 命令行参数-flagname名字flagname
    Usage    string // 帮助信息
    Value    Value  // 获取的输入的值。Value是一个接口。
    DefValue string // 默认值
}
```
flag类型代表了一个flag的状态。
>比如：autogo -f abc.txt，代码flag.String(“f”, “a.txt”, “this is usage”)，则该Flag实例（可以通过flag.Lookup(“f”)获得）相应的值为：f, this is usage , abc.txt, a.txt。

###func Lookup
```go
func Lookup(name string) *Flag
```
给定命令行参数`name`并返回Flag。如果`name`不存在则返回nil。

###type FlagSet
```go
type FlagSet struct {
    // Usage 是一个函数，当解析flag出现错误的时候调用。
    // The field is a function (not a method) that may be changed to point to
    // a custom error handler.
    Usage func()
    // contains filtered or unexported fields
}
```

###func NewFlagSet
```go
func NewFlagSet(name string, errorHandling ErrorHandling) *FlagSet
```
返回一个新的空flag集合，它带有具体的名字和错误处理属性。

>由于FlagSet中的字段没有export，其他方式获得FlagSet实例后，比如：FlagSet{}或new(FlagSet)，应该调用Init()方法，初始化name和errorHandling。

###func (*FlagSet) Arg
```go
func (f *FlagSet) Arg(i int) string
```

###func (*FlagSet) Args
```go
func (f *FlagSet) Args() []string
```

###func (*FlagSet) Bool
```go
func (f *FlagSet) Bool(name string, value bool, usage string) *bool
```

###func (*FlagSet) BoolVar
```go
func (f *FlagSet) BoolVar(p *bool, name string, value bool, usage string)
```

###func (*FlagSet) Duration
```go
func (f *FlagSet) Duration(name string, value time.Duration, usage string) *time.Duration
```

###func (*FlagSet) DurationVar
```go
func (f *FlagSet) DurationVar(p *time.Duration, name string, value time.Duration, usage string)
```

###func (*FlagSet) Float64
```go
func (f *FlagSet) Float64(name string, value float64, usage string) *float64
```

###func (*FlagSet) Float64Var
```go
func (f *FlagSet) Float64Var(p *float64, name string, value float64, usage string)
```

###func (*FlagSet) Init
```go
func (f *FlagSet) Init(name string, errorHandling ErrorHandling)
```

###func (*FlagSet) Int
```go
func (f *FlagSet) Int(name string, value int, usage string) *int
```

###func (*FlagSet) Int64
```go
func (f *FlagSet) Int64(name string, value int64, usage string) *int64
```

###func (*FlagSet) Int64Var
```go
func (f *FlagSet) Int64Var(p *int64, name string, value int64, usage string)
```

###func (*FlagSet) IntVar
```go
func (f *FlagSet) IntVar(p *int, name string, value int, usage string)
```

###func (*FlagSet) Lookup
```go
func (f *FlagSet) Lookup(name string) *Flag
```

###func (*FlagSet) NArg
```go
func (f *FlagSet) NArg() int
```

###func (*FlagSet) NFlag
```go
func (f *FlagSet) NFlag() int
```

###func (*FlagSet) Parse
```go
func (f *FlagSet) Parse(arguments []string) error
```
从参数列表中解析flag定义，不应该包含命令名字。必须在FlagSet中的所有flags都被定义并且flags被程序获取到之后调用。如果-help被设置但是没有定义，返回ErrHelp。


###func (*FlagSet) Parsed
```go
func (f *FlagSet) Parsed() bool
```

###func (*FlagSet) PrintDefaults
```go
func (f *FlagSet) PrintDefaults()
```

###func (*FlagSet) Set
```go
func (f *FlagSet) Set(name, value string) error
```

###func (*FlagSet) SetOutput
```go
func (f *FlagSet) SetOutput(output io.Writer)
```
设置usage和错误信息的输出位置。如果`output`为空，那么使用默认的os.Stderr。

###func (*FlagSet) String
```go
func (f *FlagSet) String(name string, value string, usage string) *string
```

###func (*FlagSet) StringVar
```go
func (f *FlagSet) StringVar(p *string, name string, value string, usage string)
```

###func (*FlagSet) Uint
```go
func (f *FlagSet) Uint(name string, value uint, usage string) *uint
```

###func (*FlagSet) Uint64
```go
func (f *FlagSet) Uint64(name string, value uint64, usage string) *uint64
```

###func (*FlagSet) Uint64Var
```go
func (f *FlagSet) Uint64Var(p *uint64, name string, value uint64, usage string)
```

###func (*FlagSet) UintVar
```go
func (f *FlagSet) UintVar(p *uint, name string, value uint, usage string)
```

###func (*FlagSet) Var
```go
func (f *FlagSet) Var(value Value, name string, usage string)
```

###func (*FlagSet) Visit
```go
func (f *FlagSet) Visit(fn func(*Flag))
```

###func (*FlagSet) VisitAll
```go
func (f *FlagSet) VisitAll(fn func(*Flag))
```

###type Getter
```go
type Getter interface {
    Value
    Get() interface{}
}
```

###type Value
```go
type Value interface {
    String() string
    Set(string) error
}
```

Value是一个接口，为了存取flag中的动态数据而定义。默认值用字符串来代替。

如果Value有一个IsBoolFlag()布尔方法返回true, 命令行解析器使得 -name 和-name=true等效。