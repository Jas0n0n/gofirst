package main

import "fmt"

func serveralInts(numbers ...int)  {
	fmt.Println(numbers)
}

func mix(num int,flag bool,strings ...string )  {
   fmt.Println(num,flag,strings)
}

func main()  {
	intSlice :=[]int{1,2,3}
	serveralInts(intSlice...)

	stringslice :=[]string{"a","b","c"}
	mix(1,true,stringslice...)
}