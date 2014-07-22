# archive/zip

import "archive/zip"

##简介

zip包提供了读写zip归档文件的功能。

##概述

zip包提供了读写zip归档文件的功能。  
查看：[http://www.pkware.com/documents/casestudies/APPNOTE.TXT](http://www.pkware.com/documents/casestudies/APPNOTE.TXT)  

该包不支持磁盘跨越[Disk Spanning](http://en.wikipedia.org/wiki/Disc_spanning)功能。

关于[ZIP64格式](http://en.wikipedia.org/wiki/Zip_(file_format))的说明：

为了向后兼容，FileHeader同时有32位和64位的域。64位的域总是会有正确的值，而且大多数情况下两者是相同的。对于需要ZIP64格式的归档文档来说，32位的域将被填充为0xffffffff，然后必须使用64位的域。

##内容

###常量
```go
const (
        Store   uint16 = 0
        Deflate uint16 = 8
)
```
压缩模式。

##变量
```go
var (
        // zip文件格式错误
        ErrFormat    = errors.New("zip: not a valid zip file")
        // 不支持的算法错误
        ErrAlgorithm = errors.New("zip: unsupported compression algorithm")
        // 校验和错误
        ErrChecksum  = errors.New("zip: checksum error")
)
```

###func RegisterCompressor
```go
func RegisterCompressor(method uint16, comp Compressor)
```
RegisterCompressor注册一个自定义的压缩算法，该算法被分配一个方法ID。常用的算法Store和Deflate是内置的。

###func RegisterDecompressor
```go
func RegisterDecompressor(method uint16, d Decompressor)
```
RegisterDecompressor可以为指定方法ID自定义一个解压算法。

###type Compressor
```go
type Compressor func(io.Writer) (io.WriteCloser, error)
```
Compressor返回一个向文件写入压缩数据的writer，该writer实现了io.WriteCloser接口，使用该writer将压缩后的数据写入所提供的参数io.Writer中。在关闭文件的时候，任何缓冲数据都将被写入底层文件中。

###type Decompressor
```go
type Decompressor func(io.Reader) io.ReadCloser
```
Decompressor返回一个从压缩文件读取数据的reader，该reader实现了io.ReadCloser接口，使用该reader可以从所提供的参数io.Reader中读取数据并解压。该reader将被返回给打开归档文件的调用者，这些调用者在数据读取完毕的时候会负责关闭这个reader。

###type File
```go
type File struct {
    FileHeader
    // 包含过滤掉或者未导出的域
}
```

###func (*File) DataOffset
```go
func (f *File) DataOffset() (offset int64, err error)
```
DataOffset()方法返回文件f相对于zip文件起始位置的可能被压缩的数据的偏移量。
大多数的调用都应该使用Open()函数，该函数可以透明地处理数据解压和校验和验证。

###func (*File) Open
```go
func (f *File) Open() (rc io.ReadCloser, err error)
```
Open()函数返回一个io.ReadCloser，提供对文件内容的访问。多个文件可以并发读取。

###type FileHeader
```go
type FileHeader struct {
    // Name 是这个文件的名称
    // 该名称必须是相对路径，不可以以盘符(比如C:)或者反斜杠(\)开头。可以允许正斜杠(/)存在。
    Name string

    CreatorVersion     uint16
    ReaderVersion      uint16
    Flags              uint16
    Method             uint16
    ModifiedTime       uint16 // MS-DOS 时间
    ModifiedDate       uint16 // MS-DOS 日期
    CRC32              uint32
    CompressedSize     uint32 // 已废弃; 使用 CompressedSize64
    UncompressedSize   uint32 // 已废弃; 使用 UncompressedSize64
    CompressedSize64   uint64
    UncompressedSize64 uint64
    Extra              []byte
    ExternalAttrs      uint32 // 参数含义依赖于CreatorVersion
    Comment            string
}
```
FileHeader用来描述zip文件中的一个文件。可以查看zip规范来获取详细信息。

###func FileInfoHeader
```go
func FileInfoHeader(fi os.FileInfo) (*FileHeader, error)
```
FileInfoHeader根据os.FileInfo创建一个部分域填充的FileHeader。因为os.FileInfo的Name()方法仅返回文件的短文件名，所以可能需要修改返回的FileHeader的Name域来提供文件的完整路径名。

###func (*FileHeader) FileInfo
```go
func (h *FileHeader) FileInfo() os.FileInfo
```
FileInfo()方法返回FileHeader的os.FileInfo信息。

###func (*FileHeader) ModTime
```go
func (h *FileHeader) ModTime() time.Time
```
ModTime()方法返回文件在UTC时区下的修改时间。分辨率是2秒。

###func (*FileHeader) Mode
```go
func (h *FileHeader) Mode() (mode os.FileMode)
```
Mode()方法返回FileHeader的访问权限和模式位。

###func (*FileHeader) SetModTime
```go
func (h *FileHeader) SetModTime(t time.Time)
```
SetModTime()方法用给定的UTC时间参数来设置域ModifiedTime和MidifiedDate。分辨率位2秒。

###func (*FileHeader) SetMode
```go
func (h *FileHeader) SetMode(mode os.FileMode)
```
SetMode()方法用来位FileHeader设置访问权限和模式位。

###type ReadCloser
```go
type ReadCloser struct {
    Reader
    // 包含过滤掉的或未导出的域
}
```

###func OpenReader
```go
func OpenReader(name string) (*ReadCloser, error)
```
OpenReader()函数打开参数name所指定的zip文件并返回一个ReadCloser。

###func (*ReadCloser) Close
```go
func (rc *ReadCloser) Close() error
```
关闭zip文件，标记为I/O不可用。

###type Reader
```go
type Reader struct {
    File    []*File
    Comment string
    // 包含过滤掉的或未导出的域
}
```

###func NewReader
```go
func NewReader(r io.ReaderAt, size int64) (*Reader, error)
```
NewReader函数返回一个新的Reader，该Reader从参数r里面读取数据，这个r里面假定有size个字节。

###type Writer
```go
type Writer struct {
    // 包含过滤掉的或未导出的域
}
```
Writer实现了一个写入zip文件的功能。

###func NewWriter
```go
func NewWriter(w io.Writer) *Writer
```
NewWriter()函数返回一个新的Writer，该Writer用来写入数据到zip文件。

###func (*Writer) Close
```go
func (w *Writer) Close() error
```
Close()方法通过写入中间目录来结束写入数据到zip文件。但它没有（也无法）关闭底层的writer。

###func (*Writer) Create
```go
func (w *Writer) Create(name string) (io.Writer, error)
```
Create()方法使用name参数来将一个文件添加到zip文件中。它返回一个应该向其写入数据的io.Writer。name参数必须是相对路径，它不可以以驱动器符号（例如C:）开头或者以反斜杠（\）开头，但是支持正斜杠（/）。该文件的内容必须写入io.Writer，然后你才能再次调用Create()，CreateHeader()或者Close()方法。

###func (*Writer) CreateHeader
```go
func (w *Writer) CreateHeader(fh *FileHeader) (io.Writer, error)
```
CreateHeader()方法使用表示文件元数据的FileHeader参数来将一个文件添加到zip文件中。它返回一个应该向其写入数据的io.Writer。该文件的内容必须写入io.Writer，然后你才可以再次调用Create()，CreateHeader()或者Close()方法。
