// @Title 变量类型转换、断言示例
package example

import "fmt"

// 类型转换
func ExampleConversion() {
  // 当整数值的类型的有效范围由宽变窄时，会发生截断。
  var srcInt = int16(-255)
  dstInt := int8(srcInt)  // int16类型的值-255的补码是1111111100000001，截断后是00000001。
  fmt.Printf("#1:%d\n", dstInt)

  // 接把一个整数值转换为一个string类型的值是可行的，但被转换的整数值应该可以代表一个有效的 Unicode 代码点，否则转换的结果将会是"�"（仅由高亮的问号组成的字符串值）。
  fmt.Printf("#2:%s\n", string(-1))

  // 从string类型向[]byte类型转换时代表着以 UTF-8 编码的字符串会被拆分成零散、独立的字节。
  s := "你好"
  b := []byte(s)  // ['\xe4', '\xbd', '\xa0', '\xe5', '\xa5', '\xbd']
  fmt.Printf("#3:%v\n", b)

  // rune是int32的别名，主要用途是为了区分字符值和整型值。
  // 从string类型向[]rune类型转换时代表着字符串会被拆分成一个个 Unicode 字符。
  r := []rune(s)  // ['\u4F60', '\u597D']
  fmt.Printf("#4:%v\n", r)

  // 注意: len(string) 和 len([]byte) 返回的是字节数，len([]rune)返回的是rune的个数
  fmt.Printf("#5:%d,%d,%d\n", len(s),len(b),len(r))

  // Output:
  // #1:1
  // #2:�
  // #3:[228 189 160 229 165 189]
  // #4:[20320 22909]
  // #5:6,6,2
}

// 类型断言
func ExampleAssert() {
  m := make(map[string]interface{})
  m["key"] = "value"

  // 类型断言表达式的语法形式是 x.(T)，其中x代表要被判断类型的值；这个值必须是接口类型的(本身就是接口类型，或者先转为interface{})。
  // 在go语言中，interface{}代表孔接口，任何类型都是它的实现类型，都可以转为interface{}。
  // 类型断言要使用下面的标准方法，通过ok判断断言结果；如果省略掉ok，在类型判断为否时，会引发异常。
  if _, ok := interface{}(m).(map[string]interface{}); ok {
    fmt.Printf("true\n")
  }

  // switch 断言
  if v, ok := m["key"]; ok {
    switch v.(type) {
    case string:
      fmt.Printf("true\n")
    }
  }

  // Output:
  // true
  // true
}