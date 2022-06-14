### 页分配器

<link rel="stylesheet" type="text/css" href="../images/jquery.dialog.css">
<script type=text/javascript src="../images/jquery.dialog-code.js"></script>

#### mheap.alloc()
* 堆空间充足：调用`h.pages.alloc(npages)`
* 堆空间满了：调用`h.grow(npages)`

<div class="DialogCode" data-code="alloc"></div>