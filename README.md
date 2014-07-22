# Golang标准库中文版

###翻译By
Go友团([http://golanghome.com](http://golanghome.com))

###项目地址
[GoStandardLibrary-Chinese](http://github.com/jemygraw/GoStandardLibrary-Chinese)

###成员  
[@jemygraw](http://github.com/jemygraw), [@itfanr](http://github.com/itfanr), [@getwe](http://github.com/getwe)

###致谢  
1. 感谢 [@ghosert](http://weibo.com/ghosert) 创立的 [作业部落](https://www.zybuluo.com/mdeditor) ，史上最好的Markdown编辑器。  
2. 感谢Google为我们提供的革命性编程语言[Golang](http://golang.org)。  
3. 感谢所有支持我们，为我们的努力欢呼的小伙伴们！  
4. 感谢上天给我一次来到地球的机会。哈哈！  

###包翻译状态  
1. 未翻译 － 该包未被占用也未被翻译。  
2. 已占用 － 该包正在被别人翻译，优先merge该成员的pull request，当该成员未能如期完成时，包状态恢复未翻译，可以再次申请占用。  
3. 已翻译 － 该包已经被翻译，并且pull request已经merge，你可以查看翻译，并给出建议。  

###参与方法  
1. 首先Fork项目。  
2. 然后查看项目README文件选择一个未被占用或未被翻译的包，到[这里](https://github.com/jemygraw/GoStandardLibrary-Chinese/issues/1)申请占用，并翻译。  
3. 申请占用翻译的时候，需提供你能够预期完成的时间，超过该时间后，pull request将不再被接受，并且该包状态恢复未翻译，可供他人申请占用翻译。  
4. 翻译完成之后，创建pull request，请求merge。merge完成后，包状态改为已翻译，并且署名。  
5. 翻译以包为单位，如果有子包，则上一层包和子包属于不同的包。  
6. 可以对自己或他人翻译完成的包提交校验pull request，这时将署名为校验人  。

###标准库网址

如果你无法访问golang.org可以访问以下的两个网址来作为Go Standard Library参考。

 -  http://godoc.golangtc.com/doc/
 -  http://docs.studygolang.com/doc/

###翻译文件规范  
为了统一翻译风格，请参考文件[archive_tar.md](https://github.com/jemygraw/GoStandardLibrary-Chinese/blob/master/archive/archive_tar.md)。  
主要内容有下面几点：  
1. 翻译的标准库以Go1.3的为准。  
2. 翻译文件名称以`包名.md`方式命名，如果存在子包，则以`包名_子包名.md`方式命名。存在多层子包，则依次加上`_子包名`来命名。例如：`archive_tar.md`，`database_sql_driver.md`等。  
3. 文件内容以`#包名`开头。这里的包名为import导入时的包名。  
4. 然后是import导入包的方式，比如`import "archive/tar"`。  
5. 下面是`简介(Synopsis)`，就是标准库Packages页面的包简单描述，使用二级标题(`##`)。  
6. 然后是`概览(Overview)`，为标准库详细页面的描述，不需要包括示例(Examples)，使用二级标题(`##`)。  
7. 然后是`内容`，单独占一行，使用二级标题(`##`)。  
8. 然后依次按照标准库页面内容，翻译每一项。每一项的标题使用三级标题(`###`)表示，然后是原型代码，然后是详细描述。  
9. 为了使得翻译文件内容清晰，每个自然段落多加一个空行。  
10. Merge的时候，由[@jemygraw](https://github.com/jemygraw)统一检查完格式后merge。  


---

|   包名             |               翻译人                      |  状态  |   起始时间  | 预期结束时间 | 实际结束时间  |    校验人  |
|-------------------|--------------------------|------------------------------------------|-------|------------|------------|-------------|------------|
|[archive/tar]()| [@jemygraw](https://github.com/jemygraw) | 已翻译 | 2014/07/16 | 2014/07/16 |2014/07/16   |            |
|[archive/zip]()| [@jemygraw](https://github.com/jemygraw) | 已占用 | 2014/07/21 | 2014/07/22 |             |            |
|[bufio]()|[@getwe](https://github.com/getwe)|已占用|2014/07/23|2014/07/24||||
|[builtin]()||||||||
|[bytes]()||||||||
|[compress/bzip2]()|||||||
|[compress/flate]()|||||||
|[compress/gzip]()|||||||
|[compress/lzw]() |||||||
|[compress/zlib]() |||||||
|[container/heap]() |||||||
|[container/list]() |||||||
|[container/ring]() |||||||
|[crypto]()|||||||
|[crypto/aes]()|||||||
|[crypto/cipher]()|||||||
|[crypto/des]()|||||||
|[crypto/dsa]()|||||||
|[crypto/ecdsa]()|||||||
|[crypto/elliptic]()|||||||
|[crypto/hmac]()|||||||
|[crypto/md5]()|||||||
|[crypto/rand]()|||||||
|[crypto/rc4]()|||||||
|[crypto/rsa]()|||||||
|[crypto/sha1]()|||||||
|[crypto/sha256]()|||||||
|[crypto/sha512]()|||||||
|[crypto/subtle]()|||||||
|[crypto/tls]()|||||||
|[crypto/x509]()|||||||
|[crypto/x509/pkix]()|||||||
|[database/sql]()|||||||
|[database/sql/driver]()|||||||
|[debug/dwarf]()|||||||
|[debug/elf]()|||||||
|[debug/gosym]()|||||||
|[debug/macho]()|||||||
|[debug/pe]()|||||||
|[encoding]()|||||||
|[encoding/ascii85]()|||||||
|[encoding/asn1]()|||||||
|[encoding/base32]()|||||||
|[encoding/base64]()|||||||
|[encoding/binary]()|||||||
|[encoding/csv]()|||||||
|[encoding/gob]()|||||||
|[encoding/hex]()|||||||
|[encoding/json]()|||||||
|[encoding/pem]()|||||||
|[encoding/xml]()|||||||
|[errors]()|||||||
|[expvar]()|||||||
|[flag]()|||||||
|[fmt]()|||||||
|[go/ast]()|||||||
|[go/build]()|||||||
|[go/doc]()|||||||
|[go/format]()|||||||
|[go/packages]()|||||||
|[go/parser]()|||||||
|[go/printer]()|||||||
|[go/scanner]()|||||||
|[go/token]()|||||||
|[hash]()|||||||
|[hash/adler32]()|||||||
|[hash/crc32]()|||||||
|[hash/crc64]()|||||||
|[hash/fnv]()|||||||
|[html]()|||||||
|[html/template]()|||||||
|[image]()|||||||
|[image/color]()|||||||
|[image/color/palette]()|||||||
|[image/draw]()|||||||
|[image/gif]()|||||||
|[image/jpeg]()|||||||
|[image/png]()|||||||
|[index/suffixarray]()|||||||
|[io]()|||||||
|[io/ioutil]()|||||||
|[log]()|||||||
|[log/syslog]()|||||||
|[math]()|||||||
|[math/big]()|||||||
|[math/cmplx]()|||||||
|[math/rand]()|||||||
|[mime]()|||||||
|[mime/multipart]()|||||||
|[net]()|||||||
|[net/http]()|||||||
|[net/http/cgi]()|||||||
|[net/http/cookiejar]()|||||||
|[net/http/fcgi]()|||||||
|[net/http/httptest]()|||||||
|[net/http/httputil]()|||||||
|[net/http/pprof]()|||||||
|[net/mail]()|||||||
|[net/rpc]()|||||||
|[net/rpc/jsonrpc]()|||||||
|[net/smtp]()|||||||
|[net/textproto]()|||||||
|[net/url]()|||||||
|[os]()|||||||
|[os/exec]()|||||||
|[os/signal]()|||||||
|[os/user]()|||||||
|[path]()|||||||
|[path/filepath]()|||||||
|[reflect]()|||||||
|[regexp]()|||||||
|[regexp/syntax]()|||||||
|[runtime]()|||||||
|[runtime/cgo]()|||||||
|[runtime/debug]()|||||||
|[runtime/pprof]()|||||||
|[runtime/race]()|||||||
|[sort]()|||||||
|[strconv]()|||||||
|[strings]()|||||||
|[sync]()|||||||
|[sync/atomic]()|||||||
|[syscall]()|||||||
|[testing]()|||||||
|[testing/iotest]()|||||||
|[testing/quick]()|||||||
|[text/scanner]()|||||||
|[text/tabwriter]()|||||||
|[text/template]()|||||||
|[text/template/parse]()|||||||
|[time]()|||||||
|[unicode]()|||||||
|[unicode/utf16]()|||||||
|[unicode/utf8]()|||||||
|[unsafe]()|||||||
