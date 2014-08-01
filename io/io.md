
#io包

import "io"

----------

##简介

##概览
io包提供了底层I/O操作的接口。它主要的工作是封装现有的底层实现的方法，比如os包中的方法，然后提供抽象出来的方法作为接口。

因为这些接口和操作用各种实现封装了底层，除非另有通知，多个实例不应该假设他们在并行操作时是安全的。

##变量
```go
var EOF = errors.New("EOF")
```
EOF是当没有数据输入时返回的读入错误。函数应该通过返回EOF来给出一个信号，优雅地指示输入结束。如果EOF意外的发生在一个有结构的数据流中，会产生ErrUnexpectedEOF或者其他能给出更加详细信息的错误，这些错误都是相似的。

```go
var ErrClosedPipe = errors.New("io: read/write on closed pipe")
```
ErrClosedPipe是用来指明操作一个已经关闭的管道时产生的错误。
```go
var ErrNoProgress = errors.New("multiple Read calls return no data or error")
```
当多次调用Read，没有数据返回或者读取错误，多个io.Reader的实例会返回ErrNoProgress。ErrNoProgress经常作为一个损坏的io.Reader的标记。
```go
var ErrShortBuffer = errors.New("short buffer")
```
ErrShortBuffer说明读取需要的buffer比提供的更长。
```go
var ErrShortWrite = errors.New("short write")
```
ErrShortWrite意味着一次写入得到的bytes比需要的少，但是无法返回一个显示的错误。
```go
var ErrUnexpectedEOF = errors.New("unexpected EOF")
```
ErrUnexpectedEOF意味着在读取一个固定大小的块或者数据结构时遇到了EOF。

###func Copy
```go
func Copy(dst Writer, src Reader) (written int64, err error)
```
从`src`到`dst`复制数据，直到遇到EOF或者产生错误为止。返回已经复制的数据和在复制过程中遇到的第一个错误，如果错误存在的话。

一次成功的复制返回err==nil，而不是err==EOF。因为复制定义为从`src`读取直到EOF，它不会把EOF当作错误对待。

如果`src`实现了WriterTo接口，复制就通过调用src.WriteTo(dst)实现。否则，如果`dst`实现了ReaderFrom接口，复制通过调用dst.ReadFrom(src)实现。

>可能的异常： io.ErrShortWrite:写入数据不等于读取数据。

###func CopyN
```go
func CopyN(dst Writer, src Reader, n int64) (written int64, err error)
```
从`src`向`dst`拷贝`n`字节数据直到遇到错误。返回已经考不的字节数和拷贝过程中最近的一次错误。返回时只有err==nil，写入的字节数才为`n`。

如果`dst`实现了ReaderFrom接口，拷贝就通过它实现。

###func ReadAtLeast
```go
func ReadAtLeast(r Reader, buf []byte, min int) (n int, err error)
```
从`r`向`buf`读取数据直到已经读取了至少`min`字节的数据。返回已经拷贝的数据的字节数和错误（如果读取的字节数不足）。只有没有任何数据读取时error才会是EOF。如果EOF在读取数据少于`min`字节时发生，则返回ErrUnexpectedEOF。如果min比buf的长，返回ErrShortBuffer错误。只有err==nil，在返回时n>=min才成立。

###func ReadFull
```go
func ReadFull(r Reader, buf []byte) (n int, err error)
```
从`r`到`buf`精确地读取len(buf)字节的数据。返回已经拷贝的字节数和错误（如果读取的字节不足）。只有没有数据读取时才会返回EOF。如果在读取部分数据时发生了EOF错误，则返回ErrUnexpectedEOF。在返回时，只有err==nil，n==len(buf)才成立。

###func WriteString
```go
func WriteString(w Writer, s string) (n int, err error)
```
将字符串`s`的数据写入`w`，`w`接收一个字节数组。如果`w`已经实现了WriteString方法，将会直接被调用。

###type ByteReader interface
```go
type ByteReader interface {
    ReadByte() (c byte, err error)
}
```
ByteReader是封装了ReadByte方法的接口。ReadByte 从输入读取并返回下一个字节。如果没有可用的字节，将产生错误`err`。

###type ByteScanner interface
```go
type ByteScanner interface {
    ByteReader
    UnreadByte() error
}
```
ByteScanner是一个接口，它将UnreadByte 方法和基本的ReadByte组合。UnreadByte产生下一次的ReadByte调用，来返回和上一次ReadByte调用相同数目的字节。连续调用两次UnreadByte，中间如果没有调用ReadByte，将会产生错误。

###type ByteWriter interface
```go
type ByteWriter interface {
    WriteByte(c byte) error
}
```
ByteWriter是封装了WriteByte方法的接口。

###type Closer interface
```go
type Closer interface {
    Close() error
}
```
Closer是接口，封装了基本的Close方法。在第一次调用Close后的行为是为定义的。在具体的实现中可能会记录自己的行为。

###type LimitedReader struct
```go
type LimitedReader struct {
    R Reader // underlying reader
    N int64  // max bytes remaining
}
```
LimitedReader 从`R`读取数据担心限制了返回的数据为`N`字节。每次调用读取都会更新`N`来反映新的剩余的数量。

###func (*LimitedReader) Read
```go
func (l *LimitedReader) Read(p []byte) (n int, err error)
```

###type PipeReader struct
```go
type PipeReader struct {
    // contains filtered or unexported fields
}
```
PipeReader是一个只读的单工管道。

###func Pipe
```go
func Pipe() (*PipeReader, *PipeWriter)
```
创建了在内存的一个同步的管道。它可以用来连接io.Reader和io.Writer。一端的读取和另一端的写入匹配，在两者之间直接复制数据，没有内部的缓冲。互相调用读取和写入（带有Close也可以）是安全的。Close会在即将完成的I/O完成后结束。并行地调用读取，并行地调用写入都是安全的：独立的调用都是顺序封闭的。

###func (*PipeReader) Close
```go
func (r *PipeReader) Close() error
```
关闭reader，随后向只写的单工管道写入数据将会导致ErrClosedPipe错误。

###func (*PipeReader) CloseWithError
```go
func (r *PipeReader) CloseWithError(err error) error
```
关闭reader，随后向只写的单工管道写入数据将会返回err错误。

###func (*PipeReader) Read
```go
func (r *PipeReader) Read(data []byte) (n int, err error)
```
实现了标准Read接口。从管道读取数据，直到一次写入到达或者写入端被关闭才会产生阻塞。如果写入端被关闭时产生了`err`错误，错误将会返回，否则错误是EOF。

###type PipeWriter struct
```go
type PipeWriter struct {
    // contains filtered or unexported fields
}
```
PipeWriter是一个只写的单工管道。
###func (*PipeWriter) Close
```go
func (w *PipeWriter) Close() error
```
关闭Writer，随后另一端从只读的单工管道读将不会返回数据，并返回EOF。
###func (*PipeWriter) CloseWithError
```go
func (w *PipeWriter) CloseWithError(err error) error
```
关闭Writer，随后另一端从只读的单工管道读将不会返回数据，并返回`err`错误。

###func (*PipeWriter) Write
```go
func (w *PipeWriter) Write(data []byte) (n int, err error)
```
实现了标准Write接口，它向管道写入数据，直到readers已经取完所有数据或者读取的另一端关闭才会阻塞。如果读取端被关闭并有错误，`err`将会返回，否则`err`为ErrClosedPipe。

###type ReadCloser interface
```go
type ReadCloser interface {
    Reader
    Closer
}
```
ReadCloser是将基本的Read和Close方法组合的接口。

###type ReadSeeker interface
```go
type ReadSeeker interface {
    Reader
    Seeker
}
```
ReadSeeker 是将基本的Read和Seek方法组合的接口。

###type ReadWriteCloser interface
```go
type ReadWriteCloser interface {
    Reader
    Writer
    Closer
}
```
ReadWriteCloser是将基本的Read、Write和Close方法组合的接口。

###type ReadWriteSeeker interface
```go
type ReadWriteSeeker interface {
    Reader
    Writer
    Seeker
}
```
ReadWriteSeeker是将基本的Read、Write和Seek方法组合的接口。

###type ReadWriter interface
```go
type ReadWriter interface {
    Reader
    Writer
}
```
ReadWriter是将基本的Read和Write方法组合的接口。

###type Reader interface
```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
```
Reader是封装了基本Read方法的接口。

Read将len(p)字节的数据读入`p`。它返回读取的字节数(0<=n<=len(p))和遇到的错误。即使Read返回的n<len(p)，在调用时它也会使用所有的`p`作为写入空间。如果数据是可用的，但不是len(p)字节，Read按照惯例返回可用的数据，而不是等待得到更多的数据。

如果Read在成功读取n>0字节的数据后遇到了错误或者文件的结尾，它返回读到的字节数。在相同的调用它可能返回非空错误或者在接下来的调用返回错误（并且n==0）。这种情形的一个通用的实例是一个在输入流的末尾返回非零的字节数Reader可能返回err==EOF或者err==nil。下一次的Read应该返回0和EOF，不管怎样。

调用函数应该先处理返回的n>0的字节数据，再考虑err错误。这样做可以正确地处理在读取一些字节后的I/O错误和所有允许的EOF情形。

实现Read方法时不应该返回零字节和nil错误，调用函数应该将这种情形作为no-op（未操作）。

###func LimitReader
```go
func LimitReader(r Reader, n int64) Reader
```
返回一个Reader，它从`r`读取但是在读到EOF或者读取`n`字节数据后停止。底层的实现是一个*LimitReader。

###func MultiReader
```go
func MultiReader(readers ...Reader) Reader
```
返回一个Reader，将输入的`readers`逻辑上串联在一起。它们依次读取。一旦所有的输入都返回了EOF，Read会返回EOF。如果其中任意一个reader返回一个非空的、非EOF的错误，Read将会返回那个错误。

###func TeeReader
```go
func TeeReader(r Reader, w Writer) Reader
```
返回一个Reader，它将从`r`读到的数据写入`w`。所有和这个Reader的相关的`r`的reads方法都有`w`的writes对应。没有内部的缓冲，因此write必须在read完成之前完成。任何的在写入时遇到的错误都会视为read 错误。

###type ReaderAt interface
```go
type ReaderAt interface {
    ReadAt(p []byte, off int64) (n int, err error)
}
```
ReaderAt接口封装了基本的ReadAt方法。

ReadAt从偏移量为`off`将len(p)字节的底层数据输入读入`p`。它返回读到的字节数（0<=n<=len(p)）并返回遇到的错误。

当ReadAt返回n<len(p)，它返回一个非空错误解释没有返回更多字节的原因。在这个方面，ReadAt比Read更苛刻。

尽管ReadAt返回n<len(p)，在调用期间它可能使用所有的p作为读写空间。如果一些数据是可以获得的但是不是len(p)字节，ReadAt阻塞直到所有数据都可以获得或者返回一个错误。在这个方面，ReadAt和Read不同。

如果ReadAt返回的n=len(p)在输入数据的结尾，ReadAt可能返回err==EOF或者err==nil。

如果ReadAt带有一个seek偏移量从源数据读入，ReadAt不应该影响或者被底层seek偏移量影响。

多个ReadAt的实例可以并行读取同一个源文件。

###type ReaderFrom interface
```go
type ReaderFrom interface {
    ReadFrom(r Reader) (n int64, err error)
}
```
ReaderFrom是一个封装了ReadFrom方法的接口。

ReadFrom从`r`读取数据直到EOF或者遇到错误。返回值`n`是读到的字节数。在读取时任何除了io.EOF之外的错误都会返回。

如果有可用的ReadFrom函数，Copy函数使用它。

###type RuneReader interface
```go
type RuneReader interface {
    ReadRune() (r rune, size int, err error)
}
```
 RuneReader是封装了ReadRune方法的接口。
 ReadRune读取单个UTF-8编码的字符并返回rune类型和她得字节数。如果没有可用的字符，将会返回错误`err`。
 

###type RuneScanner interface
```go
type RuneScanner interface {
    RuneReader
    UnreadRune() error
}
```
RuneScanner是将UnreadRune方法加入到基本的ReadRune方法组成的接口。

UnreadRune引起下一次的ReadRune调用，并返回上一次ReadRune相同的rune类型。调用两次UnreadRune时，如果中间没有调用ReadRune，那么将会返回错误。

###type SectionReader struct
```go
type SectionReader struct {
    // contains filtered or unexported fields
}
```
SectionReader 实现了Read、Seek和ReadAt，是底层ReaderAt接口上的一部分。

>SectionReader 类型SectionReader是一个struct（没有任何导出的字段），实现了 Read, Seek 和 ReadAt，同时，内嵌了 ReaderAt 接口。

###func NewSectionReader
```go
func NewSectionReader(r ReaderAt, off int64, n int64) *SectionReader
```
NewSectionReader返回一个SectionReader，它从`r`读取，偏移量为`off`，并在读取`n`字节数据后的EOF处停止。

###func (*SectionReader) Read
```go
func (s *SectionReader) Read(p []byte) (n int, err error)
```

###func (*SectionReader) ReadAt
```go
func (s *SectionReader) ReadAt(p []byte, off int64) (n int, err error)
```

###func (*SectionReader) Seek
```go
func (s *SectionReader) Seek(offset int64, whence int) (int64, error)
```

###func (*SectionReader) Size
```go
func (s *SectionReader) Size() int64
```

###type Seeker interface
```go
type Seeker interface {
    Seek(offset int64, whence int) (int64, error)
}
```
Seeker是一个接口，它封装了基本的Seek方法。

Seek根据`whence`设置了下一次读或写的偏移量`offset`：0意味着相对于文件的开始，1意味着相对于当前偏移量，2意味着相对于结尾。Seek返回新的偏移量或者错误（如果存在的话）。

###type WriteCloser interface
```go
type WriteCloser interface {
    Writer
    Closer
}
```
WriteCloser是一个接口，它将基本的Write和Close方法组合。

###type WriteSeeker interface
```go
type WriteSeeker interface {
    Writer
    Seeker
}
```
WriteSeeker是一个接口，它将基本的Write和Seek方法组合。

###type Writer interface
```go
type Writer interface {
    Write(p []byte) (n int, err error)
}
```
Writer封装了基本的Write方法。

Write将len(p)字节的数据从`p`写入底层数据流。它返回从`p`写入的字节数（0<=n<=len(p)），并且如果过早地结束了写数据，同时会返回错误。如果n<len(p)，Write必须返回非空的错误。Write不能改变数据切片`p`，临时改变也不行。

###func MultiWriter
```go
func MultiWriter(writers ...Writer) Writer
```
MultiWriter创建了一个writer，它将writes复制给了所有给定的writers，类似于Unix tee(1)命令。

###type WriterAt interface
```go
type WriterAt interface {
    WriteAt(p []byte, off int64) (n int, err error)
}
```
WriterAt封装了基本的WriteAt方法。

WriteAt将len(p)字节的偏移量为`off`数据从`p`写入底层数据流。它返回的字节数（0<=n<=len(p)）并且返回导致write过早停止的错误。如果返回n<len(p)，WriteAt必须返回非空错误。

如果WriteAt带有一个seek偏移量向目标写入数据，WriteAt不应该影响或者被底层seek偏移量影响。

多个WriteAt的实例可以在同一个目标上执行并行的WriteAt调用，如果读取范围不重叠。

###type WriterTo interface
```go
type WriterTo interface {
    WriteTo(w Writer) (n int64, err error)
}
```
WriteTo是一个接口，它们封装了WriteTo方法。
WriteTo向`w`写入直到没有数据写入或者遇到错误。返回的值n是写入的字节数。在写入期间遇到的任何错误都会返回。

如果有可用的WriterTo方法，Copy函数会使用它。