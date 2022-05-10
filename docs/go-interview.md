### Go 程序员面试

#### 参考说明
* [Go 程序员面试](https://golang.design/go-questions/)
* 基于Go 1.16.10

#### 第1章 逃逸分析
* 堆上对象都会通过调用runtime.newobject分配，该函数会调用runtime.mallocgc
* Go语言没有哪个关键字可以指定变量一定在堆上
* [查看Go支持的操作系统和CPU架构](https://golang.google.cn/doc/install/source)
* [Linux内存管理：释放大块内存时的阻塞问题](http://hanyunqi.me/2020/05/17/Linux%E5%86%85%E5%AD%98%E7%AE%A1%E7%90%86%EF%BC%9A%E9%87%8A%E6%94%BE%E5%A4%A7%E5%9D%97%E5%86%85%E5%AD%98%E6%97%B6%E7%9A%84%E9%98%BB%E5%A1%9E%E9%97%AE%E9%A2%98/)
  * C标准库malloc，小于128KB用brk，大于128KB用mmap
  * 测试发现munmap一块20GB的内存，会阻塞其他线程的malloc（brk/sbrk/mmap）1秒左右。
  * 页框回收算法（PFRA）：大部分内核占用的内存不可回收，大部分用户的可以回收，被mlock标记的不可回收
  * 测试释放大块内存的阻塞：munmap、madvise + munmap、mlock + madvise + munmap
* [golang内存泄漏排查](https://lvbay.github.io/2020/01/20/golang%E5%86%85%E5%AD%98%E6%B3%84%E6%BC%8F%E6%8E%92%E6%9F%A5/)
  * runtime.Memstats的解读见下图
* [Go 应用内存占用太多，让排查？](https://eddycjy.gitbook.io/golang/di-1-ke-za-tan/why-vsz-large)  
  * 分析了runtime.schedinit->runtime.mallocinit
* [top 命令的虚拟内存（VIRT）、物理内存（RES）、共享内存（SHR）](http://xuwang.online/index.php/archives/253/)
  * 匿名内存：就是没有文件背景的内存，就是无法和磁盘进行交换的，比如堆栈。
  * 共享内存（SHR）包括：
    * 程序的代码段。
    * 动态库的代码段。
    * 通过 mmap 做的文件映射。
    * 通过 mmap 做的匿名映射，但指明了MAP_SHARED属性。
    * 通过 shmget 申请的共享内存。

![Pointer](../images/interview/runtime.Memstats.png)

