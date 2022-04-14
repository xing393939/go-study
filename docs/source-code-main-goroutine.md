### 启动main协程

<link rel="stylesheet" type="text/css" href="../images/jquery.dialog.css">
<script type=text/javascript src="../images/jquery.dialog-code.js"></script>

#### Go的启动过程
```
runtime/rt0_linux_amd64.s的_rt0_amd64_linux()   JMP    _rt0_amd64(SB)
runtime/asm_amd64.s的_rt0_amd64()               JMP    runtime·rt0_go(SB)
runtime/asm_amd64.s的rt0_go()
```

#### runtime.rt0_go
<div class="DialogCode" data-code="rt0_go"></div>



