package main

import "fmt"

func main()  {
	var myInt int =11
	var myIntPointer *int
	myIntPointer = &myInt
	*myIntPointer =20
	fmt.Println(*myIntPointer)
	fmt.Println(myInt)

}