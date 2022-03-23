package main

var (
	a int64 = 1
	b int64 = 1
)

// GOSSAFUNC=testStatement go build 可以看出b=a并不是原子语句
func testStatement() {
	b = a
}

// 产生竞争场景一：一个写，一个读
func scene1() {
	go func() {
		for {
			a++
		}
	}()
	go func() {
		for {
			b = a
		}
	}()
}

// 产生竞争场景一：两个写
func scene2() {
	go func() {
		for {
			a++
		}
	}()
	go func() {
		for {
			a--
		}
	}()
}

// 可见性的问题
func visibilityBad() {
	a = 2
	go func() {
		if a == 2 {
			// 写db
		}
	}()
}

// 可见性的解决
func visibilityGood(a int) {
	defer func() {
		println(1)
	}()
	if a > 1 {
		defer func() {
			println(2)
		}()
	}

	println(3)
}

func main() {
	visibilityGood(2)
	/*scene1()
	<-make(chan bool)*/
}
