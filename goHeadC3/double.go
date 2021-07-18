package main

import "fmt"

func double(number *int)  {
	*number *=2
}

func main()  {
	myInt :=2
	double(&myInt)
	fmt.Println(myInt)
}