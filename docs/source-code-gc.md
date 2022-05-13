<link rel="stylesheet" href="../images/ideal-image-slider.css">
<link rel="stylesheet" href="../images/ideal-default-theme.css">
<script src="../images/ideal-image-slider.js"></script>
<script src="../images/ideal-iis-bullet-nav.js"></script>

### GC 分析

#### GC 全流程
* [曹大-Go 语言的 GC](https://time.geekbang.org/column/article/484271)

![dlv](../images/gc-alltime.jpg)

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