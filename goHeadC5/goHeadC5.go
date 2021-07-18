package datafile

import "fmt"

func main()  {
	var myfirstArray [3]int
	myfirstArray[0] = 1


	var mysecondArray = [3]string{"a","b","c"}


	mythirdString :=[3]int{1,2,3}

	myfourthString :=[2]string{
		"Love is the best gift you can give",
		"Do the best",
	}


	notes := [7]string{"do","ri","mi","fa","sa","la","si"}

	for i:=0;i<len(notes);i++{
		fmt.Println(notes[i])
	}

    for index,note := range notes{
    	fmt.Println(index,note)
	}

	//#不打印index值，用空白标识符
	for _,note2 :=range notes{
		fmt.Println(note2)
	}


	//sum求数组的和

	numbers :=[3]float64{1.1,2.2,3.3}

	var sum float64 =0;

	for _,num:= range numbers{
		sum +=num
	}
	fmt.Println(sum)






	fmt.Printf("%#v \n",myfourthString)


	fmt.Println(myfourthString)
	fmt.Println(mythirdString)
	fmt.Println(mysecondArray)
	fmt.Println(myfirstArray[2])

}
