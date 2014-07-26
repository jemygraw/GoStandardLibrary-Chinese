# strconv包

import "strconv"

---

##简介 

>对于`javascript`等动态语言，parseInt等函数都是必备的。可见strconv包的存在多么重要。

##概览
strconv包实现了string和其他基本数据类型的转换。


##内容

###常量
```go
const IntSize = intSize
```
IntSize是int或者uint类型的比特位数。

###变量
```go
var ErrRange = errors.New("value out of range")
```
ErrRange表明超过了目标类型的取值范围。比如将 "128" 转为 int8 就会返回这个错误。

```go
var ErrSyntax = errors.New("invalid syntax")
```
ErrSyntax表明不符合目标类型的语法。比如将""转为int类型会返回这个错误。

>然而，在返回错误的时候，不是直接将上面的变量值返回，而是通过构造一个 NumError 类型的 error 对象返回。

###func AppendBool
```go
func AppendBool(dst []byte, b bool) []byte
```
将布尔值 `b` 转换为字符串 "true" 或 "false"，然后将结果追加到 byte切片`dst`的尾部，返回追加后的byte切片。

###func AppendFloat
```go
func AppendFloat(dst []byte, f float64, fmt byte, prec int, bitSize int) []byte
```
TODO

###func AppendInt
```go
func AppendInt(dst []byte, i int64, base int) []byte
```
将整数`i`转为byte切片追加到目标byte切片中。最终，我们也可以通过返回的byte切片得到字符串。

###func AppendQuote
```go
func AppendQuote(dst []byte, s string) []byte
```
将字符串 `s`转换为“双引号”引起来的字符串， 并将结果追加到 dst 的尾部，返回追加后的 []byte, 其中的特殊字符将被转换为“转义字符”。

###func AppendQuoteRune
```go
func AppendQuoteRune(dst []byte, r rune) []byte
```
将 Unicode 字符`r`转换为“单引号”引起来的字符串， 并将结果追加到 dst 的尾部，返回追加后的 []byte。“特殊字符”将被转换为“转义字符”。

###func AppendQuoteRuneToASCII
```go
func AppendQuoteRuneToASCII(dst []byte, r rune) []byte
```
将 Unicode 字符转换为“单引号”引起来的 ASCII 字符串，并将结果追加到 dst 的尾部，返回追加后的 []byte， “非 ASCII 字符”和“特殊字符”将被转换为“转义字符”。

###func AppendQuoteToASCII
```go
func AppendQuoteToASCII(dst []byte, s string) []byte
```
将字符串 s 转换为“双引号”引起来的 ASCII 字符串，并将结果追加到 dst 的尾部，返回追加后的 []byte， “非 ASCII 字符”和“特殊字符”将被转换为“转义字符”。

###func AppendUint
```go
func AppendUint(dst []byte, i uint64, base int) []byte
```
将 uint 型整数 i 转换为字符串形式，并追加到 dst 的尾部，并返回追加后的 []byte。 base表示进位制。

###func Atoi
```go
func Atoi(s string) (i int, err error)
```
Atoi内部通过调用 ParseInt(s, 10, 0) 来实现的，是 ParseInt 的便捷版。

###func CanBackquote
```go
func CanBackquote(s string) bool
```
判断字符串 s 是否可以表示为一个单行的“反引号”字符串。字符串中不能含有控制字符（除了 \t）和“反引号”字符，否则返回 false。

###func FormatBool
```go
func FormatBool(b bool) string
```
根据输入的bool参数`b`，返回`'true'`或者`'false'`。

###func FormatFloat
```go
func FormatFloat(f float64, fmt byte, prec, bitSize int) string
```

###func FormatInt
```go
func FormatInt(i int64, base int) string
```
将int64类型`i`按照给定的基数`base`转换成字符串并返回。`base`的范围是2至36。返回值使用小写字母`'a'`至`'z'`来表示大于等于10的数字。

###func FormatUint
```go
func FormatUint(i uint64, base int) string
```
将int64类型`i`按照给定的基数`base`转换成字符串并返回。`base`的范围是2至36。返回值使用小写字母`'a'`至`'z'`来表示大于等于10的数字。

###func IsPrint
```go
func IsPrint(r rune) bool
```
IsPrint 判断 Unicode 字符 r 是否是一个可显示的字符。可否显示并不是你想象的那样，比如空格可以显示，而\t则不能显示。

> `'\t'` `'\n'`  `0` 均为不可显示。

###func Itoa
```go
func Itoa(i int) string
```
Itoa 内部直接调用 FormatInt(i, 10) 实现的，是一个更方便的函数。

>实际应用中，我们经常会遇到需要将字符串和整型连接起来，在Java中，可以通过操作符 "+" 做到。不过，在Go语言中，你需要将整型转为字符串类型，然后才能进行连接。FormatUint、FormatInt和 Itoa是strconv 包中的整型转字符串的相关函数。

###func ParseBool
```gp
func ParseBool(str string) (value bool, err error)
```
将字符串`str`解析成bool类型，接受的参数为：1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False。传入其他的值会返回错误。

###func ParseFloat
```go
func ParseFloat(s string, bitSize int) (f float64, err error)
```

###func ParseInt
```go
func ParseInt(s string, base int, bitSize int) (i int64, err error)
```
将字符串`s`按照所给定的基数`base`转换成int类型，并返回。如果`base==0`，`base`（即进制数）由`s`的前缀隐式指定。`0x`为16，'0'为8，其他为10。

参数`bitSize`指定了返回的int类型结果满足的比特位。比特位数0、8、16、32和64分别对应int、int8、int16、int32和int64。

返回的`err`构造了*NumError类型并且包含err.Num=s。如果`s`是空的或者包含无效的数字，err.Err=ErrSyntax，返回值为0,；如果和字符串`s`对应的返回值无法用给定的size转换成一个有符号整形，err.Err=ErrRange，返回值是`bitSize`能够表示的最大或最小值，而且是有符号的。

###func ParseUint
```go
func ParseUint(s string, base int, bitSize int) (n uint64, err error)
```
和ParseInt类似，只不过对应的是无符号整形。

>ParseInt、ParseUint 和Atoi三个函数都可以将字符串转换为int类型。

###func Quote
```go
func Quote(s string) string
```
Quote 将字符串 s 转换为“双引号”引起来的字符串。返回的字符串使用Go转义字符`\t, \n, \xFF, \u0100`表示控制字符和IsPrint函数所定义的不可打印字符。

>其中的特殊字符将被转换为“转义字符”，“不可显示的字符”将被转换为“转义字符”。

>我们称 `"golanghome.com"` 这种用双引号引起来的字符串为 Go 语言字面值字符串（Go string literal）。

###func QuoteRune
```go
func QuoteRune(r rune) string
```
将 Unicode 字符转换为“单引号”引起来的字符串，“特殊字符”将被转换为“转义字符”。

###func QuoteRuneToASCII
```go
func QuoteRuneToASCII(r rune) string
```
将 Unicode 字符转换为“单引号”引起来的 ASCII 字符串， “非 ASCII 字符”和“特殊字符”将被转换为“转义字符”

###func QuoteToASCII
```go
func QuoteToASCII(s string) string
```
将字符串 `s `转换为“双引号”引起来的 ASCII 字符串， “非 ASCII 字符”和“特殊字符”将被转换为“转义字符”。

###func Unquote
```go
func Unquote(s string) (t string, err error)
```
将“带引号的字符串” s 转换为常规的字符串（不带引号和转义字符）
。s 可以是“单引号”、“双引号”或“反引号”引起来的字符串（包括引号本身）。如果 s 是单引号引起来的字符串，则返回该该字符串代表的字符。

###func UnquoteChar
```go
func UnquoteChar(s string, quote byte) (value rune, multibyte bool, tail string, err error)
```

> UnquoteChar 将 s 中的第一个字符“取消转义”并解码
>
s：转义后的字符串
 quote：字符串使用的“引号符”（用于对引号符“取消转义”）
>
>value：    解码后的字符
 multibyte：value 是否为多字节字符
 tail：     字符串 s 除去 value 后的剩余部分
 error：    返回 s 中是否存在语法错误
>
 参数 quote 为“引号符”
如果设置为单引号，则 s 中允许出现 \' 字符，不允许出现单独的 ' 字符
如果设置为双引号，则 s 中允许出现 \" 字符，不允许出现单独的 " 字符
如果设置为 0，则不允许出现 \' 或 \" 字符，可以出现单独的 ' 或 " 字符

```go
type NumError

type NumError struct {
    Func string // 转换失败的函数 (ParseBool, ParseInt, ParseUint, ParseFloat)
    Num  string // 输入的字符串
    Err  error  // 错误转换的原因 (ErrRange, ErrSyntax)
}
```
如果类型转换失败的转换，可以用一个NumError结构类型的变量来记录。

###func (*NumError) Error
```go
func (e *NumError) Error() string
```
将错误打印出来。

>格式为：`"strconv." + e.Func + ": " + "parsing " + Quote(e.Num) + ": " + e.Err.Error()`