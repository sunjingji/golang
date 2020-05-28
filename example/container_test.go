//@Title 容器示例
package example

// # 容器
//   go语言container包提供了三个容器类型：List(双向链表),Ring(循环链表),Heap(最小堆)。
//   go容器的实现思路是：定义接口, 实现基于接口的一组通用操作。
//   List,Ring很像c语言实现类似数据结构的做法，以interface{}代替了void*；
//   Heap需要你自己实现接口(提供基础的Sort、Swap等操作)，参考官方基于int[]的实现 heap_test.go)。
//   用c++类比，这种写法类似算法模板的实现方式：
//     namespace heap {
//       // 我们要实现的接口
//       struct Interface {
//         virtual int Len() = 0;
//         virtual bool Less(int i, int j) = 0;
//         virtual void Swap(int i, int j) = 0;
//         virtual void Push(Interface *x) = 0;  // add x as element Len()
//         virtual Interface* Pop() = 0;         // remove and return element Len() - 1.
//       };
//
//       // heap提供的方法
//       void Init(Interface *y);
//       void Push(Interface *h, Interface *x) {}
//       Interface* Pop(h Interface) {}
//       ...
//     }