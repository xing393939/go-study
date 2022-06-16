### 编译条件

#### 参考资料
* [你一直在寻找的GO语言条件编译](https://linkscue.com/posts/2018-06-18-golang-build-constraints/)
* [Build constraints](https://pkg.go.dev/cmd/go#hdr-Build_constraints)

#### 常用规则
```
1.编译标签
// +build windows 相当于go build -tags GOOS
//go:build !cgo   若CGO_ENABLED=0则编译

2.文件后缀
mypkg_linux.go         只在 linux 系统编译
mypkg_windows_amd64.go 只在 windows amd64 平台编译

3.指定配置方法一：./myServer release 
serverMode := os.Args[1] 

4.指定配置方法二：export SERVER_MODE=debug
os.Getenv("SERVER_MODE")

5.指定配置方法三：go build -ldflags '-X main.mode=debug' main.go
var mode string && print(mode)

6.指定配置方法四：go build -tags debug -o main
config_release.go //+build release
config_debug.go   //+build debug
```