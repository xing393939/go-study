### coredump

#### 生成core dump文件
* [用gdb分析coredump文件](https://blog.csdn.net/chengqiuming/article/details/90142476)
* sysctl -w kernel.core_pattern = core // 在当前进程目录下生成core文件
* ulimit -c unlimited                  // coredump文件大小是unlimited

#### 查看core dump文件
```
C语言：
gdb a.out core
(gdb) layout split           // 查看报错处的代码，gcc编译时加-g参数可以看到c代码

Go语言：
运行main进程需要：GOTRACEBACK=crash ./main
dlv core main core
(gdb) bt
```