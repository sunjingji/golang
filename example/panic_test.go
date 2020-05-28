// @Title 异常示例
package example

import (
  "errors"
  "fmt"
)

// panic就是异常，和C++中的异常机制非常相似。
// 1)异常可以由运行时抛出，也可以自己抛出。
// 2)如果异常得不到处理，就会传播到上级调用函数，直到程序崩溃。
// 3)处理异常基于 defer/recover。
// 4)在defer语句每次执行的时候，Go语言会把它携带的defer函数及其参数值另行存储到一个栈中，函数退出时依次弹出执行。
// 5)有些异常时不可恢复的，程序只能崩溃。
//
// error是可预期的错误，panic是异常。错误和异常分开，的确要好一些，可以避免异常被滥用。在很多时候为了省事儿，有人会用try/catch代替错误检测代码，不做错误检测。
func ExamplePanic() {
  defer func() {  // 无论什么原因退出函数，defer都会被执行。
    if p := recover(); p != nil {
      fmt.Printf("2#%v", p)
    }
  }()

  // 引发panic
  panic(errors.New("something wrong"))
  p := recover()  // 抛出异常后控制权转移，不会再按照正常的流程执行。
  fmt.Printf("1#%v", p)

  // Output: 2#something wrong
}
