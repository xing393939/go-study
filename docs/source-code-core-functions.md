### 核心函数

<link rel="stylesheet" type="text/css" href="../images/jquery.dialog.css">
<script type=text/javascript src="../images/jquery.main.js"></script>

#### runtime.get_tls(CX)
* [本地线程存储](https://tiancaiamao.gitbooks.io/go-internals/content/zh/04.1.html#%E6%9C%AC%E5%9C%B0%E7%BA%BF%E7%A8%8B%E5%AD%98%E5%82%A8)
* 不同的goroutine调用get_tls，得到的就是当前的g

#### [runtime.mcall(fn(*g))](https://github.com/golang/go/blob/go1.16.10/src/runtime/asm_amd64.s#L302)
* [几个重要的汇编](https://hushi55.github.io/2017/05/10/Golang-function-call#menuIndex3)
* mcall只能在g栈被调用，不能在g0、gsingal栈调用
* fn永不返回，一般情况下由调用schedule结束，这样可以让m运行其他goroutine
* 也就是说mcall使得当前goroutine放弃当前的运行

#### [runtime.systemstack(fn)](https://github.com/golang/go/blob/go1.16.10/src/runtime/asm_amd64.s#L342)
* [几个重要的汇编](https://hushi55.github.io/2017/05/10/Golang-function-call#menuIndex3)
* 通过get_tls得到当前的g和m
* 检查g==m.gsignal，直接执行fn
* 检查g==m.g0，直接执行fn
* 检查g!=m.currg，报错
* 切换到m.g0的栈上执行fn，执行结束再切回来

#### [runtime.gogo(*gobuf)](https://github.com/golang/go/blob/go1.16.10/src/runtime/asm_amd64.s#L281)
* [调度循环](https://draveness.me/golang/docs/part3-runtime/ch06-concurrency/golang-goroutine/#655-%e8%b0%83%e5%ba%a6%e5%be%aa%e7%8e%af)
* 从m.g0栈切换到g栈，注意最后一条指令是JMP而不是RET
* 此时的SP=runtime.goexit，这样在执行完gobuf.pc函数最后RET的时候就会到runtime.goexit

#### [runtime.newproc(siz, fn)](https://github.com/golang/go/blob/go1.16.10/src/runtime/proc.go#L4018)
* siz=参数的大小、fn=函数的指针
* [goroutine的生老病死](https://tiancaiamao.gitbooks.io/go-internals/content/zh/05.2.html)
* 在systemstack上执行newg := [newproc1(fn, argp, siz, gp, pc)](https://github.com/golang/go/blob/go1.16.10/src/runtime/proc.go#L4043)
  * fn=函数的指针、argp=参数的指针、siz=参数的大小、gp=caller的g、pc=caller的pc
  * newproc1()主要工作是：
    * 通过[gfget(\_p_)](https://github.com/golang/go/blob/go1.16.10/src/runtime/proc.go#L4215)获取g，获取不到就[malg(stackSize)](https://github.com/golang/go/blob/go1.16.10/src/runtime/proc.go#L3987)
    * 将传入的参数移到Goroutine的栈上；
    * 更新Goroutine调度相关的属性
* 在systemstack上执行[runqput(_p_, newg, true)](https://github.com/golang/go/blob/go1.16.10/src/runtime/proc.go#L5789)
  * \_p_=caller的P、newg=新的g、true=把新的g放在P.runnext
  * 如果P的队列长度<P.runq(256)就加入P的队列，否则加入全局队列
* 在systemstack上执行[wakep()](https://github.com/golang/go/blob/go1.16.10/src/runtime/proc.go#L2469:6)

