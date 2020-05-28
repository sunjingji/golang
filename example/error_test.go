// @Title 错误示例
package example

import (
  "fmt"
  "os"
  "os/exec"
)
// ----------------------------------------------------------------------------
// 基础知识
// error 是一个接口，可以很方便地创建错误类型
// type error interface {
//   Error() string
// }
//
// errors.New() 创建的是errors包的缺省实现：errorString
// fmt.Errorf() 生成模板化的错误信息, 其内部调用的也是 errors.New()

// ----------------------------------------------------------------------------
// 问题1：对于具体错误的判断，Go 语言中都有哪些惯用法？
// 1)对于类型在已知范围内的一系列错误值，一般使用类型断言表达式或类型switch语句来判断；
// 2)对于已有相应变量且类型相同的一系列错误值，一般直接使用判等操作来判断；
// 3)对于没有相应变量且类型未知的一系列错误值，只能使用其错误信息的字符串表示形式来做判断。
func ExampleError() {
  // 执行os操作
  err := os.Mkdir("test", 666)

  // 获取os相关的错误
  err = underlyingError(err)

  // 对于已有相应变量且类型相同的一系列错误值，一般直接使用判等操作来判断。
  switch err {
    case os.ErrClosed:      // errors.New("file already closed")
      fmt.Printf("error(closed): %s\n", err)
    case os.ErrInvalid:     // errors.New("invalid argument")
      fmt.Printf("error(invalid): %s\n", err)
    case os.ErrPermission:  //errors.New("permission denied")
      fmt.Printf("error(permission): %s\n", err)
  }

  // Output:
}

// 获取和返回已知的操作系统相关错误的潜在错误值。
func underlyingError(err error) error {
  // 对于类型在已知范围内的一系列错误值，一般使用类型断言表达式或类型switch语句来判断；
  switch err := err.(type) {
  case *os.PathError:
    return err.Err
  case *os.LinkError:
    return err.Err
  case *os.SyscallError:
    return err.Err
  case *exec.Error:
    return err.Err
  }
  return err
}

// ----------------------------------------------------------------------------
// 问题：怎样根据实际情况给予恰当的错误值？
// 构建错误值体系的基本方式有两种，即：创建立体的错误类型体系和创建扁平的错误值列表。
//
// 1)错误类型体系
// net.Error接口除了拥有error接口的Error方法之外，还有两个自己声明的方法：Timeout和Temporary。
//   type Error interface {
//	   error
//	   Timeout() bool   // Is the error a timeout?
//	   Temporary() bool // Is the error temporary?
//   }
// net包中有很多错误类型都实现了net.Error接口，比如：
//   *net.OpError；
//   *net.AddrError；
//   net.UnknownNetworkError等等。
// 你可以把这些错误类型想象成一棵树，内建接口error就是树的根，而net.Error接口就是一个在根上延伸的第一级非叶子节点。
// 同时，你也可以把这看做是一种多层分类的手段。当net包的使用者拿到一个错误值的时候，可以先判断它是否是net.Error类型的，也就是说该值是否代表了一个网络相关的错误。
// 如果是，那么我们还可以再进一步判断它的类型是哪一个更具体的错误类型，这样就能知道这个网络相关的错误具体是由于操作不当引起的，还是因为网络地址问题引起的，又或是由于网络协议不正确引起的。
// 当我们细看net包中的这些具体错误类型的实现时，还会发现，与os包中的一些错误类型类似，它们也都有一个名为Err、类型为error接口类型的字段，代表的也是当前错误的潜在错误。
//
// 2)扁平的错误值列表
// 当我们只是想预先创建一些代表已知错误的错误值时候，用这种扁平化的方式就很恰当了。
// 不过，由于error是接口类型，所以通过errors.New函数生成的错误值只能被赋给变量，而不能赋给常量，又由于这些代表错误的变量需要给包外代码使用，所以其访问权限只能是公开的。
// 这就带来了一个问题，如果有恶意代码改变了这些公开变量的值，那么程序的功能就必然会受到影响。
// 因为在这种情况下我们往往会通过判等操作来判断拿到的错误值具体是哪一个错误，如果这些公开变量的值被改变了，那么相应的判等操作的结果也会随之改变。
// 两个解决方案：
//   - 私有化此类变量，也就是说，让它们的名称首字母变成小写，然后编写公开的用于获取错误值以及用于判等错误值的函数。
//     比如，对于错误值os.ErrClosed，先改写它的名称，让其变成os.errClosed，然后再编写ErrClosed函数和IsErrClosed函数。
//   - 使用errors.Is(err, target error)
//     当err是可比较的，或err所属类型提供了: Is(target error) bool 接口时时，就可以使用。
//     例如 syscall.Errno 类型，
//        type Errno uintptr                           #uintptr的重定义类型，可以被赋值给常量
//        func (e Errno) Error() string {...}          #Error接口
//        func (eErrno) Is(target error) bool {...}    #Is接口
//        const EPIPE = ...                            #常量无法被修改
//        examples,
//          _,_,err := syscall.Syscall(...)
//          if errors.Is(err, syscall.EPIPE) {...}