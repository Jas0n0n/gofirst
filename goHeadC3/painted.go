package main

import (
	"fmt"
	"log"
)

var meterPerLitter float64 = 10.0

func paiteNeeded(width float64,height float64) (float64,error)  {

	if width <0 {
		return 0,fmt.Errorf("宽度 %0.2f 是不可用的",width)
	}
	if height <0 {
		return 0,fmt.Errorf("宽度 %0.2f 是不可用的",height)
	}

	area := width * height

	return area/meterPerLitter,nil
	
}

func main()  {

    amount,err := paiteNeeded(4.2,-11)
    if err !=nil {
		log.Fatal(err)
	}

    fmt.Printf("需要 %0.2f 升",amount)
}