package main

type Myintinterface interface {
	fun()
}

type Myint int

func (m Myint) fun() {}

//go:nosplit
func test() {
	m := Myint(1) // CALL "".Myint.fun(SB)
	m.fun()

	m2 := &m
	m2.fun() // 还是CALL "".Myint.fun(SB)
}

//go:nosplit
func test2() {
	/*
	 * LEAQ	go.itab."".Myint,"".Myintinterface(SB), AX
	 * LEAQ	""..stmp_0(SB), AX                          // stmp_0=12
	 */
	var mii Myintinterface = Myint(12)
	mii.fun() // CALL "".Myint.fun(SB)

	/*
	 * LEAQ	go.itab.*"".Myint,"".Myintinterface(SB), CX
	 * LEAQ	""..autotmp_4+32(SP), AX                    // autotmp_4=0
	 */
	var mii2 Myintinterface = new(Myint)
	mii2.fun() // 还是CALL "".Myint.fun(SB)
}
