// @Title 变量定义、重命名、作用域
package example

import (
  "fmt"
)

// 变量重命名
func ExampleRename() {
  // 变量声明
  var s1 string
  s1 = "s1"

  // 变量声明同时赋值, 类型推断
  var s2 = "s2"

  // 短变量声明, 类型推断
  s3 := 10

  // 变量重声明允许我们在使用短变量声明时不用理会被赋值的多个变量中是否包含旧变量。
  // 注1: 变量的类型在其初始化时就已经确定了，所以对它的再次声明赋予的类型必须与其原本的类型相同。
  // 注2：被“声明并赋值”的变量必须是多个，并且至少有一个新的变量。
  s3, s4 := sort(1,2)  // s3是旧变量，s4是新变量

  // 没有声明新变量，直接使用 "=" 赋值。
  s3, _ = sort(1,2)

  fmt.Printf("%s,%s,%d,%d", s1,s2,s3,s4)

  // Output: s1,s2,1,2
}

// 变量作用域
func ExampleScope() {
  s := "s1"

  // 变量的重声明只可能发生在一个作用域内，内层作用域的变量会覆盖外层代码中的变量。
  {
    s := "s2"
    fmt.Printf("%s\n", s)  // 内层作用域，定义了一个新变量
  }

  fmt.Printf("%s\n", s)

  // Output:
  // s2
  // s1
}

func sort(x int, y int) (lt int, lg int) {
  if x <= y {
    return x,y
  } else {
    return y,x
  }
}