package main

// 解析
// 考点:defer和函数返回值
// 需要明确一点是defer需要在函数结束前执行。
// 函数返回值名字会在函数起始处被初始化为对应类型的零值并且
// 作用域为整个函数 DeferFunc1有函数返回值t作用域为整个函数，
// 在return之前defer会被执行，所以t会被修改，返回4;
// DeferFunc2函数中t的作用域为函数，返回1; DeferFunc3返回3

func main() {
	println(DeferFunc1(1))
	println(DeferFunc2(1))
	println(DeferFunc3(1))
}

func DeferFunc1(i int) (t int) {
	t = i
	defer func() {
		t += 3
	}()
	return t
}

func DeferFunc2(i int) int {
	t := i
	defer func() {
		t += 3
	}()
	return t
}

func DeferFunc3(i int) (t int) {
	defer func() {
		t += i
	}()
	return 2
}
