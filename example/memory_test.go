//@Title 内存模型
package example

// # 内存模型
// Go规定了不同goroutine之间的内存可见性模型。
// 一个goroutine对内存的操作，在另一个goroutine看到的结果是不确定的，可能存在"乱序"的情况，也可能根据就看不到。
// 因此，如果在不同goroutine之间共享数据，必须进行同步。
// channel: 对channel(要注意区分带缓存和不带缓存的channel)的读写，相当于插入了一个"内存屏障"，类似c++的ack/rel，可以channel读写前后的内存访问。
// sync/atomic: 原子操作，没有对非atomic内存操作排序的作用。
// sync/mutex: 互斥锁。
//
// 参考go官方博客：http://golang.org/ref/mem
