
//4-4 GOPATH以及目录结构

// GOPATH环境变量
// 默认在~/go(unix, linux), %USERPROEILE%\go(windows)
// 官方推荐，所有项目和第三方库都放在同一个GOPATH下
// 也可以将每个项目放在不同的GOPATH

// 在liunx中查看GOPATH
// $ echo $GOPATH
// $ cat ~/.base_profile
// 可以在上述文件中增加 
// export GOPATH = /Users/nancy/学习/go/imooc
// PATH 增加多 "$GOPATH/bin:"

// 一般import包，会默认在GOROOT或者GOPATH目录下寻找

// 一般我们要下载go的包，可以使用go get命令
// $ go get goland.org/x/tools/cmd/goimports
// 但由于网络问题，可能拉取不下来
// 可以用gopm来获取无法下载的包
// 需要先下载gopm工具
// $ go get -v github.com/gpmgo/gopm
// 拉取下github.com/gpmgo/gopm后，在GOPATH的src中就会出现github.com目录
// 在GOPATH的bin目录下，就有gopm
// 下载完gopm后，就可以使用gopm来下载包了
// $ gopm get -g -v -u goland.org/x/tools/cmd/goimports
// 下载完goland.org/x/tools/cmd/goimports就可以在GOPATH下src下看到
// 但在bin中并没有goimports， 需要build gopm
// 可以使用
// $ go build goland.org/x/tools/cmd/goimports //编译
// $ go install goland.org/x/tools/cmd/goimports //产生pkg文件和可执行文件
// 安装成功就可以在GOPATH下的bin文件看到
// $ go run //编译执行

// GOPATH下目录结构
// src 
//     git repository 1
//     git repository 2
// pkg
//     git repository 1
//     git repository 2
// bin
//     执行文件1， 2， 3

// 每个文件夹都只能有一个main，所以在编译的时候，如果想每个文件都带着main函数，需要将每个文件放到一个文件夹里
