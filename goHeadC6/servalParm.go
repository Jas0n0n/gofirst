package main

import "fmt"

func mulityPara(strings ...string) {
	fmt.Println(strings)
}

func main() {
	mulityPara("a", "b", "c")
}
