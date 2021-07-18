package main

import "fmt"

func negate(myboolean *bool)  {
	*myboolean  = !*myboolean
}

func main()  {
	truth := true
	negate(&truth)
   fmt.Println(truth)
}