### Go 语言项目开发实战

#### 参考资料
* [Go 语言项目开发实战](https://time.geekbang.org/column/intro/100079601?tab=catalog)

#### 第1站 规范设计
```
Commit Message 规范：
<type>[optional scope]: <description>

[optional body]

[optional footer]
```
* type说明了commit的类型，常用类型见[链接](../images/combat/commit-message.png)
* scope说明了commit的影响范围的，见[示例](https://github.com/marmotedu/iam/blob/master/docs/devel/zh-CN/scope.md)
* description是对commit的简短描述
* body是对commit的更详细的描述
* footer常用的有两种：
  * BREAKING CHANG: xxx。表示不兼容的改动
  * Closes #1, #2。关闭了2个issue
* 功能分支工作流，[三种合并代码方式的差别](https://www.chenshaowen.com/blog/the-difference-of-tree-ways-of-merging-code-in-github.html)
* Git Flow工作流，介绍见[链接](https://blog.csdn.net/weixin_46674610/article/details/115396404)

![img](../images/combat/pattern.png)

#### 第2站 基础功能

