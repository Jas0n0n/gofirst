package main

import "fmt"

func average3(numbers ...float64) float64 {

	//此处的sum为定义的初始值
	var sum float64 =0

	for _,number := range numbers{
		sum += number
	}
	return sum/float64(len(numbers))
}

func main()  {
	fmt.Println(average3(11,22,33,44,55,66))
}

