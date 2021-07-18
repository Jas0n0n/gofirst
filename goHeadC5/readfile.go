package datafile

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main()  {

	//打开文件
	file,err := os.Open("data.txt")

	//读取失败后输出错误信息
	if err !=nil{
		log.Fatal(err)
	}

	//读取文件流
    scanner :=bufio.NewScanner(file)

    //从文件中读取一行
    //循环到文件结尾，scnner.Scan返回false
	for scanner.Scan() {
        fmt.Println(scanner.Text())
	}

	//关闭文件流
	err = file.Close()

	//如果文件关闭时异常则输出
    if err !=nil{
    	log.Fatal(err)
	}

	//如果扫描文件时出现错误则退出
	if scanner.Err() !=nil{
		log.Fatal(scanner.Err())
	}

}