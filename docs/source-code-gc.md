<link rel="stylesheet" href="../images/ideal-image-slider.css">
<link rel="stylesheet" href="../images/ideal-default-theme.css">
<script src="../images/ideal-image-slider.js"></script>
<script src="../images/ideal-iis-bullet-nav.js"></script>

### GC 流程图

#### GC 全流程
* [曹大-Go 语言的 GC](https://time.geekbang.org/column/article/484271)

![dlv](../images/gc-alltime.png)

#### 后台标记worker的工作过程

<div class="IdealImageSlider">
    <img src="../images/gc-mark/1649752192-1.jpg" />
    <img src="../images/gc-mark/1649752193-2.jpg" />
    <img src="../images/gc-mark/1649752193-3.jpg" />
    <img src="../images/gc-mark/1649752193-4.jpg" />
    <img src="../images/gc-mark/1649752194-5.jpg" />
    <img src="../images/gc-mark/1649752194-6.jpg" />
    <img src="../images/gc-mark/1649752194-7.jpg" />
    <img src="../images/gc-mark/1649752195-8.jpg" />
    <img src="../images/gc-mark/1649752195-9.jpg" />
</div>

#### gcDrain.scanobject的具体过程

<div class="IdealImageSlider">
    <img src="../images/gc-scanobject/1649753414-1.jpg" />
    <img src="../images/gc-scanobject/1649753414-2.jpg" />
    <img src="../images/gc-scanobject/1649753415-3.jpg" />
    <img src="../images/gc-scanobject/1649753415-4.jpg" />
</div>

#### 插入写屏障对象丢失问题

<div class="IdealImageSlider">
    <img src="../images/insert-barrier/1649748070-1.jpg" />
    <img src="../images/insert-barrier/1649748071-2.jpg" />
    <img src="../images/insert-barrier/1649748071-3.jpg" />
    <img src="../images/insert-barrier/1649748072-4.jpg" />
</div>

#### 删除写屏障对象丢失问题

<div class="IdealImageSlider">
    <img src="../images/delete-barrier/1649748868-1.jpg" />
    <img src="../images/delete-barrier/1649748869-2.jpg" />
    <img src="../images/delete-barrier/1649748869-3.jpg" />
    <img src="../images/delete-barrier/1649748870-4.jpg" />
    <img src="../images/delete-barrier/1649748870-5.jpg" />
</div>

#### write barrier 的执行过程

<div class="IdealImageSlider">
    <img src="../images/write-barrier/1649746680-1.jpg" />
    <img src="../images/write-barrier/1649746681-2.jpg" />
    <img src="../images/write-barrier/1649746681-3.jpg" />
</div>

#### markWorker和mutator同时工作

<div class="IdealImageSlider">
    <img src="../images/gc-mark-mutator/1649754644-1.jpg" />
    <img src="../images/gc-mark-mutator/1649754644-2.jpg" />
    <img src="../images/gc-mark-mutator/1649754644-3.jpg" />
    <img src="../images/gc-mark-mutator/1649754645-4.jpg" />
    <img src="../images/gc-mark-mutator/1649754645-5.jpg" />
    <img src="../images/gc-mark-mutator/1649754646-6.jpg" />
</div>

#### 扫描栈

<div class="IdealImageSlider">
    <img src="../images/Go_GC/Go_GC_page-0001.jpg" />
    <img src="../images/Go_GC/Go_GC_page-0002.jpg" />
    <img src="../images/Go_GC/Go_GC_page-0003.jpg" />
    <img src="../images/Go_GC/Go_GC_page-0004.jpg" />
    <img src="../images/Go_GC/Go_GC_page-0005.jpg" />
    <img src="../images/Go_GC/Go_GC_page-0006.jpg" />
    <img src="../images/Go_GC/Go_GC_page-0007.jpg" />
    <img src="../images/Go_GC/Go_GC_page-0008.jpg" />
    <img src="../images/Go_GC/Go_GC_page-0009.jpg" />
    <img src="../images/Go_GC/Go_GC_page-0010.jpg" />
    <img src="../images/Go_GC/Go_GC_page-0011.jpg" />
    <img src="../images/Go_GC/Go_GC_page-0012.jpg" />
    <img src="../images/Go_GC/Go_GC_page-0013.jpg" />
    <img src="../images/Go_GC/Go_GC_page-0014.jpg" />
    <img src="../images/Go_GC/Go_GC_page-0015.jpg" />
    <img src="../images/Go_GC/Go_GC_page-0016.jpg" />
    <img src="../images/Go_GC/Go_GC_page-0017.jpg" />
    <img src="../images/Go_GC/Go_GC_page-0018.jpg" />
    <img src="../images/Go_GC/Go_GC_page-0019.jpg" />
    <img src="../images/Go_GC/Go_GC_page-0020.jpg" />
    <img src="../images/Go_GC/Go_GC_page-0021.jpg" />
    <img src="../images/Go_GC/Go_GC_page-0022.jpg" />
    <img src="../images/Go_GC/Go_GC_page-0023.jpg" />
    <img src="../images/Go_GC/Go_GC_page-0024.jpg" />
    <img src="../images/Go_GC/Go_GC_page-0025.jpg" />
    <img src="../images/Go_GC/Go_GC_page-0026.jpg" />
    <img src="../images/Go_GC/Go_GC_page-0027.jpg" />
    <img src="../images/Go_GC/Go_GC_page-0028.jpg" />
    <img src="../images/Go_GC/Go_GC_page-0029.jpg" />
    <img src="../images/Go_GC/Go_GC_page-0030.jpg" />
    <img src="../images/Go_GC/Go_GC_page-0031.jpg" />
</div>