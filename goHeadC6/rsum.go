package main

import "fmt"

func inRanger(min float64,max float64,numbers ...float64) []float64  {

	var result []float64

	for _,number :=range numbers{
		if number >=min && number<=max{
			result = append(result, number)
		}
	}
	
	return result
}

func main()  {
	fmt.Println(inRanger(100,200,111,222,333))
}