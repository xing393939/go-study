<link rel="stylesheet" href="../extra/ideal-image-slider.css">
<link rel="stylesheet" href="../extra/ideal-default-theme.css">
<script src="../extra/ideal-image-slider.js"></script>
<script src="../extra/ideal-iis-bullet-nav.js"></script>
<script>
var gitbook = gitbook || [];
gitbook.push(function() {
    let slider = new IdealImageSlider.Slider('.IdealImageSlider');
    slider.addBulletNav();
})
</script>

### GC 分析

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

