### M的管理

<link rel="stylesheet" type="text/css" href="../images/jquery.dialog.css">
<script type=text/javascript src="../images/jquery.dialog-code.js"></script>

#### M的创建
* [工作线程的唤醒及创建(19)](https://www.cnblogs.com/abozhang/p/10916701.html)
* 在ready()和newproc()时会调用wakep()
* wakep发现目前没有自旋的M，则调用startm来获取一个空闲的M或者新建一个M
* 新建的M将执行mstart，它要么执行G，要么休眠，永远不会退出

<div class="DialogCode" data-code="wakep"></div>

#### 辛勤工作的M
* [深入golang runtime的调度](https://zboya.github.io/post/go_scheduler/#m%E7%9A%84%E7%AE%A1%E7%90%86)
* M执行mstart会调用schedule。后面就是循环的调用schedule：
  * 调用globrunqget：每隔61次调度，从全局队列取G
  * 调用runqget：从P的本地队列取G
  * 调用findrunnable：
    * 调用runqget
    * 调用globrunqget
    * 从netpool取
    * 随机从其他P偷取G
    * 再次调用globrunqget
    * 检查有没有空闲的P且没有绑定M，与之绑定
    * 再次从netpool取
    * 实在找不到G，调用stopm休眠

<div class="DialogCode" data-code="findrunnable"></div>




