#iotuil包

import "io/ioutil"

---

##概览
ioutil包实现了一些实用的I/O函数。

##变量
```go
var Discard io.Writer = devNull(0)
```
Discard是一个io.Writer，对其进行的所有Write调用都会成功但不会做任何实际的操作。 

###func NopCloser
```go
func NopCloser(r io.Reader) io.ReadCloser
```
NopCloser返回一个ReadCloser对象，带有no-op Close方法的函数，并封装了给定的Reader对象 `r`。

###func ReadAll
```go
func ReadAll(r io.Reader) ([]byte, error)
```
从r读取直到遇到error或EOF并返回读取的数据。 如果调用成功，返回的err为nil，而不是EOF。因为ReadAll定义为从源读取数据直到EOF，它不会将从r读取的EOF视为报错。

###func ReadDir
```go
func ReadDir(dirname string) ([]os.FileInfo, error)
```
接收`dirname`指定的目录参数，并返回一个有序的且带有子目录信息的文件列表。 

###func ReadFile
```go
func ReadFile(filename string) ([]byte, error)
```
从`filename`指定的文件中读取数据并返回文件的内容。 若调用成功，返回的err为nil，而不是EOF。因为ReadFile定义为从源读取数据直到EOF，它不会将从r读取的EOF视为报错。 

###func TempDir
```go
func TempDir(dir, prefix string) (name string, err error)
```
在指定的`dir`目录里创建一个新的、使用`prfix`作为前缀的临时文件夹，并返回文件夹的路径。 如果`dir`是空字符串，TempDir使用默认的临时目录（参考os.TempDir函数）。 如果多个程序同时调用该函数的话，将会创建不同的临时目录（因此是线程安全的）。调用本函数的程序有责任在不需要临时文件夹时删除它。

###func TempFile
```go
func TempFile(dir, prefix string) (f *os.File, err error)
```
在`dir`目录下创建一个新的、使用`prefix`为前缀的临时文件，以读写模式打开该文件并返回os.File指针。 如果`dir`是空字符串，TempFile使用默认的临时目录（参考os.TempDir函数）。 如果多个程序调用该函数的话，将会创建不同的临时文件（因此是线程安全的）。
调用者可以使用f.Name()函数来找到文件的pathname。
调用本函数的程序有责任在不需要临时文件时删除它。 

###func WriteFile
```go
func WriteFile(filename string, data []byte, perm os.FileMode) error
```
向名为filename的文件写入数据。如果文件不存在，则创建带有权限信息的文件，否则在写入前截断。

