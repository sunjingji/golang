//@Title Gob序列化示例
package example

//# Gob
//  Gob设计的目标：
//    易于使用。因为go有反射，不需要额外的序列化接口，这一点很容易实现。
//    效率是非常重要的，文本表示，以XML和JSON为例，在通信网络中太慢了，二进制编码是必要的。
//    Gob流必须是自描述的。每一个gob流，从一开始就读取，包含足够的信息，可以由一个对其内容一无所知的代理解析整个流。这个属性意味着您将始终能够解码存储在文件中的gob流，即使您已经忘记了它所代表的数据。
//
//  Gob的设计规避了protobuf设计上的一些“缺陷”：
//    必须知道数据的定义(.proto定义)，才能解码数据。
//    protobuf只适用于struct结构，你没办法在顶层直接使用integer或者array，而只能把它们包裹在一个struct中。
//    "required"特性的支持，使得编解码更复杂，性能也受影响。
//    默认值的处理。
//
//  Gob的一些特性：
//    除非你设了值，默认值不编码，缺省就是相应类型的“零值”。
//    gob编码有点像go中的常量，对数值它只区分有符号无无符号，不考虑你是int8还是uint16类型，以最小长度编码。 指针也以类似数值的方式编码。
//    解码结构时，按照字段名称和类型匹配，例如，
//        编码时使用结构：type T struct{ X, Y, Z int }
//        解码时使用结构：type U struct{ X, Y *int8 }，那么Z字段就被丢掉了，这为gob提供了扩展数据格式的便利。
//
//    参考go官方博客: http://golang.org/blog/gobs-of-data