# strings包

import "strings"

---

##简介

strings包内包含字符串或字符的查找、替换、拆分、连接、映射、大小写转换、去头去尾等方法。

##概览

strings包实现了操作字符串的简单函数。

##内容

###func Contains
```go
func Contains(s, substr string) bool
```
当且仅当字符串`s`包含字符串`substr`时，返回 true。

###func ContainsAny
```go
func ContainsAny(s, chars string) bool
```
如果`s`包含`chars`中的任意一个字符，则返回true。

###func ContainsRune
```go
func ContainsRune(s string, r rune) bool
```
如果`s`包含`r`中的任意一个rune类型字符，则返回true。

###func Count
```go
func Count(s, sep string) int
```
返回字符串`s`中包含的字符串`sep`中的字符的数目。特别注意如果`sep`为空格时的情形。

###func EqualFold
```go
func EqualFold(s, t string) bool
```
比较字符串`s`和`t`，当采用UTF-8编码，两者全部小写时若相等，则返回ture。

###func Fields
```go
func Fields(s string) []string
```
按字符串`s`内的空格或者连续的空格来切割字符串`s`，并返回子串组成的切片。空格由unicode.IsSpace来定义。如果`s`只包含空格，那么返回空切片。

###func FieldsFunc
```go
func FieldsFunc(s string, f func(rune) bool) []string
```
按字符串`s`内的满足某种要求的rune字符来切割字符串`s`，并返回子串组成的切片。这种要求通过函数`f`来自定义。如果s只包含空格，那么返回空切片。

###func HasPrefix
```go
func HasPrefix(s, prefix string) bool
```
判断字符串`s`是否以字符串`prefix`为前缀。如果是，则返回true。

###func HasSuffix
```go
func HasSuffix(s, suffix string) bool
```
判断字符串`s`是否以字符串`prefix`为后缀。如果是，则返回true。

###func Index
```go
func Index(s, sep string) int
```
返回字符串`sep`在字符串`s`中第一次出现处的索引。如果`s`中不包含`sep`，则返回-1。

###func IndexAny
```go
func IndexAny(s, chars string) int
```
返回字符串`chars`中的任意字符在字符串`s`中第一次出现处的索引。如果`s`中不包含`chars`中的所有字符，则返回-1。

###func IndexByte
```go
func IndexByte(s string, c byte) int
```
返回byte字符`c`在字符串`s`中第一次出现处的索引。如果`s`中不包含`c`，则返回-1。

###func IndexFunc
```go
func IndexFunc(s string, f func(rune) bool) int
```
返回符合某种要求的字符在字符串`s`中第一次出现处的索引，这种要求由函数`f`自定义。如果`s`中不存在这样的字符，则返回-1。

###func IndexRune
```go
func IndexRune(s string, r rune) int
```
返回rune字符`c`在字符串`s`中第一次出现处的索引。如果`s`中不包含`r`，则返回-1。

###func Join
```go
func Join(a []string, sep string) string
```
把string切片连接成一个字符串并返回，切片元素用字符串`sep`连接。

###func LastIndex
```go
func LastIndex(s, sep string) int
```
返回字符串`sep`在字符串`s`中最后一次出现处的索引。如果`s`中不包含`sep`，则返回-1。

###func LastIndexAny
```go
func LastIndexAny(s, chars string) int
```
返回字符串`chars`中的任意字符在字符串`s`中最后一次出现处的索引。如果`s`中不包含`chars`中的所有字符，则返回-1。

###func LastIndexFunc
```go
func LastIndexFunc(s string, f func(rune) bool) int
```
返回符合某种要求的字符在字符串`s`中最后一次出现处的索引，这种要求由函数`f`自定义。如果`s`中不存在这样的字符，则返回-1。

###func Map
```go
func Map(mapping func(rune) rune, s string) string
```
按照某种规则将字符串`s`中的每个字符做映射处理，然后返回字符串`s`。这个规则通过函数`mapping`来自定义。如果`mapping`返回负数，则该字符被丢弃。

###func Repeat
```go
func Repeat(s string, count int) string
```
返回一个由`count`个字符串`s`组成的新字符串

###func Replace
```go
func Replace(s, old, new string, n int) string
```
将字符串`s`中的`n`个`old`字符串用`new`来替换。如果`n`小于0，则对要替换的`old`字符串个数没有限制。

###func Split
```go
func Split(s, sep string) []string
```
以字符串`sep`来拆分字符串`s`，然后返回由该字符串拆分形成的子字符串切片。如果`sep`为空，那么按照UTF-8编码分隔每一个字符，等效于SplitN函数当`count`取值为-1的情形。

###func SplitAfter
```go
func SplitAfter(s, sep string) []string
```
以字符串`sep`来拆分字符串`s`，然后返回由该字符串拆分形成的子字符串切片,除了最后的切片元素，每个元素以sep结尾。如果`sep`为空，那么按照UTF-8编码分隔每一个字符，等效于SplitN函数当`count`取值为-1的情形。

###func SplitAfterN
```go
func SplitAfterN(s, sep string, n int) []string
```
TODO

###func SplitN
```go
func SplitAfterN(s, sep string, n int) []string
```
TODO
###func Title
```go
func Title(s string) string
```
将字符中串的每个单词的首字母大写，并返回字符串。

###func ToLower
```go
func ToLower(s string) string
```
将字符串的每个字符全部小写，然后返回字符串。

###func ToLowerSpecial
```go
func ToLowerSpecial(_case unicode.SpecialCase, s string) string
```
TODO

###func ToTitle
```go
func ToTitle(s string) string
```
将字符串的每个字符全部大写，然后返回字符串。

###func ToTitleSpecial
```go
func ToTitleSpecial(_case unicode.SpecialCase, s string) string
```
TODO

###func ToUpper
```go
func ToUpper(s string) string
```
将字符串的每个字符全部大写，然后返回字符串。

###func ToUpperSpecial
```go
func ToUpperSpecial(_case unicode.SpecialCase, s string) string
```
TODO

###func Trim
```go
func Trim(s string, cutset string) string
```
去掉字符串`s`的头部和尾部的一些字符并返回字符串。这些字符由字符串`cutset`自定义。

###func TrimFunc
```go
func TrimFunc(s string, f func(rune) bool) string
```
去掉字符串`s`的头部和尾部的一些字符并返回字符串。这些字符由满足函数f的字符自定义。

###func TrimLeft
```go
func TrimLeft(s string, cutset string) string
```
去掉字符串`s`的头部的一些字符并返回字符串。这些字符由字符串cutset自定义。

###func TrimLeftFunc
```go
func TrimLeftFunc(s string, f func(rune) bool) string
```
去掉字符串`s`的头部的一些字符并返回字符串。这些字符由满足函数`f`的字符自定义。

###func TrimPrefix
```go
func TrimPrefix(s, prefix string) string
```
去掉字符`s`串头部的字符串`prefix`并返回子字符串。如果`s`不以`prefix`为头部，字符串`s`保持不变。

###func TrimRight
```go
func TrimRight(s string, cutset string) string
```
去掉字符串`s`的尾部的一些字符并返回字符串。这些字符由字符串`cutset`自定义。

###func TrimRightFunc
```go
func TrimRightFunc(s string, f func(rune) bool) string
```
去掉字符串`s`的尾部的一些字符并返回字符串。这些字符由满足函数`f`的字符自定义。

###func TrimSpace
```go
func TrimSpace(s string) string
```
去掉头部和尾部的空白符，包括空格、制表符、换行符。

###func TrimSuffix
```go
func TrimSuffix(s, suffix string) string
```
去掉字符`s`串尾部的字符串`suffix`并返回子字符串。如果`s`不以`suffix`为尾部，字符串`s`保持不变。

###type Reader  struct
```go
type Reader struct {
    // contains filtered or unexported fields
}
```
实现了io.Reader, io.ReaderAt, io.Seeker, io.WriterTo, io.ByteScanner, and io.RuneScanner接口。

###func  NewReader
```go
func NewReader(s string) *Reader
```
读取字符串`s`，并返回一个Reader指针。类似于bytes.NewBuffer，但是更加高效并且是只读的。

###func (*Reader) Len
```go
func (r *Reader) Len() int
```
返回未读到的字符串的长度。

###func (*Reader) Read
```go
func (r *Reader) Read(b []byte) (n int, err error)
```
将Reader类型`r`中的数据读入字节切片中。实现了io.Reader接口。

###func (*Reader) ReadAt
```go
func (r *Reader) ReadAt(b []byte, off int64) (n int, err error)
```
将Reader类型`r`中的数据读入字节切片中，`off`指定了读取的偏移量。
###func (*Reader) ReadByte
```go
func (r *Reader) ReadByte() (b byte, err error)
```
从Reader类型r中读出一个字节`b`并返回。

###func (*Reader) ReadRune
```go
func (r *Reader) ReadRune() (ch rune, size int, err error)
```
TODO

###func (*Reader) Seek
```go
func (r *Reader) Seek(offset int64, whence int) (int64, error)
```
TODO

###func (*Reader) UnreadByte
```go
func (r *Reader) UnreadByte() error
```
TODO

###func (*Reader) UnreadRune
```go
func (r *Reader) UnreadRune() error
```
TODO

###func (*Reader) WriteTo
```go
func (r *Reader) WriteTo(w io.Writer) (n int64, err error)
```
将r中的数据读出并写入io.Writer对象w中。实现了io.WriterTo方法。

###type Replacer
```go
type Replacer struct {
    // contains filtered or unexported fields
}
```

###func NewReplacer
```go
func NewReplacer(oldnew ...string) *Replacer
```
传入多个成对的string参数old、new，返回一个Replacer指针。

###func (*Replacer) Replace
```go
func (r *Replacer) Replace(s string) string
```
按照`r`所指定的替换规则来替换字符串`s`中的字符并返回字符串。

###func (*Replacer) WriteString
```go
func (r *Replacer) WriteString(w io.Writer, s string) (n int, err error)
```
按照Replacer类型`r`所指定的替换规则替换字符串`s`中的字符，并将处理后的字符串写入io.Writer对象`w`中。

