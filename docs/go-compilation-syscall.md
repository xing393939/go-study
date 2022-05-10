### 汇编+Syscall

#### 开始工作
* 为了了解运行程序时的内存指标(VIRT、RES)，因而想运行一个干净的程序，于是想到了汇编

#### 编译运行一个简单的汇编
* [x86_64汇编之六：系统调用（system call）](https://blog.csdn.net/qq_29328443/article/details/107250889)
* [系统调用函数大全](https://filippo.io/linux-syscall-table/)
* 实现一个程序调用系统调用nanosleep和exit：go-demo/memstats/sleep1.s
* 查看VIRT、RES：top -p 进程ID
* 查看内存映射：pmap -p 进程ID

#### 如何快速写出Syscall的汇编
* 先写一个简单go程序并编译，里面已经包含了常用的Syscall代码
* 执行：objdump -d --no-show-raw-insn a.out | grep '<runtime.usleep>:' -A 9