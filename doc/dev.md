# 开发环境部署

## Go安装包
开发环境是Windows+Linux，安装最新版版本即可。

    Windows版本：go1.14.3.windows-amd64.msi
    Linux版本  ：go1.14.3.linux-amd64.tar.gz

Go安装需要配置`GOROOT`和`GOPATH`两个环境变量，虽然包管理已经不再需要`GOPATH`，但是`mod`下载的包会存放在`$GOPATH/pkg`目录。

## MOD
Go现在支持`mod`管理项目的依赖包，非常好用，为了启用`mod`管理，可能需要设置下环境变量。

    # go mod功能开关，默认是auto，在gopath中不启用
    # 可设置为on强制启用
    export GO111MODULE=on

在项目中，如果需要使用已经下载到本地的库，或者干脆就是项目内部的库，可以在`mod`文件中使用`replace`设置。

    replace (
      github.com/gohouse/goroom => /path/to/go/src/github.com/gohouse/goroom
      by => ../by
    )

## GOPROXY
我们在开发过程中，经常会用到Go语言官方包，例如golang.org/x/...，国内网络环境无法访问。

Go镜像库托管在github.com/golang/上，可以从github下载，然后再移动到golang.org目录。

好在Go get支持代理服务器，在V1.13之后通过设置GOPROXY可以修改代理服务器，可以使用国内的代理服务器，安装包就很容易了。

    # 通过环境变量GOPROXY设置代理
    export GOPROXY=https://goproxy.io

    # 国内推荐使用 七牛代理
    export GOPROXY=https://github.com/goproxy/goproxy.cn

## 自定义代码包远程导入路径
如果希望在代码中不在出现github.com等托管地址，或者发布自己打包(大致个人品牌，使用自己的域名)，让使用者的项目与代码托管网站隔离所作出的努力，可以尝试用自定义远程包导入路径。

    # go get
	https://github.com/hyper0x/go_command_tutorial/blob/master/0.3.md

    实质上是提供一个web接口，响应go get的请求，做了一个重定向。 具体做法是在响应的html头中加入下面格式的元数据：
        <meta name="go-import" content="import-prefix vcs repo-root">

## 实用工具
Go安装包默认不包含`godoc`等实用工具，如果需要，需要自行安装。

    # 安装godoc工具
    go get -v golang.org/x/tools/cmd/godoc

`godoc`是一个非常好用的工具，可以自动扫描库和项目包，自动生成在线文件

    godoc -http=:80

## Golang IDE
Go开发的不二之选 - Golang，建议安装2020.1版本，对`mod`支持比较完善。

    # 官方安装包及破解包
    goland-2020.1.1-Crack-MacOS.zip:   https://089u.com/file/19978301-439708594
    goland-2020.1.1-Crack-Windows.zip: https://089u.com/file/19978301-439713520
    goland-2020.1.1-MacOS.zip:         https://089u.com/file/19978301-439708723
    goland-2020.1.1-Windows.zip:       https://089u.com/file/19978301-439709874

在项目中开启`mod`支持，File->Settings->Go->Go Modules选项卡 选择`Enable Go Modules Integration`。

## 工程管理
工程管理的一些实践：

	# 每个目录下只能有一个 package。

    # package 名称最好与目录名称一致(main包除外)，否则很容易迷惑包的用户。
      我们在构建或者安装这个代码包的时候，提供给go命令的路径应该是目录的相对路径，就像这样：
         go install puzzlers/article3/q2/lib  #假如package的名称是lib5，我们故意把包名和目录名称搞成不一样
      该命令会成功完成，当前工作区的 pkg 子目录下会产生相应的归档文件，具体的相对路径是:pkg/darwin_amd64/puzzlers/article3/q2/lib.a
      当我们实际使用该包的程序实体时，需要用的却是包的名称，例如: lib5.Hello()。

    # package 当我们不希望 package 当中的实体被外部引用时，最简单的办法是首字母小写。 还有一个办法是创建一个名称为"internal"的
      包，这个包智能在模块内部使用。