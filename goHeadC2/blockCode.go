package main

import "fmt"

var a = "a"

func main() {
	a = "a"
	b := "b"
	if true {
		c := "c"
		if true {
			d := "d"
			fmt.Println(a)
			fmt.Println(b)
			fmt.Println(c)
			fmt.Println(d)
		}

		fmt.Println(a)
		fmt.Println(b)
		fmt.Println(c)
		//变量d不再作用域内
		//fmt.Println(d)

	}
	fmt.Println(a)
	fmt.Println(b)
	//变量c再作用域内
	//fmt.Println(c)
	//变量d不再作用域内
	//fmt.Println(d)
}
