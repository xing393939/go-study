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

#### write barrier 的执行过程

<div class="IdealImageSlider">
    <img src="../images/write-barrier/1649746680-1.jpg" />
    <img src="../images/write-barrier/1649746680-2.jpg" />
    <img src="../images/write-barrier/1649746680-3.jpg" />
</div>
