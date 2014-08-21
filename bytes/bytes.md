#bytes

import "bytes"

##简介

bytes包实现了操作字节切片的函数.它是类似于strings包的实现.

##概览

bytes包实现了操作字节切片的函数.它是类似于strings包的实现.


###常量

```go
const MinRead = 512
// MinRead是Buffer.ReadFrom调用Read方法时传递的最小切片大小.
// 只要Buffer拥有至少MinRead个字节,超过保存r的内容的长度需求,
// ReadFrom就不会扩大底层缓冲区.
```

   

###变量
```go
var ErrTooLarge = errors.New("bytes.Buffer: too large")
// ErrTooLarge在无法分配内存的时候传递给panic
```


###func Compare
```go
func Compare(a, b []byte) int
```
Compare比较两个字节切片然后返回一个数字.0表示a==b,-1表示a < b,+1表示 a > b.空切片跟nil参数相等.


###func Contains
```go
func Contains(b, subslice []byte) bool
```
Contains返回b是否包含subslice.


###func Count
```go
func Count(s, sep []byte) int
```
Count计算不重叠情况下s包含sep的次数.


###func Equal
```go
func Equal(a, b []byte) bool
```
Equal比较a跟b是否长度一样且包含相同的字节.空切片跟nil参数相等.



###func EqualFold
```go
func EqualFold(s, t []byte) bool
```
EqualFold比较字符串s和t，即当采用UTF-8编码，两者全部大写时若相等，则返回ture。


###func Fields
```go
func Fields(s []byte) [][]byte
```
按切片s内的空白字符或者连续的空白字符来切割切片s，并返回子切片组成的切片.如果s只包含空白字符，那么返回空切片。



###func FieldsFunc
```go
func FieldsFunc(s []byte, f func(rune) bool) [][]byte
```
FieldsFunc把s当成UTF-8编码的Unicode字符序列.按切片s内的满足某种要求的rune字符来切割切片s，并返回子切片组成的切片。这种要求通过函数f来自定义。如果s只包含空白字符，那么返回空切片。


###func HasPrefix
```go
func HasPrefix(s, prefix []byte) bool
```
判断字节切片s是否以字节切片prefix为前缀。如果是，则返回true。


###func HasSuffix
```go
func HasSuffix(s, suffix []byte) bool
```
判断字节切片s是否以字节切片suffix为后缀。如果是，则返回true。


###func Index
```go
func Index(s, sep []byte) int
```
返回切片sep在切片s中第一次出现处的索引。如果s中不包含sep，则返回-1。


###func IndexAny
```go
func IndexAny(s []byte, chars string) int
```
IndexAny把s当成UTF-8编码的Unicode字符序列.返回字符串chars中的任意字符在切片s中第一次出现处的索引。如果`s`中不包含`chars`中的任何字符，则返回-1



###func IndexByte
```go
func IndexByte(s []byte, c byte) int
```
返回byte字符c在切片s中第一次出现处的索引。如果s中不包含c，则返回-1。


###func IndexFunc
```go
func IndexFunc(s []byte, f func(r rune) bool) int
```
IndexAny把s当成UTF-8编码的Unicode字符序列.返回符合某种要求的字符在切片s中第一次出现处的字节索引，这种要求由函数f自定义。如果s中不存在这样的字符，则返回-1。



###func IndexRune
```go
func IndexRune(s []byte, r rune) int
```
IndexRune把s当成UTF-8编码的Unicode字符序列.返回rune字符r在切片s中第一次出现处的索引。如果s中不包含r，则返回-1。



###func Join
```go
func Join(s [][]byte, sep []byte) []byte
```
Join把[]byte切片连接成一个切片并返回，切片元素用切片sep连接。


###func LastIndex
```go
func LastIndex(s, sep []byte) int
```
返回切片sep在切片s中最后一次出现处的索引。如果s中不包含sep，则返回-1。


###func LastIndexAny
```go
func LastIndexAny(s []byte, chars string) int
```
IndexAny把s当成UTF-8编码的Unicode字符序列.返回字符串chars中的任意字符在切片s中最后一次出现处的索引。如果`s`中不包含`chars`中的任何字符，则返回-1


###func LastIndexFunc
```go
func LastIndexFunc(s []byte, f func(r rune) bool) int
```
LastIndexFunc把s当成UTF-8编码的Unicode字符序列.返回符合某种要求的字符在切片s中最后一次出现处的索引，这种要求由函数f自定义。如果s中不存在这样的字符，则返回-1。



###func Map
```go
func Map(mapping func(r rune) rune, s []byte) []byte
```
按照某种规则将字节切片s中的每个字符做映射处理，然后返回每个元素被映射后的字节数组。这个规则通过函数mapping来自定义。如果mapping返回负数，则该字符被丢弃.s和输出的字符都按照当UTF-8编码的Unicode字符.



###func Repeat
```go
func Repeat(b []byte, count int) []byte
```
返回一个由count个切片b组成的新切片.



###func Replace
```go
func Replace(s, old, new []byte, n int) []byte
```
将切片s中的n个old切片用new切片来替换,复制返回新切片。如果n小于0，则对要替换的old切片个数没有限制。



###func Runes
```go
func Runes(s []byte) []rune
```
Runes返回跟s相等的字符切片(Unicode编码).


###func Split
```go
func Split(s, sep []byte) [][]byte
```
以切片sep来拆分切片s，然后返回由该切片拆分形成的子字节切片。如果sep为空，那么按照UTF-8编码分隔每一个字符，等效于SplitN函数当count取值为-1的情形。



###func SplitAfter
```go
func SplitAfter(s, sep []byte) [][]byte
```
以切片sep来拆分切片s，然后返回由该切片拆分形成的子字节切片,除了最后的切片元素，每个元素以sep结尾（即保留sep）。如果sep为空，那么按照UTF-8编码分隔每一个字符，等效于SplitN函数当count取值为-1的情形。



###func SplitAfterN
```go
func SplitAfterN(s, sep []byte, n int) [][]byte
```
以切片sep来拆分切片s，然后返回由该切片拆分形成的子字节切片,除了最后的切片元素，每个元素以sep结尾（即保留sep）。如果sep为空，那么按照UTF-8编码分隔每一个字符。可以通过最后一个参数 n 控制返回的结果中的 slice 中的元素个数

* 当 n > 0 时，表示返回的 slice 中最多只有 n 个元素，其中，最后一个元素不会分割
* 当 n == 0 时，返回的结果是 nil
* 当 n < 0 时，返回所有的子字符串



###func SplitN
```go
func SplitN(s, sep []byte, n int) [][]byte
```
以切片sep来拆分切片s，然后返回由该切片拆分形成的子字节切片。如果sep为空，那么按照UTF-8编码分隔每一个字符。可以通过最后一个参数 n控制返回的结果中的 slice 中的元素个数

* 当 n > 0 时，表示返回的 slice 中最多只有 n 个元素，其中，最后一个元素不会分割
* 当 n == 0 时，返回的结果是 nil
* 当 n < 0 时，返回所有的子字符串


###func Title
```go
func Title(s []byte) []byte
```
复制输入字节切片,将字节切片中的每个单词的首字母转为大写，并返回复制处理后的字节切片.


###func ToLower
```go
func ToLower(s []byte) []byte
```
将切片的每个字符转换为小写，然后返回一个新的切片。



###func ToLowerSpecial
```go
func ToLowerSpecial(_case unicode.SpecialCase, s []byte) []byte
```
// TODO
ToLowerSpecial returns a copy of the byte slice s with all Unicode
letters mapped to their lower case, giving priority to the special
casing rules.

###func ToTitle
```go
func ToTitle(s []byte) []byte
```
ToTitle复制输入字节切片,把字节切片中的每一个字符转化为大写后返回.


###func ToTitleSpecial
```go
func ToTitleSpecial(_case unicode.SpecialCase, s []byte) []byte
```
// TODO
ToTitleSpecial returns a copy of the byte slice s with all Unicode
letters mapped to their title case, giving priority to the special
casing rules.

###func ToUpper
```go
func ToUpper(s []byte) []byte
```
ToUpper复制输入字节切片,把每一个字符转换为大写后返回.


###func ToUpperSpecial
```go
func ToUpperSpecial(_case unicode.SpecialCase, s []byte) []byte
```
// TODO    
ToUpperSpecial returns a copy of the byte slice s with all Unicode
letters mapped to their upper case, giving priority to the special
casing rules.


###func Trim
```go
func Trim(s []byte, cutset string) []byte
```
Trim去掉字节切片`s`的头部和尾部的一些UTF-8编码的字符并返回一个子切片,这些字符由字符串`cutset`自定义。


###func TrimFunc
```go
func TrimFunc(s []byte, f func(r rune) bool) []byte
```
TrimFunc去掉字节切片`s`的头部和尾部的一些UTF-8编码的字符并返回一个子切片,字符是否符合要求由函数f定义.


###func TrimLeft
```go
func TrimLeft(s []byte, cutset string) []byte
```
Trim去掉字节切片`s`的头部的一些UTF-8编码的字符并返回一个子切片,这些字符由字符串`cutset`自定义。


###func TrimLeftFunc
```go
func TrimLeftFunc(s []byte, f func(r rune) bool) []byte
```
TrimFunc去掉字节切片`s`的头部的一些UTF-8编码的字符并返回一个子切片,字符是否符合要求由函数f定义.


###func TrimPrefix
```go
func TrimPrefix(s, prefix []byte) []byte
```
TrimPrefix去掉切片`s`起始的切片`prefix`并返回子切片。如果`s`不以`prefix`为头部，返回的切片`s`保持不变。



###func TrimRight
```go
func TrimRight(s []byte, cutset string) []byte
```
Trim去掉字节切片`s`的尾部的一些UTF-8编码的字符并返回一个子切片,这些字符由字符串`cutset`自定义。


###func TrimRightFunc
```go
func TrimRightFunc(s []byte, f func(r rune) bool) []byte
```
TrimFunc去掉字节切片`s`的尾部的一些UTF-8编码的字符并返回一个子切片,字符是否符合要求由函数f定义.


###func TrimSpace
```go
func TrimSpace(s []byte) []byte
```
去掉头部和尾部的空白符，包括空格、制表符、换行符。


###func TrimSuffix
```go
func TrimSuffix(s, suffix []byte) []byte
```
TrimSuffix去掉切片`s`末尾的切片`suffix`并返回子切片。如果`s`不以`suffix`为结尾，返回的切片`s`保持不变。


###type Buffer
```go
type Buffer struct {
    // 包含了被过滤或者未被导出的字段
}
```
Buffer是一个变长的缓冲区,实现了Read和Write方法. 
// TODO: The zero value for Buffer is an empty buffer ready to use.


###func NewBuffer
```go
func NewBuffer(buf []byte) *Buffer
```
NewBuffer创建初始化一个Buffer,使用切片buf作为其初始化内容.它的目的是准备好一个缓冲区去读取已存在的数据.它也可以用于改变内部缓冲区的大小,这样做时,buf应该拥有所需的容量但是长度为0.

在大多数情况下,new(Buffer) (或者仅仅声明一个Buffer变量)就足够初始化一个Buffer.


###func NewBufferString
```go
func NewBufferString(s string) *Buffer
```
NewBufferString创建初始化一个Buffer,使用字符串s作为其初始化内容.它的目的是准备好一个缓冲区去读取一个已存在的字符串.

在大多数情况下,new(Buffer) (或者仅仅声明一个Buffer变量)就足够初始化一个Buffer.



###func (*Buffer) Bytes
```go
func (b *Buffer) Bytes() []byte
```
Bytes返回缓冲区中未读部分的内容.len(b.Bytes()) == b.Len(). 在没有其它对缓冲区的干预方法的情况下,如果调用者修改了返回切片中的内容,缓冲区中的内容会改变.


###func (*Buffer) Grow
```go
func (b *Buffer) Grow(n int)
```
Grow保证有足够的n个字节的空间,在有必要的情况下扩大缓冲区的容量.在调用Grow(n)之后,至少有n个字节可以写入缓冲区而不需要其它分配操作.如果n是负数,Grow会发生panic.如果缓冲区无法增长,会发生panic,抛出异常ErrTooLarge.


###func (*Buffer) Len
```go
func (b *Buffer) Len() int
```
Len返回未读字节数.b.Len() == len(b.Bytes())


###func (*Buffer) Next
```go
func (b *Buffer) Next(n int) []byte
```
Next返回一个切片包含了缓冲区中的n个字节,同时推进缓冲区就像是这些字节被Read读取了返回一样.如果缓冲区中少于n各字节,Next返回整个缓冲区.返回的切片只在后续其它读写操作之前有效.



###func (*Buffer) Read
```go
func (b *Buffer) Read(p []byte) (n int, err error)
```
Read从缓冲区读取长度为len(p)的字节,直到缓冲区读取完.返回值n是读取到的字节数.如果缓冲区中没有数据可以返回,返回错误是io.EOF(除非len(p)是0);其它情况下,返回值err是nil.



###func (*Buffer) ReadByte
```go
func (b *Buffer) ReadByte() (c byte, err error)
```
ReadByte读取返回缓冲区的下一个字节.如果没有更多的可读字节,返回错误io.EOF.


###func (*Buffer) ReadBytes
```go
func (b *Buffer) ReadBytes(delim byte) (line []byte, err error)
```
ReadBytes一直读取直到第一次在输入中遇到delim,返回一个字节数组,包含了读取到的字节和分隔符delim.如果ReadBytes在找到分隔符之前出现了错误,返回遇到错误之前已经读取到的数据和错误本身(通常就是io.EOF).ReadBytes当且仅当在数据不以delim为结尾的情况下返回err != nil.


###func (*Buffer) ReadFrom
```go
func (b *Buffer) ReadFrom(r io.Reader) (n int64, err error)
```
ReadFrom从r读取数据直到结束,把所读到的数据追加到缓冲区中,有需要的情况下会扩大缓冲区大小,返回值中n是读取的字节数.读取过程中遇到的任何错误,除了io.EOF会被返回.如果缓冲区变得太大,ReadFrom会发生panic,抛出异常ErrTooLarge.



###func (*Buffer) ReadRune
```go
func (b *Buffer) ReadRune() (r rune, size int, err error)
```
ReadRune读取缓冲区返回下一个UTF-8编码的Unicode字符.如果没有可用字节,返回的错误是io.EOF.如果字节是错误的UTF-8编码,函数消耗一个字节然后返回U+FFFD, 1.

    

###func (*Buffer) ReadString
```go
func (b *Buffer) ReadString(delim byte) (line string, err error)
```
ReadString一直读取直到第一次在输入中遇到delim,然后返回一个字符串line,line包含读取到的数据和分隔符.如果ReadString在找到分隔符之前出现了错误,返回遇到错误之前已经读取到的数据和错误本身(通常就是io.EOF).ReadString当且仅当在数据不以delim为结尾的情况下返回err != nil


###func (*Buffer) Reset
```go
func (b *Buffer) Reset()
```
Reset重置清空缓冲区.b.Reset()相当于b.Truncate(0).


###func (*Buffer) String
```go
func (b *Buffer) String() string
```
String把未读的内容转换为一个字符串后返回.如果Buffer是一个nil指针,它返回"<nil>".



###func (*Buffer) Truncate
```go
func (b *Buffer) Truncate(n int)
```
Truncate截断保留缓冲区上的n个未读字节.如果n是负数或者比缓冲区的大小还大,会发生panic操作.



###func (*Buffer) UnreadByte
```go
func (b *Buffer) UnreadByte() error
```
UnreadByte把最近一次读操作返回的字节标记为未读.如果在最后一次读操作后发生了写操作,UnreadByte返回一个错误.


###func (*Buffer) UnreadRune
```go
func (b *Buffer) UnreadRune() error
```
UnreadRune把ReadRune返回的字符标记为未读.如果最近一次在缓冲区上的读写操作不是ReadRune,UnreadRune会返回一个错误.(在这方面,它比UnreadByte严格,UnreadByte可以对任何读操作的最后一个字节进行操作)

  

###func (*Buffer) Write
```go
func (b *Buffer) Write(p []byte) (n int, err error)
```
Write追加切片p的内容到缓冲区,有需要的情况下会扩大缓冲区大小.返回值中n是切片p的大小,err永远都是nil.如果缓冲区变得太大,WriteByte会发生panic,抛出异常ErrTooLarge.


###func (*Buffer) WriteByte
```go
func (b *Buffer) WriteByte(c byte) error
```
WriteByte追加字节c到缓冲区,有需要的情况下会扩大缓冲区大小.返回的错误永远都是nil,目的是使接口跟bufio.Write的WriteByte保持一致.如果缓冲区变得太大,WriteByte会发生panic,抛出异常ErrTooLarge.



###func (*Buffer) WriteRune
```go
func (b *Buffer) WriteRune(r rune) (n int, err error)
```
WriteRune追加UTF-8编码的字符r到缓冲区,返回r的长度和一个错误err.err永远都是nil.返回一个err的目的是使接口跟bufio.Writer的WriteRune保持一致.有需要的情况下会扩大缓冲区大小.如果缓冲区变得太大,WriteRune会发生panic,抛出异常ErrTooLarge.



###func (*Buffer) WriteString
```go
func (b *Buffer) WriteString(s string) (n int, err error)
```
WriteString追加s的内容到缓冲区,有需要的情况下会扩大缓冲区大小.返回值n是s的长度,err总是nil.如果缓冲区变得太大,WriteString会发生panic,抛出异常ErrTooLarge.



###func (*Buffer) WriteTo
```go
func (b *Buffer) WriteTo(w io.Writer) (n int64, err error)
```
WriteTo写入数据到w,直到内部缓冲区输出完或者发生了错误.返回的n表示写入的字节数,n的类型是int64是为了匹配接口io.Writer.写入过程中遇到的任何错误会一并返回.


###type Reader
```go
type Reader struct {
    // contains filtered or unexported fields
}
```
Reader实现了io.Reader, io.ReaderAt, io.WriterTo, io.Seeker,io.ByteScanner, 和io.RuneScanner方法,它从一个字节切片读取数据.跟Buffer不一样,Reader只支持读操作,同时支持Seek操作.


###func NewReader
```go
func NewReader(b []byte) *Reader
```
NewReader返回一个Reader指针,从`b`读取数据.


###func (*Reader) Len
```go
func (r *Reader) Len() int
```
Len返回未读到的字节数.


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

###func (*Reader) ReadRune
```go
func (r *Reader) ReadRune() (ch rune, size int, err error)
```
从Reader类型r中读出一个字节`b`并返回。


###func (*Reader) Seek
```go
func (r *Reader) Seek(offset int64, whence int) (int64, error)
```
Seek实现了io.Seeker接口.
    

###func (*Reader) UnreadByte
```go
func (r *Reader) UnreadByte() error
```

###func (*Reader) UnreadRune
```go
func (r *Reader) UnreadRune() error
```

###func (*Reader) WriteTo
```go
func (r *Reader) WriteTo(w io.Writer) (n int64, err error)
```
WriteTo实现了io.WriterTo接口.


