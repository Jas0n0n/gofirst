package main

import "fmt"

func SayHi(){
	for x:=1;x<3;x++{
		fmt.Println("hi")
	}
}



func manyReturns() (int,bool,string) {

	return 1,true,"hi"
}

func main()  {
	myInt,myBool,myString := manyReturns()
	fmt.Println(myInt,myBool,myString)
}

