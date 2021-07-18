package main

import (
	"fmt"
	"math"
)

func squareRoot(number float64) (float64,error)  {
	if number<0{
        return 0,fmt.Errorf("数字应大于0")
	}
	return math.Sqrt(number),nil
}

func main()  {
	root,err := squareRoot(10)
	if err != nil{
		fmt.Println(err)
	}else {
		fmt.Printf("%0.3f",root)
	}


}