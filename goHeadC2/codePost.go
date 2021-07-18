package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	//获取文件信息
	fileInfo, err := os.Stat("my.txt")

	//如果检测到err，打印出来
	if err != nil {
		log.Fatal(err)
	}

	//打印出文件的大小
	fmt.Println(fileInfo.Size())

}
