#http包

import "net/http"

----------

##简介

##概览
http包提供了HTTP的客户端和服务端实现。
Get, Head, Post, and PostForm构造了HTTP或者HTTPS的请求。

```go
resp, err := http.Get("http://example.com/")
...
resp, err := http.Post("http://example.com/upload", "image/jpeg", &buf)
...
resp, err := http.PostForm("http://example.com/form",
	url.Values{"key": {"Value"}, "id": {"123"}})
```

当完成响应body后，客户端必须关闭它。
```go
resp, err := http.Get("http://example.com/")
if err != nil {
	// handle error
}
defer resp.Body.Close()
body, err := ioutil.ReadAll(resp.Body)
// ...
```

为了控制HTTP客户端头、重定向策略和其他设置，创建一个Client：
```go
client := &http.Client{
	CheckRedirect: redirectPolicyFunc,
}

resp, err := client.Get("http://example.com")
// ...

req, err := http.NewRequest("GET", "http://example.com", nil)
// ...
req.Header.Add("If-None-Match", `W/"wyzzy"`)
resp, err := client.Do(req)
// ...
```
为了控制代理、TLS配置、keep-alives、压缩和其他设置，创建一个Transport：
```go
tr := &http.Transport{
	TLSClientConfig:    &tls.Config{RootCAs: pool},
	DisableCompression: true,
}
client := &http.Client{Transport: tr}
resp, err := client.Get("https://example.com")
```
对于多个goroutines的并发，Clients和Transport是安全的。为了获得高效率，应该被创建一次然后重用。

ListenAndServe用给定的地址和handler开启一个HTTP服务器。handler经常是空的，那意味着使用了DefaultServeMux。Handle和HandleFunc将handlers加入了DefaultServeMux：
```go
http.Handle("/foo", fooHandler)
http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
})
log.Fatal(http.ListenAndServe(":8080", nil))
```
创建一个自定义的Server可以更好的控制服务器的行为：
```go
s := &http.Server{
	Addr:           ":8080",
	Handler:        myHandler,
	ReadTimeout:    10 * time.Second,
	WriteTimeout:   10 * time.Second,
	MaxHeaderBytes: 1 << 20,
}
log.Fatal(s.ListenAndServe())
```

##变量

```go
var (
    ErrHeaderTooLong        = &ProtocolError{"header too long"}
    ErrShortBody            = &ProtocolError{"entity body too short"}
    ErrNotSupported         = &ProtocolError{"feature not supported"}
    ErrUnexpectedTrailer    = &ProtocolError{"trailer header without chunked transfer encoding"}
    ErrMissingContentLength = &ProtocolError{"missing ContentLength in HEAD response"}
    ErrNotMultipart         = &ProtocolError{"request Content-Type isn't multipart/form-data"}
    ErrMissingBoundary      = &ProtocolError{"no multipart boundary param in Content-Type"}
)
```

```go
var (
    ErrWriteAfterFlush = errors.New("Conn.Write called after Flush")
    ErrBodyNotAllowed  = errors.New("http: request method or response status code does not allow body")
    ErrHijacked        = errors.New("Conn has been hijacked")
    ErrContentLength   = errors.New("Conn.Write wrote more than the declared Content-Length")
)
```
以上是由HTTP server 引发的错误。

```go
var DefaultClient = &Client{}
```
DefaultClient是默认的Client，可以发起Get、Head和Post请求。

```go
var DefaultServeMux = NewServeMux()
```
DefaultServeMux是Server所用的默认的ServerMux。

```go
var ErrBodyReadAfterClose = errors.New("http: invalid Read on closed Body")
```
body被关闭之后读取Request和Response，返回ErrBodyReadAfterClose 。HTTP Handler在它的ResponseWriter上调用WriterHeader或者Write 方法之后，如果此时读取body，将会引起这种错误，这是一种很典型的情形。

```go
var ErrHandlerTimeout = errors.New("http: Handler timeout")
```
当ResponseWriter的Write在已经超时的handlers中调用，返回ErrHandlerTimeout。

```go
var ErrLineTooLong = errors.New("header line too long")
```
```go
var ErrMissingFile = errors.New("http: no such file")
```
当提供的文件的field name不知request中或者不是文件field，FromFile会返回ErrMissingFile。

```go
var ErrNoCookie = errors.New("http: named cookie not present")
```
```go
var ErrNoLocation = errors.New("http: no Location header in response")
```

###func CanonicalHeaderKey
```go
func CanonicalHeaderKey(s string) string
```
返回header key `s`的正规格式。它将首字母和hyphen后面的任意字母转为大写形式，其余的转为小写。比如`"accept-encoding"`的canonical key是` "Accept-Encoding"`。

###func DetectContentType
```go
func DetectContentType(data []byte) string
```
实现了`http://mimesniff.spec.whatwg.org`描述的算法，用来确定给定数据的内容类型。它考虑了最多512字节的数据。DetectContentType常常返回一个有效的MIME类型：如果它无法确定一个具体的类型，它返回` "application/octet-stream"`。

###func Error
```go
func Error(w ResponseWriter, error string, code int)
```
Error向请求回答带有具体错误信息的request和HTTP code。错误信息应该是纯文本格式。

###func Handle
```go
func Handle(pattern string, handler Handler)
```
Handle在DefaultServeMux注册了给定形式（pattern）的handler。ServeMux 的相关文档解释了形式（patterns ）如何被解析。

###func HandleFunc
```go
func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
```
HandleFunc在DefaultServeMux注册了给定形式（pattern）的handler 函数。ServeMux 的相关文档解释了形式（patterns ）如何被解析。

###func ListenAndServe
```go
func ListenAndServe(addr string, handler Handler) error
```
ListenAndServe在TCP网络地址上监听，然后带着handler来处理来到的链接并调用Serve。Handler典型值为nil，在这种情况下使用DefaultServerMux。
以下是一个小小的server：
```go
package main

import (
	"io"
	"net/http"
	"log"
)

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func main() {
	http.HandleFunc("/hello", HelloServer)
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
```

###func ListenAndServeTLS
```go
func ListenAndServeTLS(addr string, certFile string, keyFile string, handler Handler) error
```
ListenAndServeTLS的行为与ListenAndServe完全相同，只不过它等待的是HTTPS连接。而且，带有证书和匹配的private key的文件必须提供给server。如果证书被证书颁发机构签署，在server的证书后面是CA的证书，然后应是`certFile`。
```go
import (
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is an example server.\n"))
}

func main() {
	http.HandleFunc("/", handler)
	log.Printf("About to listen on 10443. Go to https://127.0.0.1:10443/")
	err := http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", nil)
	if err != nil {
		log.Fatal(err)
	}
}
```
可以使用 crypto/tls包中的generate_cert.go来生成cert.pem 和key.pem。

###func MaxBytesReader
```go
func MaxBytesReader(w ResponseWriter, r io.ReadCloser, n int64) io.ReadCloser
```
MaxBytesReader类似于LimitReader ，但是它的目的是限制到来的request bodies。与io.LimitReader相比，MaxBytesReader的结果是一个io.ReadCloser，它对于Read超过了limit返回non-EOF错误，并且当它的Close方法被调用的时候会关闭底层的reader。

###func NotFound
```go
func NotFound(w ResponseWriter, r *Request)
```
NotFound向请求回答HTTP 404 not found错误。

###func ParseHTTPVersion
```go
func ParseHTTPVersion(vers string) (major, minor int, ok bool)
```
解析HTTP版的字符串。"HTTP/1.0" 返回 (1, 0, true).

###func ParseTime
```go
func ParseTime(text string) (t time.Time, err error)
```
解析time header (比如 the Date: header)，它会尝试每一种HTTP/1.1允许的格式: TimeFormat, time.RFC850, and time.ANSIC。

###func ProxyFromEnvironment
```go
func ProxyFromEnvironment(req *Request) (*url.URL, error)
```
ProxyFromEnvironment返回给定request的代理url。一般该URL由用户的环境变量 $HTTP_PROXY and $NO_PROXY （or $http_proxy and $no_proxy）指定。如果用户的全局代理环境无效则返回一个错误。 如果全局环境变量没有定义或者，则会返回一个nil的URL和一个nil的错误。

一种特殊的情形，如果req.URL.Host是"localhost"（带有或者不带有端口号），会返回一个nil的URL和一个nil的错误。

###func ProxyURL
```go
func ProxyURL(fixedURL *url.URL) func(*Request) (*url.URL, error)
```
返回一个代理函数（在Transport中使用），它常常返回相同的URL。

###func Redirect
```go
func Redirect(w ResponseWriter, r *Request, urlStr string, code int)
```
Redirect向请求回答一个url重定向，它可能是一个相对于请求路径的路径。

###func Serve
```go
func Serve(l net.Listener, handler Handler) error
```
Serve在Listener`l`上接收HTTP连接，为每一个创建一个新的service goroutine。service goroutine读取请求然后调用handler来回答它们。Handler典型值为nil，在这种情况下使用DefaultServerMux。

###func ServeContent
```go
func ServeContent(w ResponseWriter, req *Request, name string, modtime time.Time, content io.ReadSeeker)
```
ServeContent使用提供的ReadSeeker中的content来回答request。ServeContent比io.Copy更好的地方主要是它恰当地处理Range request、设置MIME类型和处理 If-Modified-Since请求。

如果响应的内容类型头没有设置,该函数首先会尝试从文件的文件扩展名推断文件类型。如果推断不出来，则会读取文件的第一个块并传送给DetectContentType来检测类型。 文件名称也可以不使用。 如果文字名称为空，则服务器不会传送给响应。

如果修改时间不为0，ServeContent会把它放在服务器响应的Last-Modified头里面。如果客户端请求中包含了If-Modified-Since头，ServeContent会使用modtime来判断是否把内容传给客户端。

content的Seek方法必须能够工作。 ServeContent通过定位到文件结尾来确定文件大小。 

如果调用函数已经设置w的ETag 头，ServeContent使用它来处理使用 If-Range 和 If-None-Match的请求。

*os.File中实现了io.ReadSeeker接口。

###func ServeFile
```go
func ServeFile(w ResponseWriter, r *Request, name string)
```
ServeFile向请求输出带有名字的文件或者目录。

###func SetCookie
```go
func SetCookie(w ResponseWriter, cookie *Cookie)
```
SetCookie向给定的 ResponseWriter的headers增加Set-Cookie头。

###func StatusText
```go
func StatusText(code int) string
```
StatusText返回对应于HTTP 状态码对应的文字。如果code未知，则返回空字符串。

###type Client struct
```go
type Client struct {
    Transport RoundTripper
    CheckRedirect func(req *Request, via []*Request) error
    Jar CookieJar   
    Timeout time.Duration
}
```
Client是HTTP client。它的零值（DefaultClient）是一个使用DefaultTransport的有用的client。

Client的Transport 典型地有内部状态（缓存的TCP连接），所以Clients应该被重用而不是因需创建。Clients被多个goroutines使用时是并发安全的。

相对于RoundTripper（比如Transport），Client是高层次的，并且可以额外地处理HTTP的细节，比如cookies和redirects。

###func (*Client) Do
```go
func (c *Client) Do(req *Request) (resp *Response, err error)
```
Do发送一个HTTP request并返回HTTP响应，在客户端配置的策略（比如redirects、cookies和auth）之后。

如果由于客户端策略（比如CheckRedirect）引起了错误或者出现HTTP 协议错误，这个错误会返回。non-2xx响应不会引起错误。

当err是空的，resp常常包含一个非空的resp体。

当完成读取resp.Body之后，调用者应该关闭它。如果resp.Body没有关闭，Client的底层RoundTripper（典型是Transport）可能不会重用一个持续的向服务器的TCP连接来回答接下来的“keep-alive”请求。

请求体如果非空，会被底层Transport关闭，即使遇到错误。

通常，使用Get、Post或者PostFrom而不是Do。

###func (*Client) Get
```go
func (c *Client) Get(url string) (resp *Response, err error)
```
Get分发指定URL的GET请求。如果响应是以下重定向码，Get跟在重定向后面，最多10个重定向：
```go
301 (Moved Permanently)
302 (Found)
303 (See Other)
307 (Temporary Redirect)
```
如果有太多的重定向或者有HTTP 协议错误，将会返回一个错误。一个non-2xx响应不会引起错误。

当`err`为空，`resp`常常包含一个non-nil resp.Body。当完成了从resp.Body读数据，调用者应该关闭它。

###func (*Client) Head
```go
func (c *Client) Head(url string) (resp *Response, err error)
```
Head向指定的URL分发HEAD。如果响应是以下的重定向码，Head调用客户端的CheckRedirect函数后跟随重定向：
```go
301 (Moved Permanently)
302 (Found)
303 (See Other)
307 (Temporary Redirect)
```

###func (*Client) Post
```go
func (c *Client) Post(url string, bodyType string, body io.Reader) (resp *Response, err error)
```
Post分发指定的URL的POST请求。当完成了从resp.Body读数据，调用者应该关闭它。

如果给定的body也是一个io.Closer，在请求后会关闭。

###func (*Client) PostForm
```go
func (c *Client) PostForm(url string, data url.Values) (resp *Response, err error)
```
PostForm分发一个指定URL的POST请求，它带有`data`的keys和URL-encoded 的values作为请求体。

当`err`为空，`resp`常常包含一个non-nil resp.Body。当完成了从resp.Body读数据，调用者应该关闭它。

###type CloseNotifier interface
```go
type CloseNotifier interface {
    CloseNotify() <-chan bool
}
```
CloseNotifier接口通过ResponseWriters接口实现。当底层连接消失之后ResponseWriters允许检测。

这个机制可以用在：如果客户端在response准备好之前已经失去连接，那么取消长时间操作服务器。

###type ConnState
```go
type ConnState int
```
ConnState 代表客户端向服务器的连接状态。它可以被可选的Server.ConnState hook使用。
```go
const (
    StateNew ConnState = iota
    StateActive 
    StateIdle
    StateHijacked
    StateClosed
)
```

###func (ConnState) String
```go
func (c ConnState) String() string
```

###type Cookie struct
```go
type Cookie struct {
    Name       string
    Value      string
    Path       string
    Domain     string
    Expires    time.Time
    RawExpires string
    MaxAge   int
    Secure   bool
    HttpOnly bool
    Raw      string
    Unparsed []string // Raw text of unparsed attribute-value pairs
}
```
Cookie代表HTTP cookie ，它在HTTP响应的Set-Cookie头或者HTTP请求的Cookie头中发送。

###func (*Cookie) String
```go
func (c *Cookie) String() string
```
String返回序列化的cookie，它在Cookie头中使用（只要设置了Name和Value）或者在Set-Cookies响应头中使用（如果其他fields都被设置）。

###type CookieJar interface
```go
type CookieJar interface {
    SetCookies(u *url.URL, cookies []*Cookie)
    Cookies(u *url.URL) []*Cookie
}
```
###type Dir
```go
type Dir string
```
Dir使用限制在具体的目录树的本地文件系统实现了http.FileSystem。空Dir会作为`"."`。

###func (Dir) Open
```go
func (d Dir) Open(name string) (File, error)```

###type File interface
```go
type File interface {
    io.Closer
    io.Reader
    Readdir(count int) ([]os.FileInfo, error)
    Seek(offset int64, whence int) (int64, error)
    Stat() (os.FileInfo, error)
}
```
File是FileSystem 的Open方法返回的，它可以通过FileServer的实现来提供服务。其方法应该和*os.File的方法行为一样。

###type FileSystem interface
```go
type FileSystem interface {
    Open(name string) (File, error)
}
```
FileSystem 实现了获取一系列有名字的文件的方法。在文件路径中的元素用正斜杠('/', U+002F) 字符分隔，忽略主机操作系统的转换。

###type Flusher interface
```go
type Flusher interface {
    // Flush sends any buffered data to the client.
    Flush()
}
```
Flusher 是一个通过ResponseWriters实现的接口。ResponseWriters允许HTTP handler 向客户端冲洗缓冲数据。

注意虽然ResponseWriters支持Flash，如果客户端通过HTTP代理连接，缓冲的数据可能直到响应完成才会到达客户端。

>Flusher的作用是被Handler调用来将写缓存中的数据推给客户端

###type Handler interface
```go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```
实现了Handler接口的类型可以被注册，用来服务一个特定的路径或者在HTTP 服务器的subtree。

ServeHTTP应该向ResponseWriter 写入回答头和数据，然后返回。返回意味着请求结束，HTTP 服务器可以移向连接上的下一次的请求。

###func FileServer
```go
func FileServer(root FileSystem) Handler
```
FileServer 返回一个handler，它用文件系统根的内容来服务于提供HTTP请求。

为了使用操作系统的文件系统实现，使用http.Dir：
```go
http.Handle("/", http.FileServer(http.Dir("/tmp")))
```

###func NotFoundHandler
```go
func NotFoundHandler() Handler
```
返回一个简单的请求handler，它用404 page not found来回复每个请求。

###func RedirectHandler
```go
func RedirectHandler(url string, code int) Handler
```
RedirectHandler返回一个简单的请求handler，它使用给定的状态码对应的url重定向每个请求。

###func StripPrefix
```go
func StripPrefix(prefix string, h Handler) Handler
```
StripPrefix返回一个handler，它通过从请求的url的路径去掉给定的前缀并处罚handler h来提供服务。如果请求路径不以prefix为前缀，StripPrefix 用 HTTP 404 not found error处理它们。 

###func TimeoutHandler
```go
func TimeoutHandler(h Handler, dt time.Duration, msg string) Handler
```
TimeoutHandler返回一个Handler，它运行带有给定时间限制的`h`。

新的Handler调用h.ServeHTTP来处理每个请求，但是如果一个调用运行超过了时间限制，Handler回应503 Service Unavailable错误和body中给出的信息。（如果信息为空的，则会发送一个适当的默认信息）。在这个超时之后，h向它的ResponseWriter写入的writes会返回ErrHandlerTimeout。

###type HandlerFunc
```go
type HandlerFunc func(ResponseWriter, *Request)
```
HandlerFunc是一个适配器来允许使用普通的函数函数作为HTTP handlers。如果函数带有合适的签名，HandlerFunc(f)是调用f的Handler类型。

###func (HandlerFunc) ServeHTTP
```go
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request)
```
ServeHTTP 调用`f(w,r)`。

>这里需要多回味一下了，这个HandlerFunc定义和ServeHTTP合起来是说明了什么？说明HandlerFunc的所有实例是实现了ServeHttp方法的。另，实现了ServeHttp方法就是什么？实现了接口Handler!
所以你以后会看到很多这样的代码：

```go
func AdminHandler(w ResponseWriter, r *Request) {
    ...
}
handler := HandlerFunc(AdminHandler)
handler.ServeHttp(w,r)
 ```
 
>请不要讶异，你明明没有写ServeHttp，怎么能调用呢？ 实际上调用ServeHttp就是调用AdminHandler。

###type Header
```go
type Header map[string][]string
```
一个Header类型的数据代表HTTP 头中的键值对。

###func (Header) Add
```go
func (h Header) Add(key, value string)
```

###func (Header) Del
```go
func (h Header) Del(key string)
```

###func (Header) Get
```go
func (h Header) Get(key string) string
```

###func (Header) Set
```go
func (h Header) Set(key, value string)
```
设置头条目相关联的键为单值。如果存在key对应的值，则替换。

###func (Header) Write
```go
func (h Header) Write(w io.Writer) error
```
以wire格式写入header。

###func (Header) WriteSubset
```go
func (h Header) WriteSubset(w io.Writer, exclude map[string]bool) error
```
WriteSubset以wire格式写入头部。如果`exclude`不是空的，满足 exclude[key] == true的keys不会写入。

###type Hijacker interface
```go
type Hijacker interface {
	// 这个方法让调用者主动管理连接
     Hijack() (net.Conn, *bufio.ReadWriter, error)
}
```
> Hijacker的作用是被Handler调用来关闭连接的

###type ProtocolError struct
```go
type ProtocolError struct {
    ErrorString string
}
```

###func (*ProtocolError) Error
```go
func (err *ProtocolError) Error() string
```

###type Request struct
```go
type Request struct {   
    Method string
    URL *url.URL
    Proto      string // "HTTP/1.0"
    ProtoMajor int    // 1
    ProtoMinor int    // 0
    Header Header
    Body io.ReadCloser
    ContentLength int64
    TransferEncoding []string
    Close bool
    Host string
    Form url.Values
    PostForm url.Values
    MultipartForm *multipart.Form
    Trailer Header
    RemoteAddr string
    RequestURI string
    TLS *tls.ConnectionState
}
```

###func NewRequest
```go
func NewRequest(method, urlStr string, body io.Reader) (*Request, error)
```
给定method、URL和可选的body，NewRequest返回一个新的请求。

如果给定的body也是一个io.Closer，返回的Request.Body设为body，并且会被 Client 方法 Do、 Post和 PostForm以及Transport.RoundTrip关闭。

###func ReadRequest
```go
func ReadRequest(b *bufio.Reader) (req *Request, err error)
```
ReadRequest读取并解析`b`的请求。

###func (*Request) AddCookie
```go
func (r *Request) AddCookie(c *Cookie)
```
AddCookie向请求加入cookie。每个RFC 6265 section 5.4，AddCookie不会加入多于一个Cookie头域。这意味着所有的cookies，都被写入同一行，用分号隔开。

###func (*Request) Cookie
```go
func (r *Request) Cookie(name string) (*Cookie, error)
```
返回请求中命名的cookie或者ErrNoCookie错误，如果没有找到。

###func (*Request) Cookies
```go
func (r *Request) Cookies() []*Cookie
```
解析并返回在请求中发送的HTTP cookies 。

###func (*Request) FormFile
```go
func (r *Request) FormFile(key string) (multipart.File, *multipart.FileHeader, error)
```
FormFile返回`key`对应的第一个文件。如果有必要，FormFile调用ParseMultipartForm 和 ParseForm 。

###func (*Request) FormValue
```go
func (r *Request) FormValue(key string) string
```
返回与query中的命名的成员对应的第一个值。POST和PUT body 参数优先于URL query字符串值。如果有必要，FormValue 调用ParseMultipartForm 和 ParseForm。为了获得同一个key对应的更多的值，使用ParseFom。

###func (*Request) MultipartReader
```go
func (r *Request) MultipartReader() (*multipart.Reader, error)
```
如果这是一个 multipart/form-data的POST请求，MultipartReader返回一个MIME多部分的reader，或者返回nil和一个错误。使用这个函数而不是ParseMultipartForm 来将请求体作为流来处理。

###func (*Request) ParseForm
```go
func (r *Request) ParseForm() error
```
ParseForm从URL的原始的query解析并更新r.Form。

对于POST和PUT请求，它也会将请求体作为form解析，然后将结果写入r.PostForm和r.Form。POST和PUT体参数优先于r.Form中的URL query 字符串值。

如果请求体的大小没有被MaxBytesReader限制，那么其容量大小为10MB。

ParseMultipartForm自动调用ParseForm。它是独立的。

###func (*Request) ParseMultipartForm
```go
func (r *Request) ParseMultipartForm(maxMemory int64) error
```
ParseMultipartForm解析 multipart/form-data的 request。整个 request的 body都会被解析，文件中最多有maxMemory字节的被存储在内存中，其余存储在临时文件中。如果需要 ParseMultipartForm会自行调用 ParseForm。调用完 ParseMultipartForm，后续的各种方法的调用不受影响。

###func (*Request) PostFormValue
```go
func (r *Request) PostFormValue(key string) string
```
对于POST或者PUT中命名的成员，PostFormValue返回它的第一个值。忽略URL query参数。如果有必要，PostFormValue调用ParseMultipartForm 和 ParseForm。

###func (*Request) ProtoAtLeast
```go
func (r *Request) ProtoAtLeast(major, minor int) bool
```
ProtoAtLeast返回 request使用的协议是否不低于 major.minor指定的版本。

###func (*Request) Referer
```go
func (r *Request) Referer() string
```
Referer返回一个表示引用的 URL，如果 request里有的话。
request中的 Referer拼写错了，这是 HTTP早期时犯的错误。这个值也可以通过 map类型变量 的Header["Referer"]来取得；使用方法的好处是编译器可以诊断那些使用正确拼法的程序（调用req.Referrer），但却不能诊断使用Header["Referrer"]的程序。

###func (*Request) SetBasicAuth
```go
func (r *Request) SetBasicAuth(username, password string)
```
SetBasicAuth设置request的Authorization header以便使用HTTP Basic Authentication，它带有username、password两个参数。

使用HTTP Basic Authentication，不会加密用户名和密码。

###func (*Request) UserAgent
```go
func (r *Request) UserAgent() string
```
如果在请求中发送client的User-Agent，则返回它。

###func (*Request) Write
```go
func (r *Request) Write(w io.Writer) error
```
Write写入HTTP/1.1请求的头和体，以wire格式。它考虑以下的请求域：
```go
Host
URL
Method (defaults to "GET")
Header
ContentLength
TransferEncoding
Body
```
如果有Body，Content-Length<=0，TransferEncoding 没有被设为”identity“，Write向头部加入Transfer-Encoding: chunked。在它被发送之后，Body被关闭。

###func (*Request) WriteProxy
```go
func (r *Request) WriteProxy(w io.Writer) error
```
WriteProxy类似Write，但是期望以HTTP代理的形式写入。特别注意，WriteProxy使用绝对URL写入request的原始Request-URL line，每部分按照section 5.1.2 of RFC 2616，包括scheme和host。在其他情形，WriteProxy还会使用r.Host或r.URL.Host写入Host header。

###type Response struct
```go
type Response struct {
    Status     string // e.g. "200 OK"
    StatusCode int    // e.g. 200
    Proto      string // e.g. "HTTP/1.0"
    ProtoMajor int    // e.g. 1
    ProtoMinor int    // e.g. 0
    Body io.ReadCloser   
    ContentLength int64   
    TransferEncoding []string
    Close bool    
    Trailer Header
    Request *Request
    TLS *tls.ConnectionState
}
```
>Response实现了ResponseWriter,Flusher,Hijacker这三个接口

###func Get
```go
func Get(url string) (resp *Response, err error)
```
Get分发指定URL的GET请求。如果响应是以下重定向码，Get跟在重定向后面，最多10个重定向：
```go
301 (Moved Permanently)
302 (Found)
303 (See Other)
307 (Temporary Redirect)
```
如果有太多的重定向或者有HTTP 协议错误，将会返回一个错误。一个non-2xx响应不会引起错误。

当`err`为空，`resp`常常包含一个non-nil resp.Body。当完成了从resp.Body读数据，调用者应该关闭它。

Get是DefaultClient.Get的封装。

###func Head
```go
func Head(url string) (resp *Response, err error)
```
Head向指定的URL分发HEAD。如果响应是以下的重定向码，Head调用客户端的CheckRedirect函数后跟随重定向：
```go
301 (Moved Permanently)
302 (Found)
303 (See Other)
307 (Temporary Redirect)
```
Head是DefaultClient.Head的封装。

###func Post
```go
func Post(url string, bodyType string, body io.Reader) (resp *Response, err error)
```
Post分发指定的URL的POST请求。

当完成了从resp.Body读数据，调用者应该关闭它。

Post是DefaultClient.Post的封装。

###func PostForm
```go
func PostForm(url string, data url.Values) (resp *Response, err error)
```
PostForm分发一个指定URL的POST请求，它带有`data`的keys和URL-encoded 的values作为请求体。

当`err`为空，`resp`常常包含一个non-nil resp.Body。当完成了从resp.Body读数据，调用者应该关闭它。

PostForm是DefaultClient.PostForm的封装。

###func ReadResponse
```go
func ReadResponse(r *bufio.Reader, req *Request) (*Response, error)
```
func ReadResponse从`r`读取并返回一个HTTP 响应。req参数根据需要确定了和这个响应符合的request。如果是空的，那么假设为GET请求。当完成了resp.Body的读取后，客户端必须调用resp.Body.Close。调用之后，额客户端可以检查 resp.Trailer来找到响应trailer中的键值对。

###func (*Response) Cookies
```go
func (r *Response) Cookies() []*Cookie
```
解析并返回在Set-Cookie头中的cookies。

###func (*Response) Location
```go
func (r *Response) Location() (*url.URL, error)
```

###func (*Response) ProtoAtLeast
```go
func (r *Response) ProtoAtLeast(major, minor int) bool
```

###func (*Response) Write
```go
func (r *Response) Write(w io.Writer) error
```
向响应（header、body和trailer）以wire格式写入。这个方法考虑了以下的响应域:
```go
StatusCode
ProtoMajor
ProtoMinor
Request.Method
TransferEncoding
Trailer
Body
ContentLength
Header, values for non-canonical keys will have unpredictable behavior
```
在它发送之后，Body关闭。

###type ResponseWriter interface
```go
type ResponseWriter interface {
	//这个方法返回Response返回的Header供读写
    Header() Header
    // 这个方法写Response的Body
    Write([]byte) (int, error)
	 // 这个方法根据HTTP State Code来写Response的Header
    WriteHeader(int)
}
```
ResponseWriter 接口被HTTP handler用来构建HTTP 响应。

> ResponseWriter的作用是被Handler调用来组装返回的Response的

###type RoundTripper interface
```go
type RoundTripper interface {
    RoundTrip(*Request) (*Response, error)
}
```

###func NewFileTransport
```go
func NewFileTransport(fs FileSystem) RoundTripper
```
NewFileTransport返回一个新的RoundTripper，服务于给定的FileSystem。返回的RoundTripper忽略到达的请求中的URL host，向大部分其他请求属性一样。

典型使用NewFileTransport情形是用Transport注册文件协议：
```go
t := &http.Transport{}
t.RegisterProtocol("file", http.NewFileTransport(http.Dir("/")))
c := &http.Client{Transport: t}
res, err := c.Get("file:///etc/passwd")
...
```

###type ServeMux struct
```go
type ServeMux struct {
    // contains filtered or unexported fields
}
```
>它就是http包中的路由规则器。你可以在ServerMux中注册你的路由规则，当有请求到来的时候，根据这些路由规则来判断将请求分发到哪个处理器（Handler）。

>当一个请求request进来的时候，server会依次根据ServeMux.m中的string（路由表达式）来一个一个匹配，如果找到了可以匹配的muxEntry,就取出muxEntry.h,这是个handler，调用handler中的ServeHTTP（ResponseWriter, *Request）来组装Response，并返回。

###func NewServeMux
```go
func NewServeMux() *ServeMux
```

###func (*ServeMux) Handle
```go
func (mux *ServeMux) Handle(pattern string, handler Handler)
```
根据给定的patter注册了handler。如果已经存在与patter对应的handler,Handle引发异常。

###func (*ServeMux) HandleFunc
```go
func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request))
```
根据给定pattern，注册了handler函数。

###func (*ServeMux) Handler
```go
func (mux *ServeMux) Handler(r *Request) (h Handler, pattern string)
```

###func (*ServeMux) ServeHTTP
```go
func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request)
```
ServeHTTP向最匹配请求URL的handler分发请求。

>这个说明，ServeHttp也实现了Handler接口，它实际上也是一个Handler，内部实现调用handler。

###type Server struct
```go
type Server struct {
	//服务监听地址 
    Addr           string        // // TCP address to listen on, ":http" if empty
	//实现Handler接口的对象
    Handler        Handler       // handler to invoke, http.DefaultServeMux if nil
	//读超时时间
    ReadTimeout    time.Duration // maximum duration before timing out read of the request
	//写超时时间
    WriteTimeout   time.Duration // maximum duration before timing out write of the response
	//读取头数据的最大值
    MaxHeaderBytes int           // maximum size of request headers, DefaultMaxHeaderBytes if 0
    TLSConfig      *tls.Config   // optional TLS config, used by ListenAndServeTLS
    TLSNextProto map[string]func(*Server, *tls.Conn, Handler)   
    ConnState func(net.Conn, ConnState)   
    ErrorLog *log.Logger   // contains filtered or unexported fields
}
```
Server定义了启动HTTP 服务器的参数。零值的Server是有效的配置。

>ReadTimeout是读取TCP连接中的数据的超时时间，设置这个值有以下用处：
1. 避免长时间读取请求数据时候长时间阻塞
2. 在客户端以Keep Alive方式发起请求的时候，如果下一次请求迟迟不来，而服务端没有设置ReadTimtout值，则服务端会长时间挂起连接，而有读不到数据，这页很容易被用来攻击。

###func (*Server) ListenAndServe
```go
func (srv *Server) ListenAndServe() error
```
ListenAndServe监听TCP网络地址srv.Addr，然后调用Serve来处理到达的请求。如果srv.Addr是blank，那么使用“.http”。

###func (*Server) ListenAndServeTLS
```go
func (srv *Server) ListenAndServeTLS(certFile, keyFile string) error
```

###func (*Server) Serve
```go
func (srv *Server) Serve(l net.Listener) error
```
Serve在Listener `l`上接收连接，为每一个连接创建一个goroutine。goroutine读取请求然后调用srv.Handler来回复它们。

###func (*Server) SetKeepAlivesEnabled
```go
func (s *Server) SetKeepAlivesEnabled(v bool)
```
SetKeepAlivesEnabled控制HTTP keep-alive是否生效。默认情况下，keep-alive常常是使能的。只有在资源非常有限的环境或者服务器正在关闭时才应该禁用它们。


###type Transport struct
```go
type Transport struct {   
    Proxy func(*Request) (*url.URL, error)
    Dial func(network, addr string) (net.Conn, error)
    TLSClientConfig *tls.Config
    TLSHandshakeTimeout time.Duration
    DisableKeepAlives bool
    DisableCompression bool
    MaxIdleConnsPerHost int
    ResponseHeaderTimeout time.Duration   
}
```

###func (*Transport) CancelRequest
```go
func (t *Transport) CancelRequest(req *Request)
```
CancelRequest通过关闭连接取消了in-flight 请求。

###func (*Transport) CloseIdleConnections
```go
func (t *Transport) CloseIdleConnections()
```
CloseIdleConnections关闭任意一个这样的连接：它在上一次请求连接，但是现在却在keep-alive无所事事。它不会打断正在使用的连接。

###func (*Transport) RegisterProtocol
```go
func (t *Transport) RegisterProtocol(scheme string, rt RoundTripper)
```
RegisterProtocol用scheme注册了一个新的协议。Transport会使用给定的scheme向`rt`传递请求。`rt`有责任模拟HTTP请求的语义。

RegisterProtocol可以被其他包使用来提供像ftp或者file这样的协议方案。

###func (*Transport) RoundTrip
```go
func (t *Transport) RoundTrip(req *Request) (resp *Response, err error)
```
RoundTrip实现了RoundTripper接口。

为了高层的HTTP客户端支持（比如处理cookies和redirects），请看 Get、 Post、 和 the Client type。
