package main

import "fmt"

func inRange(min float64, max float64, numbers ...float64) []float64 {

	var result []float64

	for _, number := range numbers {

		if number >= min && number <= max {
			result = append(result, number)
		}
	}
	return result
}

func main() {

	fmt.Println(inRange(11, 222, 333, 111, 66, 1, 8888))

}
