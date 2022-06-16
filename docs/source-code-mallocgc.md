### 对象分配器

<link rel="stylesheet" type="text/css" href="../images/jquery.dialog.css">
<script type=text/javascript src="../images/jquery.dialog-code.js"></script>

#### mallocgc
* 微对象分配（0~16B，noscan）
* 小对象分配（16B~32KB）：依次从mcache、mcentral、mheap取span
* 大对象分配（大于32KB）：直接从mheap取span
* [三类对象分配图解](https://speakerdeck.com/deepu105/go-memory-allocation)

<div class="DialogCode" data-code="mallocgc"></div>
