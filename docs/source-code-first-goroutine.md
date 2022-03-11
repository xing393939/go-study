### 启动第一个协程

<link rel="stylesheet" type="text/css" href="../images/jquery.dialog.css">
<script>
window.onload = function() {
$.getScript("../images/jquery.dialog.js", function() {
    $(".dialog1").dialog({
        'title':'警告'
    },function(api){
        $('#btn1').click(function(){
            api.open();
        });
    });
});
}
</script>

<button id="btn1" class="button">打开对话框</button>
<div class="dialog dialog1">你确定我够帅吗？</div>


