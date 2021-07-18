package main

import "fmt"

func main() {

	var myslice []string
	myslice = make([]string, 3)

	myslice[0] = "one"
	myslice[1] = "two"
	myslice[2] = "three"

	fmt.Println(myslice)
	for _, myslicenum := range myslice {
		fmt.Println(myslicenum)
	}

	//使用make初始化一个切片
	notes := make([]int, 3)

	notes[0] = 1
	notes[1] = 2
	notes[2] = 3

	fmt.Println(notes)
	for _, note := range notes {
		fmt.Println(note)
	}

	underlinearray := [5]string{"a", "b", "c", "d", "e"}

	//0到3之间的元素，尾部开区间
	slice1 := underlinearray[0:3]

	slice2 := underlinearray[:5]

	slice3 := underlinearray[1:]

	slicechange := underlinearray[1:]
	slicechange[0] = "B"

	//append在切片后追加
	slice5 := []string{"a", "b"}
	fmt.Println(slice5, len(slice5))

	slice5 = append(slice5, "c")
	fmt.Println(slice5, len(slice5))

	//append s4的变更影响s3,不影响s2
	s1 := []string{"s1", "s1"}
	s2 := append(s1, "s2", "s2")
	s3 := append(s2, "s3", "s3")
	s4 := append(s3, "s4", "s4")

	fmt.Println(s1, s2, s3, s4)

	s4[0] = "X"
	fmt.Println(s1, s2, s3, s4)

	//如果没有赋值，则默认取零值；slice自身的零值是nil
	intslice := make([]int, 10)
	var stringslice []string
	fmt.Printf("intslice: %#v,strslice: %#v", intslice, stringslice)

	var zeroslice []string
	if len(zeroslice) == 0 {
		zeroslice = append(zeroslice, "first element")

	}
	fmt.Printf("%#v", zeroslice)

	fmt.Println(underlinearray)
	fmt.Println(slice3)
	fmt.Println(slice2)
	fmt.Println(slice1)

	//practice
	array1 := [5]string{"a", "b", "c", "d", "e"}

	slice6 := array1[1:3]

	slice6 = append(slice6, "x")

	slice6 = append(slice6, "y", "z")

	for _, letter := range slice6 {
		fmt.Println(letter)
	}

}
