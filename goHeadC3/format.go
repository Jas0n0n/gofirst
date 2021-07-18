package main

import (
	"errors"
	"fmt"
)

func main()  {
	fmt.Printf("%12s | %s\n","Product","Cost in Cents")

	fmt.Println("-------------------------------")

	fmt.Printf("%12s | %4d\n","Stamps",5000)
	fmt.Printf("%12s | %4d\n","Paper Clipse",5)
	fmt.Printf("%12s | %4d\n","Tape",1)
	fmt.Printf("%12s | %4.2f\n","Pencil Case",6.666)
	err := errors.New("高度和宽度不能为负数")
	fmt.Printf(err.Error())
}
