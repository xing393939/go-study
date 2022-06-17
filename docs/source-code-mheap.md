### 页分配器

<link rel="stylesheet" type="text/css" href="../images/jquery.dialog.css">
<script type=text/javascript src="../images/jquery.dialog-code.js"></script>

#### mheap.alloc()
* 堆空间充足：调用`h.pages.alloc(npages)`
* 堆空间满了：调用`h.grow(npages)`

#### 几个数学题
* [堆内存分配：mallocgc函数 - B站视频](https://www.bilibili.com/video/BV1gT4y1o7H1)
* 题1：已知一个堆内存p，求p在第几个arena？
  * arena编号 = (p - arenaBaseOffset) / heapArenaBytes
* 题2：在linux_amd64下，arena大小和对齐边界都是64M，线性地址有48b，求可划分多少arena？
  * arena总数 = 2 ^ 48 / 2 ^ 26 = 4M
* 题3：已知一个堆内存p，求p在第几个page？
  * page编号 = (p / pageSize) % pagesPerArena
* 题4：已知堆内存的page编号，求对应的mspan？
  * mspan = heapArena.spans\[page编号]

<div class="DialogCode" data-code="alloc"></div>