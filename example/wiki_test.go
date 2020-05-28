//@Title Wiki
package example

// # 高效Go编程
//   https://golang.google.cn/doc/effective_go.html，总结了很多go语言高效编程的建议，
//
//
// # 识别Race
//   使用-race编译运行程序，可以识别运行时潜在的共享变量竞争访问。
//   参考go语言官方博客： http://golang.org/blog/race-detector
//
//  # ast
//    Json库的用法，参考官方博客：http://golang.org/blog/json-and-go
//
//
//  # time
//    v.Format("2006-01-02 15:05:05.999")
//
//  # sql
//    使用go语言database/sql库从数据库中读取null值的问题，以及如何向数据库中插入null值。本文在这里使用的是sql.NullString, sql.NullInt64, sql.NullFloat64等结构体
//
//
//  # excel
//    excel库："github.com/extrame/xls"
//    xlsx库： 360EntSecGroup-Skylar
//
//
//  # protobuf
//    protoc --go_out=plugins=grpc:. hello.proto
//
//
//  # 限制go程序使用的cpu核数: 环境变量GOMAXPROCS
//
//
//  # godoc
//    注释：与注释的函数、包开头用相同的单词   http://127.0.0.1/blog/godoc-documenting-go-code