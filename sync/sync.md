#sync

import "sync"

##简介

sync提供了基本的同步操作原语.

##概览

sync提供了基本的同步操作原语,例如互斥锁之类.除了Once和WaitGroup之外,其它大多数是用于底层库.高层同步操作最好是通过channel和通信实现.

包含了这个包所定义的类型的值不可复制.


##内容

###type Cond
```go
type Cond struct {
    // 锁L用于观察和改变条件
    L Locker
    // 包含了被过滤或者未导出的字段
}
```
Cond实现了一个条件变量,给协程提供了一个集合点用于等待或者宣布一个事件的发生.

每个Cond拥有一个关联的Locker L(通常是一个\*Mutex或者\*RWMutex),当要改变条件和调用Wait方法时候都需要先对L加锁.

Cond可以创建为其它结构体的一部分.Cond在第一次使用后不能被复制.


###func NewCond
```go
func NewCond(l Locker) *Cond
```
NewCond返回一个新的关联Locker L的Cond.

###func (*Cond) Broadcast
```go
func (c *Cond) Broadcast()
```
BroadCase唤醒全部等待c的协程.

在调用的过程中,允许但不是必需要求调用者对c.L加锁.


###func (*Cond) Signal
```go
func (c *Cond) Signal()
```
Signal唤醒一个等待c的协程.

在调用的过程中,允许但不是必需要求调用者持有对c.L加锁.


###func (*Cond) Wait
```go
func (c *Cond) Wait()
```
Wait自动解锁c.L,然后暂停调用协程的执行.在稍后的恢复执行时,Wait在返回前先锁住c.L.不像其它系统,Wait不会返回除非被Broadcast或者Signal唤醒.

由于c.L在等待第一次恢复时没有加锁,通常情况下,调用者在Wait返回的时候不能认为条件就是为真.因此,调用者应该在循环中使用Wait:

```go
c.L.Lock()
for !condition() {
	c.Wait()
}
... make use of conditio ...
c.L.Unlock()
```

###type Locker
```go
type Locker interface {
    Lock()
    Unlock()
}
```
一个Locker表示一个对象可以加锁和解锁.

###type Mutex
```go
type Mutex struct {
    // 包含了被过滤或者未导出的字段
}
```
Mutex是一个互斥锁.Mutex可以作为其它结构体的一部分;零值Mutex是一个未加锁的互斥锁.


###func (*Mutex) Lock
```go
func (m *Mutex) Lock()
```
加锁.如果已经加过锁了,调用协程会一直阻塞直到互斥锁可以加锁.


###func (*Mutex) Unlock
```go
func (m *Mutex) Unlock()
```
解锁.如果在进入Unlock时m还没有加锁,会出现运行时错误.

一个加了锁的Mutex不跟特定的协程绑定.允许在一个协程中对Mutex加锁而在另外的协程中解锁.


###type Once
```go
type Once struct {
    // 包含了被过滤或者未导出的字段
}
```
Once是一个只能执行一次动作的对象.

###func (*Once) Do
```go
func (o *Once) Do(f func())
```
只有当前这个Once的实例第一次调用Do方法的时候,Do才会去执行函数f.换句话说,有以下实例:
  
    var once Once
    
如果 once.Do(f) 被调用了很多次,只有第一次调用Do时会调用函数f,即使每次调用Do的时候f的值不一样.对于每一个需要被执行的函数,都需要一个新的Once实例.

Do是为了满足那些需要只运行一次初始化的需求.由于f是无参函数,所以可能需要使用字面函数(function literral)捕捉Do要执行的函数所需要的参数.

    config.once.Do(func() { config.init(filename) })

在调用f那次返回之前,没有其它对Do的调用能够返回,因此如果f引起Do再次被调用,就会引起死锁.


###type Pool
```go
type Pool struct {

    // New是可选的,指定了一个函数用于生成一个新的值.
    // 在Get返回nil的情况下调用New.
    // It may not be changed concurrently with calls to Get.
    New func() interface{}
    // 包含了被过滤或者未导出的字段
}
```
Pool是一组可以单独存储和检索的临时对象集合.

任何存储在Pool中的元素可能在任意时间被自动删除,且不会有任何通知.如果Pool持有一个元素的唯一索引,那么元素可能会被释放.

Pool是多协程并发安全.

Pool的目的是缓存那些已经分配但是还未使用的元素供后续使用,这样可以缓解垃圾回收的压力.这就是,它可以很方便地建立有效的,线程安全的自由列表.虽然它并不适用于全部自由列表.

一个合适的使用场景是:一个包需要为其并发使用又互相独立的用户在后台默默共享和重用一些临时元素,可以使用Pool来管理.Pool提供了一种方法,可以在众多客户端之间摊销分配成本.

一个很好的使用Pool的例子是在`fmt`包中,管理着大小不定的临时输出缓冲区库.库在高负载下(当很多协程频繁进行输出)会增大,在不活跃时时会收缩.

另一方面,自由列表作为短生命期对象的一部分不是Pool的合适用法,因为在这种情况下,开销没法很好均摊.让这样的对象实现自己的自由列表是更有高效的.(free list译作自由列表有点牵强)


###func (*Pool) Get
```go
func (p *Pool) Get() interface{}
```
Get从Pool中任意选择一个元素,从Pool中删除,然后返回给调用者.Get可能会选择忽视内存池把它当成空的.调用者不应该假设传递给Put的值跟Get返回的值这两者之间有任何关系.

如果Get在其它情况下将要返回nil,p.New又不为空,那么Get将返回调用p.New的结果.


###func (*Pool) Put
```go
func (p *Pool) Put(x interface{})
```
Put把x放到内存池中.

###type RWMutex
```go
type RWMutex struct {
    // 包含了被过滤或者未导出的字段
}
```
RWMutex是一个读写互斥锁.锁可以被任意数量的读取者或者一个写入者持有.RWMutex可以创建为其它结构体的一部分;零值的RWMutex是一个未加锁的互斥量.


###func (*RWMutex) Lock
```go
func (rw *RWMutex) Lock()
```
Lock加写锁.如果已经加过读锁或者写锁,Lock操作会一直阻塞直到可以加锁为止.为了保证最终能够获得写锁,已经阻塞住的Lock方法会阻止新的读取者获取锁权限.


###func (*RWMutex) RLock
```go
func (rw *RWMutex) RLock()
```
RLock加只读锁.

###func (*RWMutex) RLocker
```go
func (rw *RWMutex) RLocker() Locker
```
RLocker返回一个Locker类型接口,调用rw.RLock和rw.RUnlock实现了Lock和Unlock接口.

###func (*RWMutex) RUnlock
```go
func (rw *RWMutex) RUnlock()
```
RUnlock撤销一个RLock调用;这不影响其它并发的读取者.如果在进入RUnlock时rm还没有加过读锁,会出现运行时错误.


###func (*RWMutex) Unlock
```go
func (rw *RWMutex) Unlock()
```
Unlock撤销rw的写锁.如果在进入Unlock时rw还没有加过写锁,会出现运行时错误.

跟Mutex一样,一个加了锁的Mutex不跟特定的协程绑定.允许在一个协程中对RWMutex加锁而在另外的协程中解锁.


###type WaitGroup
```go
type WaitGroup struct {
    // 包含了被过滤或者未导出的字段
}
```
WaitGroup用于等待一个协程集合完成运行.在主协程中调用Add方法设置需要等待的协程数.然后各个协程在运行结束的时候调用Done.同时Wait用于阻塞直到全部协程运行结束.


###func (*WaitGroup) Add
```go
func (wg *WaitGroup) Add(delta int)
```
Add加上变量delta到WaitGroup计数器中,delta可以是负数的.如果计数器变成0,全部阻塞在Wait操作的协程会被唤醒.如果计数器变成负数,Add会引发panic.

注意应该在调用Wait之前带正数的参数调用Add,否则Wait可能会等到较小的一组协程.通常这意味着需要在创建协程的语句之前,或者其它等待的事件之前就应该执行调用Add方法.


###func (*WaitGroup) Done
```go
func (wg *WaitGroup) Done()
```
Done减少WaitGroup的计数器.

###func (*WaitGroup) Wait
```go
func (wg *WaitGroup) Wait()
```
Wait阻塞等待WaitGroup计数器变成0后返回.


