package main

import "fmt"

func sayHi(){
	for x:=1;x<3;x++{
		fmt.Println("hi")
	}
}

func repeatLine(line string,times int)  {
	for i:=1;i<=times;i++{
		fmt.Println(line)
	}
}

func main()  {
	sayHi()
	repeatLine("imsun2y",3)
}