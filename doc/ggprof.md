# Go 性能分析

## 1. 性能分析工具
Go 语言为程序开发者们提供了丰富的性能分析 API，和非常好用的标准工具。这些 API 主要存在于：

- runtime/pprof；
- net/http/pprof；
- runtime/trace；

这三个代码包中。

另外，runtime代码包中还包含了一些更底层的 API。它们可以被用来收集或输出 Go 程序运行过程中的一些关键指标，并帮助我们生成相应的概要文件以供后续分析时使用。

至于标准工具，主要有go tool pprof和go tool trace这两个。它们可以解析概要文件中的信息，并以人类易读的方式把这些信息展示出来。

此外，go test命令也可以在程序测试完成后生成概要文件。

如此一来，我们就可以很方便地使用前面那两个工具读取概要文件，并对被测程序的性能加以分析。这无疑会让程序性能测试的一手资料更加丰富，结果更加精确和可信。

在 Go 语言中，用于分析程序性能的概要文件有三种，分别是：

- CPU 概要文件（CPU Profile）
- 内存概要文件（Mem Profile）
- 阻塞概要文件（Block Profile）

pprof 工具使用，参考郜林老师的文档 [go tool pprof](https://github.com/hyper0x/go_command_tutorial/blob/master/0.12.md)

## 2. runtime/pprof
Go语言工具链中的 pprof 可以帮助开发者快速分析及定位各种性能问题，如 CPU 消耗、内存分配及阻塞分析。

性能分析首先需要使用 runtime.pprof 包嵌入到待分析程序的入口和结束处。runtime.pprof 包在运行时对程序进行每秒 100 次的采样，最少采样 1 秒。然后将生成的数据输出，让开发者写入文件或者其他媒介上进行分析。

pprof 工具链配合 Graphviz 图形化工具可以将 runtime.pprof 包生成的数据转换为 PDF 格式，以图片的方式展示程序的性能分析结果。

### 安装第三方图形化显式分析数据工具（Graphviz）
Graphviz 是一套通过文本描述的方法生成图形的工具包。描述文本的语言叫做 DOT。

在 www.graphviz.org 网站可以获取到最新的 Graphviz 各平台的安装包。

CentOS 下，可以使用 yum 指令直接安装：
    
    $ yum install graphiviz

### 安装第三方性能分析来分析代码包
runtime.pprof 提供基础的运行时分析的驱动，但是这套接口使用起来还不是太方便，例如：

- 输出数据使用 io.Writer 接口，虽然扩展性很强，但是对于实际使用不够方便，不支持写入文件。
- 默认配置项较为复杂。

很多第三方的包在系统包 runtime.pprof 的技术上进行便利性封装，让整个测试过程更为方便。这里使用 github.com/pkg/profile 包进行例子展示，使用下面代码安装这个包：

    $ go get github.com/pkg/profile

### CPU分析代码
下面代码故意制造了一个性能问题，同时使用 github.com/pkg/profile 包进行性能分析

基准测试代码如下：

	package main
	import (
    	"github.com/pkg/profile"
	    "time"
	)
	func joinSlice() []string {
    	var arr []string
	    for i := 0; i < 100000; i++ {
    	    // 故意造成多次的切片添加(append)操作, 由于每次操作可能会有内存重新分配和移动, 性能较低
    	    arr = append(arr, "arr")
	    }
    	return arr
	}
	func main() {
    	// 开始性能分析, 返回一个停止接口
	    stopper := profile.Start(profile.CPUProfile, profile.ProfilePath("."))
    	// 在main()结束时停止性能分析
	    defer stopper.Stop()
    	// 分析的核心逻辑
	    joinSlice()
    	// 让程序至少运行1秒
	    time.Sleep(time.Second)
	}

性能分析需要可执行配合才能生成分析结果，因此使用命令行对程序进行编译，代码如下：

	$ go build -o cpu cpu.go
	$ ./cpu
	$ go tool pprof --pdf cpu cpu.pprof > cpu.pdf

这个过程中会调用 Graphviz 工具，Windows 下需将 Graphviz 的可执行目录添加到环境变量 PATH 中。

重新优化代码，在已知切片元素数量的情况下直接分配内存，代码如下：

	func joinSlice() []string {
	    const count = 100000
	    var arr []string = make([]string, count)
	    for i := 0; i < count; i++ {
	        arr[i] = "arr"
	    }
	    return arr
	}

重新运行上面的代码进行性能分析，最终得到的 cpu.pdf 中将不会再有耗时部分。

### 内存分析代码

    针对内存概要信息的采样会按照一定比例收集 Go 程序在运行期间的堆内存使用情况。
    设定内存概要信息采样频率的方法很简单，只要为runtime.MemProfileRate变量赋值即可。
    这个变量的含义是，平均每分配多少个字节，就对堆内存的使用情况进行一次采样。
    如果把该变量的值设为0，那么，Go 语言运行时系统就会完全停止对内存概要信息的采样。该变量的缺省值是512KB，也就是512千字节。

    在这之后，当我们想获取内存概要信息的时候，还需要调用runtime/pprof包中的WriteHeapProfile函数。该函数会把收集好的内存概要信
    息，写到我们指定的写入器中。注意，我们通过WriteHeapProfile函数得到的内存概要信息并不是实时的，它是一个快照，是在最近一次的内存
    垃圾收集工作完成时产生的。

      pprof.StopCPUProfile()
      doSomethig()
      pprof.WriteHeapProfile(w)

### 阻塞分析代码

    调用runtime包中的SetBlockProfileRate函数，即可对阻塞概要信息的采样频率进行设定。该函数有一个名叫rate的参数，它是int类型的。
    这个参数的含义是，只要发现一个阻塞事件的持续时间达到了多少个纳秒，就可以对其进行采样。

    当我们需要获取阻塞概要信息的时候，需要先调用runtime/pprof包中的Lookup函数并传入参数值"block"，从而得到一个*runtime/pprof.Profile类型的值（以下简称Profile值）。在这之后，我们还需要调用这个Profile值的WriteTo方法，以驱使它把概要信息写进我们指定的写入器中。
    
    这个WriteTo方法有两个参数，一个参数就是我们刚刚提到的写入器，它是io.Writer类型的。而另一个参数则是代表了概要信息详细程度的int类型参数debug，可以是0/1/2。

    Loopup的用法参考图： lookup.png。


## 3. net/http/pprof

如何为基于 HTTP 协议的网络服务添加性能分析接口？

我们在一般情况下只要在程序中导入net/http/pprof代码包就可以了，就像这样：

	import _ "net/http/pprof"

然后，启动网络服务并开始监听，比如：

	log.Println(http.ListenAndServe("localhost:8082", nil))

在运行这个程序之后，我们就可以通过在网络浏览器中访问http://localhost:8082/debug/pprof这个地址看到一个简约的网页。

我们可以通过go tool pprof工具直接读取这样的 HTTP 响应，例如：

	go tool pprof http://localhost:6060/debug/pprof/profile?seconds=60

除此之外，还有一个值得我们关注的路径，即：/debug/pprof/trace。在这个路径下，程序主要会利用runtime/trace代码包中的 API 来处理我们的请求。

前面说的这些 URL 路径都是固定不变的。这是默认情况下的访问规则。我们还可以对它们进行定制。