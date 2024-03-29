package main

import "fmt"

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret) //20 1, 2, 3
	return ret                    //3
}

func main() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b)) //a = 0
	defer calc("2", a, calc("20", a, b)) //b = 1
}
