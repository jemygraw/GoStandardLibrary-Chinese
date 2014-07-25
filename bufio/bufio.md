#bufio

import "bufio"

##简介

带buffer机制的I/O操作包.

##概览

bufio包实现了带buffer机制的I/O操作包.通过创建新的对象(Reader或者Writer)的方式对io.Reader或io.Writer进行封装,
实现了同样的接口,但是在内部提供了buffer机制,同时还支持一些实用的文本I/O接口.

##内容
###常量

```go
const (
        // Scanner中使用.每次扫描最多扫描的token的长度.实际读取到的长度可能比这个值小,
        // 因为缓冲区还需要包括分隔符,比如说换行符.
        MaxScanTokenSize = 64 * 1024
      )
```

```go
var (
        // bufio: 非法调用`UnreadByte`
        ErrInvalidUnreadByte = errors.New("bufio: invalid use of UnreadByte")
        // bufio: 非法调用`UnreadRune`
        ErrInvalidUnreadRune = errors.New("bufio: invalid use of UnreadRune")
        // bufio: 缓冲区满
        ErrBufferFull        = errors.New("bufio: buffer full")
        // bufio: 负计数
        ErrNegativeCount     = errors.New("bufio: negative count")
    )

var (
        // bufio.Scanner: token超长
        ErrTooLong         = errors.New("bufio.Scanner: token too long")
        // bufio.Scanner: 切割函数SplitFunc返回了负数
        ErrNegativeAdvance = errors.New("bufio.Scanner: SplitFunc returns negative advance count")
        // bufio.Scanner: 切割函数SplitFunc返回数超出了输入
        ErrAdvanceTooFar   = errors.New("bufio.Scanner: SplitFunc returns advance count beyond input")
    )
```

###Scanner
```go
type Scanner struct{}
```
Scanner提供了方便的接口用于读取数据,例如读取按行分隔的文本文件.成功调用`Scan`方法后会读取文件中的一个`token`,自动忽略了token之间的字节.token是由类型为SplitFunc的函数定义的;默认的切割函数根据终端换行符把输入切割成行.包里面定义的切割函数是用于读取文件后按照行,字节,UTF-8编码字符和空格分隔的单词进行分割.用户可以自定义实现切割函数.

扫描在遇到EOF,I/O错误或者token太大放不下缓冲区时会停止,这是无法恢复的.When a scan stops,the reader may have advanced arbitrarily far past the last token.一个程序如果需要多控制错误处理,处理大token或者在一个reader上顺序扫描,最好使用bufio.Reader替代.

###func NewScanner
```go
func NewScanner(r io.Reader) *Scanner
```
NewScanner返回从io.Reader读取数据的Scanner.默认分隔函数是ScanLines.


###func (*Scanner) Bytes
```go
func (s *Scanner) Bytes() []byte
```
Bytes返回最近一次调用Scan产生的token.这个函数**不可重入**,返回的数组指向的底层数据可能会在下一次调用Scan后被重写.Bytes函数调用不会发生内存分配操作.
  

###func (*Scanner) Err
```go
func (s *Scanner) Err() error
```
Err返回内部第一次出现的错误,不包括EOF


###func (*Scanner) Scan
```go
func (s *Scanner) Scan() bool
```
Scan扫描获取下一个token,扫描完成后可以通过`Bytes`或者`Text`接口获取.扫描至输入结束或者遇到错误的时候返回false.如果Scan返回了false,调用`Err`会返回扫描中遇到的错误,如果遇到的错误是EOF,`Err`接口会返回nil.


###func (*Scanner) Split
```go
func (s *Scanner) Split(split SplitFunc)
```
Split用于设置分隔函数.必须在调用`Scan`前设置.默认分隔函数是`ScanLines`.


###func (*Scanner) Text
```go
func (s *Scanner) Text() string
```
Text返回最近一次调用`Scan`生成的token,内部进行了一次内存分配,把bytes转换为string.


###type SplitFunc
```go
type SplitFunc func(data []byte, atEOF bool) (advance int, token []byte, err error)
```
SplitFunc是用于分隔输入数据的分隔函数的签名.输入参数是剩下未处理数据的初始子串和一个标识,atEof标识表示Reader是否还有更多的数据可以读取.函数返回值是在输入流中向前读取的字节数,下一个返回给用户的token和错误.如果数据还不足以组成一个完整的token,比如说读取行数据的时候没有换行符,SplitFunc可以返回(0,nil,nil),让Scanner去读取更多的数据到slice,然后使用在输入流中相同起点但是包括了更多数据的slice尝试进行一次分隔.

如果返回的error不为nil,扫描操作会停止,同时error会返回给用户.

除非atEOF为true,不然调用函数的时候不会传入空数组.如果atEOF为true,data数组也是可能不为空,里面存储着最后未处理的数据.


###func ScanBytes
```go
func ScanBytes(data []byte, atEOF bool) (advance int, token []byte, err error)
```
ScanBytes是一个分割函数,用于Scanner中把每一个byte分隔为一个token.

###func ScanLines
```go
func ScanLines(data []byte, atEOF bool) (advance int, token []byte, err error)
```
ScanLines是一个分割函数,用于Scanner中把每一行文本分隔为一个token,去除掉末尾的换行标识符.返回行可能为空.换行标识符是一个可选的回车符后面紧跟着一个换行符,用正则表达式来表示就是`\r?\n`.最后一个非空行也会返回即使最后没有换行符.


###func ScanRunes
```go
func ScanRunes(data []byte, atEOF bool) (advance int, token []byte, err error)
```
ScanRunes是一个分割函数,用于Scanner中把每一个UTF-8编码的字符分隔为一个token. 返回的字符序列相当于遍历整个输入作为字符串,这意味着错误的UTF-8编码将转换为 U+FFFD =  "\xef\xbf\xbd".因为`Scan`接口,这使得用户没有办法区分正确编码的替换字符和编码错误.


###func ScanWords
```go
func ScanWords(data []byte, atEOF bool) (advance int, token []byte, err error)
```
ScanWords是一个分割函数,用于Scanner中把每一个空格隔开的单词分隔为token,单词两边的空格会被自动删除.她不会返回任何空字符串.是否为空格是由unicode.isSpace定义的.


###ReadWriter
```go
type ReadWriter struct {
    *Reader
    *Writer
}
```
ReadWriter存储了Reader和Writer的指针,它实现了io.ReadWriter接口.

###func NewReadWriter
```go
func NewReadWriter(r *Reader, w *Writer) *ReadWriter
```
NewReadWriter分配一个新的ReadWriter.

###Reader
```go
type Reader struct {
}
```
Reader实现了带buffer机制的io.Reader.

###func NewReader
```go
func NewReader(rd io.Reader) *Reader
```
NewReader生成返回一个新的Reader,内部buffer为默认大小. 

###func NewReaderSize
```go
func NewReaderSize(rd io.Reader, size int) *Reader
```
NewReaderSize生成返回一个新的Reader,参数指定了内部buffer的最小大小.如果io.Reader参数本身就是一个Reader且空间足够大,它将直接返回底层的Reader.


###func (*Reader) Buffered
```go
func (b *Reader) Buffered() int
```
Buffered返回当前buffer中还能够读取的字节数.



###func (*Reader) Peek
```go
func (b *Reader) Peek(n int) ([]byte, error)
```
Peek返回后续的n个字节,但是不真正进行读取.返回的数组在下次调用读操作之前有效.如果Peek返回小于n各字节,那么会同时返回一个错误解释为什么读到数据不够.如果n比缓冲区还大则返回错误ErrBufferFull.


###func (*Reader) Read
```go
func (b *Reader) Read(p []byte) (n int, err error)
```
Read读取数据存储至p.它返回读取到的字节数.最多调用底层Reader的Read方法一次,因此n可能小于len(p).在遇到EOF的时候,返回io.EOF和0.


###func (*Reader) ReadByte
```go
func (b *Reader) ReadByte() (c byte, err error)
```
ReadByte读取返回一个字节.如果没有可读字节,返回错误.


###func (*Reader) ReadBytes
```go
func (b *Reader) ReadBytes(delim byte) (line []byte, err error)
```
ReadBytes一直读取直到第一次在输入中遇到delim,返回一个字节数组,包含了读取到的字节和分隔符delim.如果ReadBytes在找到分隔符之前出现了错误,返回遇到错误之前已经读取到的数据和错误本身(通常就是io.EOF).ReadString当且仅当在数据不以delim为结尾的情况下返回err != nil.有时候简单的情况下,使用Scanner会更加方便.


###func (*Reader) ReadLine
```go
func (b *Reader) ReadLine() (line []byte, isPrefix bool, err error)
```
ReadLine是一个底层的按行读取原始操作.大多数用户应该使用ReadBytes('\n)或者ReadString('\n')或者使用Scanner.

ReadLine尝试返回简单一行数据,不包括末尾字节.如果输入行对缓冲区来说太长,那么isPrefix会被设置为true,然后返回输入行的开头部分.输入行的剩下部分会在其它调用中返回.在返回输入行的最后一个片段时isPrefix会设置为false.函数是不可重入的,返回的缓冲区只在下次调用ReadLine前有效.ReadLine要不返回非空的行,要不返回一个错误,不会同时返回.

ReadLine返回的文本不包括行末结束符("\r\n"或者"\n").如果输入流的末尾没有结束符,也不会有提示或者错误发生.在ReadLine之后调用`UnreadByte`也会标记最后一个字节未读(这个字节可能是属于行末结束符的一个字节),即使这个字节不是ReadLine返回的数据中的一部分.


###func (*Reader) ReadRune
```go
func (b *Reader) ReadRune() (r rune, size int, err error)
```
ReadRune读取一个UTF-8编码的字符,返回字符以及其占用字节数.如果字符是无效的,它消耗一个字节然后返回unicode.ReplacementChar (U+FFFD)和大小1.


###func (*Reader) ReadSlice
```go
func (b *Reader) ReadSlice(delim byte) (line []byte, err error)
```
ReadSlice一直读取直到第一次在输入中遇到delim,返回一个slice指向缓冲区中的字节.再下次读取操作之时bytes将失效.如果ReadSlice在找到分隔符之前出现了错误,返回缓冲区中的所有数据和错误本身(通常就是io.EOF).如果缓冲区不存在delim,ReadSlice失败返回错误ErrBufferFull.因为ReadSlice返回的数据会被下次I/O操作重写,用户基本上应该使用ReadBytes或者ReadString.ReadSlice当且仅当数据不以delim为结尾的情况下返回err != nil.


###func (*Reader) ReadString
```go
func (b *Reader) ReadString(delim byte) (line string, err error)
```
ReadString一直读取直到第一次在输入中遇到delim,然后返回一个字符串line,line包含读取到的数据和分隔符.如果ReadString在找到分隔符之前出现了错误,返回遇到错误之前已经读取到的数据和错误本身(通常就是io.EOF).ReadString当且仅当在数据不以delim为结尾的情况下返回err != nil.一般情况下,使用Scanner会更加方便.


###func (*Reader) Reset
```go
func (b *Reader) Reset(r io.Reader)
```
Reset丢弃任何缓存的数据,清除所有状态,切换Reader从r读取数据.


###func (*Reader) UnreadByte
```go
func (b *Reader) UnreadByte() error
```
UnreadByte标记上一个读出的字节为未读状态.只有最近一个读出来的字节才可以标记.

###func (*Reader) UnreadRune
```go
func (b *Reader) UnreadRune() error
```
UnreadRune标记上一个读出的字符为未读状态.如果最近一次调用的读操作不是`ReadRune`,UnreadRune将返回一个错误.(在这方面,它比UnreadByte更加严格,UnreadByte对任何读操作都适用.)

 

###func (*Reader) WriteTo
```go
func (b *Reader) WriteTo(w io.Writer) (n int64, err error)
```
WriteTo实现了io.WriterTo接口.

###Writer
```go
type Writer struct {
}
```
Writer实现了带buffer机制的io.Writer.如果在写入过程出现了错误,不会再接受写入更多数据,且随后的所有写操作都会返回错误.在所有数据都写入结束后,用户应该调用`Flush`保证所有数据转发到底层的io.Writer.


###func NewWriter
```go
func NewWriter(w io.Writer) *Writer
```
NewWriter生成返回一个新的Writer,内部buffer为默认大小. 


###func NewWriterSize
```go
func NewWriterSize(w io.Writer, size int) *Writer
```
NewWriterSize生成返回一个新的Writer,参数指定了内部buffer的最小大小.如果io.Writer参数本身就是一个Writer且空间足够大,它将直接返回底层的Writer.


###func (*Writer) Available
```go
func (b *Writer) Available() int
```
Available返回缓冲区中还有多少未使用空间.

###func (*Writer) Buffered
```go
func (b *Writer) Buffered() int
```
Buffered返回已经写到当前缓冲区的数据的字节数.


###func (*Writer) Flush
```go
func (b *Writer) Flush() error
```
Flush把缓冲区的数据到写出到底层的io.Writer.

###func (*Writer) ReadFrom
```go
func (b *Writer) ReadFrom(r io.Reader) (n int64, err error)
```
 ReadFrom实现了io.ReaderFrom接口.

###func (*Writer) Reset
```go
func (b *Writer) Reset(w io.Writer)
```
Reset丢弃缓冲区上全部数据,清楚任何错误信息,重新设置b输出到w.


###func (*Writer) Write
```go
func (b *Writer) Write(p []byte) (nn int, err error)
```
Write把字节数组p写入缓冲区,返回写入的字节数.如果写入的字节数不够len(p),会同时返回一个错误表示原因.


###func (*Writer) WriteByte
```go
func (b *Writer) WriteByte(c byte) error
```
WriteByte写入一个字节.

###func (*Writer) WriteRune
```go
func (b *Writer) WriteRune(r rune) (size int, err error)
```
WriteRune写入一个字符,返回写入的字节数和错误.


###func (*Writer) WriteString
```go
func (b *Writer) WriteString(s string) (int, error)
```
WriteString写入一个字符串,返回写入的字节数.如果写入的字节数不够len(s),会同时返回一个错误表示原因.


