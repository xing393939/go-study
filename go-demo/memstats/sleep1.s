# 编译运行：as -o a.s sleep1.s && ld -s -o a.out a.s
# 打印输出：echo $?
# 参数说明：调用号=rax，参数=rdi,rsi,rdx,r10,r8,r9
.text                # 代码段声明
.global _start       # 指定入口函数

_start:
    callq runtime_usleep
    movl $3,%ebx     # 参数一：退出代码
    movl $1,%eax     # 系统调用号(sys_exit)
    int  $0x80       # 调用内核功能

runtime_usleep:
    sub    $0x18,%rsp
    mov    %rbp,0x10(%rsp)
    lea    0x10(%rsp),%rbp
    mov    $0x0,%edx
    mov    $999000000,%eax # sleep 999秒
    mov    $0xf4240,%ecx
    div    %ecx
    mov    %rax,(%rsp)
    mov    $0x3e8,%eax
    mul    %edx
    mov    %rax,0x8(%rsp)
    mov    %rsp,%rdi
    mov    $0x0,%esi
    mov    $0x23,%eax
    syscall
    mov    0x10(%rsp),%rbp
    add    $0x18,%rsp
    retq
