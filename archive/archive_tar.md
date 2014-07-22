# archive/tar

import "archive/tar"

##简介

tar包实现了访问tar归档文件的方法。

##概览

tar包实现了访问tar归档文件的方法。该包目标覆盖各种变种，包括GNU和BSD的tar归档文件。

参考：

[http://www.freebsd.org/cgi/man.cgi?query=tar&sektion=5](http://www.freebsd.org/cgi/man.cgi?query=tar&sektion=5)  
[http://www.gnu.org/software/tar/manual/html_node/Standard.html](http://www.gnu.org/software/tar/manual/html_node/Standard.html)  
[http://pubs.opengroup.org/onlinepubs/9699919799/utilities/pax.html](http://pubs.opengroup.org/onlinepubs/9699919799/utilities/pax.html)  

##内容

###常量

```go
const (
        // 文件类型
        TypeReg           = '0'    // 常规文件
        TypeRegA          = '\x00' // 常规文件
        TypeLink          = '1'    // 硬链接
        TypeSymlink       = '2'    // 符号链接
        TypeChar          = '3'    // 字符设备节点
        TypeBlock         = '4'    // 块设备节点
        TypeDir           = '5'    // 目录
        TypeFifo          = '6'    // 先入先出节点
        TypeCont          = '7'    // 保留
        TypeXHeader       = 'x'    // 扩展头部
        TypeXGlobalHeader = 'g'    // 全局扩展头部
        TypeGNULongName   = 'L'    // 下一个文件文件名很长
        TypeGNULongLink   = 'K'    // 下一个符号链接链接到的文件名称很长
        TypeGNUSparse     = 'S'    // 稀疏文件
)
```

###变量
```go
var (
        // 写入内容太长错误
        ErrWriteTooLong    = errors.New("archive/tar: write too long")
        // 头部内容太长错误
        ErrFieldTooLong    = errors.New("archive/tar: header field too long")
        // 在关闭文件后写入错误
        ErrWriteAfterClose = errors.New("archive/tar: write after close")
)

var (
        // 非法头部错误
        ErrHeader = errors.New("archive/tar: invalid tar header")
)
```

###type Header
```go
type Header struct {
        Name       string    // 头部名称，一般设置为文件名全路径
        Mode       int64     // 权限和模式位
        Uid        int       // 用户id
        Gid        int       // 用户组id
        Size       int64     // 按字节表示长度
        ModTime    time.Time // 修改时间
        Typeflag   byte      // 头部条目类型
        Linkname   string    // 链接的目标名称
        Uname      string    // 用户名
        Gname      string    // 用户组名
        Devmajor   int64     // 字符或块主设备号
        Devminor   int64     // 字符或块次设备号
        AccessTime time.Time // 访问时间
        ChangeTime time.Time // 状态改变时间
        Xattrs     map[string]string
}
```
头部表示一个tar归档文件的一个头部信息。头部信息的一些域可以不填充数据。

###func FileInfoHeader
```go
func FileInfoHeader(fi os.FileInfo, link string) (*Header, error)
```
FileInfoHeader根据fi创建一个域部分填充的头部。如果fi表示一个符号链接的话，FileInfoHeader就把链接当作链接目标。如果fi是一个目录的话，文件名称会被追加一个斜杠（/）。因为os.FileInfo的方法Name()返回的是文件的短文件名，而不是全路径，所以可能需要修改返回的tar头部的Name域，以提供一个文件的全路径。

###func (*Header) FileInfo
```go
func (h *Header) FileInfo() os.FileInfo
```
FileInfo 返回一个tar头部的os.FileInfo信息。

###type Reader
```go
type Reader struct {
        // 包含过滤掉的或未导出的域
}
```
Reader提供了对一个tar归档文件内容的顺序访问。一个tar归档文件由一系列文件组成。Next()方法指向归档文件中的每个文件（包括第一个文件）的开始处，然后就可以使用io.Reader来访问文件的数据。

###func NewReader
```go
func NewReader(r io.Reader) *Reader
```
NewReader()方法从一个io.Reader创建一个新的tar的Reader。

###func (*Reader) Next
```go
func (tr *Reader) Next() (*Header, error)
```
Next()方法指向tar归档文件中的下一个文件（包括第一个文件）的开始处。

###func (*Reader) Read
```go
func (tr *Reader) Read(b []byte) (n int, err error)
```
Read()方法从当前指向的tar归档文件中的文件的开始处读取数据。当读取到当前文件的末尾时，它返回0和io.EOF。当Next()再被调用时，重新从下一个文件的开始处读取数据。

###type Writer
```go
type Writer struct {
        // 包含过滤掉的或未导出的域
}
```
Writer提供了对tar归档文件（POSIX.1格式）内容的顺序写入。tar归档文件由一系列文件组成。调用WriteHeader来开始创建一个新文件，然后调用Write方法来将数据写入文件中，一共可以写入最多hdr.Size个字节。

###func NewWriter
```go
func NewWriter(w io.Writer) *Writer
```
NewWriter()方法创建一个向io.Writer写入数据的tar的Writer。

###func (*Writer) Close
```go
func (tw *Writer) Close() error
```
Close()方法关闭tar归档文件，将所有没有写入到底层writer的数据写入。

###func (*Writer) Flush
```go
func (tw *Writer) Flush() error
```
Flush()方法结束写入数据到当前文件（可选）。

###func (*Writer) Write
```go
func (tw *Writer) Write(b []byte) (n int, err error)
```
Writer()方法向tar归档文件中当前指向的文件写入数据。如果在调用WriteHeader()方法后写入文件的字节数大于hdr.Size的时候返回错误ErrWriteTooLong。

###func (*Writer) WriteHeader
```go
func (tw *Writer) WriteHeader(hdr *Header) error
```
WriterHeader()方法写入tar头部hdr，然后准备接收文件的内容。WriterHeeader()方法会调用Flush()方法，如果这不是第一个头部的话。在tar文件关闭之后调用该方法会返回ErrWriteAfterClose错误。
