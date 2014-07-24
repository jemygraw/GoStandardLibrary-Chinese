# strconv包

import "strconv"

---

##简介 

##概览
strconv包实现了string和其他基本数据类型的转换。对于`javascript`等动态语言，parseInt、toString等函数都是必备的。可见strconv包的存在多么重要。

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
ErrRange表明超过了目标类型的取值范围。

```go
var ErrSyntax = errors.New("invalid syntax")
```
ErrSyntax表明不符合目标类型的语法。

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
TODO

###func AppendQuoteRune
```go
func AppendQuoteRune(dst []byte, r rune) []byte
```
TODO

###func AppendQuoteRuneToASCII
```go
func AppendQuoteRuneToASCII(dst []byte, r rune) []byte
```
TODO

###func AppendQuoteToASCII
```go
func AppendQuoteToASCII(dst []byte, s string) []byte
```
TODO

###func AppendUint
```go
func AppendUint(dst []byte, i uint64, base int) []byte
```

###func Atoi
```go
func Atoi(s string) (i int, err error)
```

###func CanBackquote
```go
func CanBackquote(s string) bool
```

###func FormatBool
```go
func FormatBool(b bool) string
```

###func FormatFloat
```go
func FormatFloat(f float64, fmt byte, prec, bitSize int) string
```

###func FormatInt
```go
func FormatInt(i int64, base int) string
```

###func FormatUint
```go
func FormatUint(i uint64, base int) string
```

###func IsPrint
```go
func IsPrint(r rune) bool
```

###func Itoa
```go
func Itoa(i int) string
```

###func ParseBool
```gp
func ParseBool(str string) (value bool, err error)
```

###func ParseFloat
```go
func ParseFloat(s string, bitSize int) (f float64, err error)
```

###func ParseInt
```go
func ParseInt(s string, base int, bitSize int) (i int64, err error)
```

###func ParseUint
```go
func ParseUint(s string, base int, bitSize int) (n uint64, err error)
```

###func Quote
```go
func Quote(s string) string
```

###func QuoteRune
```go
func QuoteRune(r rune) string
```

###func QuoteRuneToASCII
```go
func QuoteRuneToASCII(r rune) string
```

###func QuoteToASCII
```go
func QuoteToASCII(s string) string
```

###func Unquote
```go
func Unquote(s string) (t string, err error)
```

###func UnquoteChar
```go
func UnquoteChar(s string, quote byte) (value rune, multibyte bool, tail string, err error)
```

```go
type NumError

type NumError struct {
    Func string // the failing function (ParseBool, ParseInt, ParseUint, ParseFloat)
    Num  string // the input
    Err  error  // the reason the conversion failed (ErrRange, ErrSyntax)
}
```

###func (*NumError) Error
```go
func (e *NumError) Error() string
```