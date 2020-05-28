//@Title Context
package example

// # Context
//   context的实际值大体上分为三种，即：根context(Background()、TODO())、可撤销的context(有分为只可手动撤销的Context，和可以定时撤销的Context)和带value值的context。
//   所有的Context值共同构成了一颗上下文树。
//   一个撤销事件，会向所有可撤销子context传播。
//
//   在context调用Value()函数，当前context是不带值的context时，会向parent context传递调用。
//   在当前函数退出时，函数内创建的goroutine可能会继续执行下去，这就是所谓的"goroutine leak"，使用context可以优雅得解决这个问题。
//
//   // gen is a generator that can be cancellable by cancelling the ctx.
//   func gen(ctx context.Context) <-chan int {
//     ch := make(chan int)
//	   go func() {
//	     var n int
//	     for {
//		     select {
//	 	       case <-ctx.Done():
//	            return // avoid leaking of this goroutine when ctx is done.
//		       case ch <- n:
//		          n++
//		     }
//	     }
//     }()
//	   return ch
//   }
//
//  ctx, cancel := context.WithCancel(context.Background())
//  defer cancel() // make sure all paths cancel the context to avoid context leak
//
//  for n := range gen(ctx) {
//    fmt.Println(n)
//    if n == 5 {
//      cancel()
//      break
//     }
//  }
//
//  // ...
//
//  参考：Go Concurrency Patterns: Context https://blog.golang.org/context
//  参考：Go Concurrency Patterns: Pipelines and cancellation  https://blog.golang.org/pipelines
