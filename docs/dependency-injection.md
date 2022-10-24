### 依赖注入

#### 参考
* [Fx: Golang中的依赖注入](http://ishare.site/post/2020/01/simple-dependency-injection-in-go/)
* [Golang依赖注入提升开发效率](https://mp.weixin.qq.com/s/Mj-EqwYWZBMr8XNIHxUqDA)
* [简洁代码之道（2）：避免全局可变状态](https://blog.xiaohansong.com/avoid-global-state.html)
* Go依赖注入的开源库
  * facebook的inject基于反射，运行时注入(功能有点弱，也不维护了)。
  * google的wire基于AST，编译期注入。
  * uber的dig基于反射，运行时注入。
  * https://github.com/uber-go/fx，在dig基础上包了一层, 用起来更爽
* 依赖注入的好处
  * 对象的创建和使用解耦（一般创建都交给了框架或者自己扩展的模块）。
  * 不用自己创建和组装，不用关心依赖创建顺序，根据使用顺序自动推导。

#### 举例说明
```
// 例1
func query() (email string) {
    db, err := sql.Open("postgres", "user=postgres dbname=test ...")
    if err != nil {
        panic(err)
    }
    err = db.QueryRow(`SELECT email FROM "user" WHERE id = $1`, 1).Scan(&email)
    if err != nil {
        panic(err)
    }
    return email
}

// 例2
func query(db *sql.DB) (email string) {
    err = db.QueryRow(`SELECT email FROM "user" WHERE id = $1`, 1).Scan(&email)
    if err != nil {
        panic(err)
    }
    return email
}

// 例3
func query() (email string) {
    err = globalDb.QueryRow(`SELECT email FROM "user" WHERE id = $1`, 1).Scan(&email)
    if err != nil {
        panic(err)
    }
    return email
}
```

* 上面的两个例子，query方法依赖sql.DB句柄，有三种方法做到
  * 例1的方法，由自己构造，缺点是耦合了db具体类；单元测试无法使用mockDB
  * 例2的方法，靠传参来注入
  * 例3的方法，使用全局变量，缺点是第三方包可能会修改全局变量
