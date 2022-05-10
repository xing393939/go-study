### Go 程序员面试

#### 参考说明
* [Go 程序员面试](https://golang.design/go-questions/)
* 基于Go 1.16.10

#### 第1章 逃逸分析
* 堆上对象都会通过调用runtime.newobject分配，该函数会调用runtime.mallocgc
* Go语言没有哪个关键字可以指定变量一定在堆上
* [查看Go支持的操作系统和CPU架构](https://golang.google.cn/doc/install/source)
* [内存分配之系统调用sbrk、brk、mmap、munmap](https://blog.csdn.net/Apollon_krj/article/details/54565768)



