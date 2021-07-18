package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	//提示用户输入成绩
	fmt.Println("请输入同学的成绩:")

	//从键盘读取文本的"缓冲读取器"
	reader := bufio.NewReader(os.Stdin)

	//使用空白标识符忽略错误返回值
	//返回用户输入的所有内容，直到按下<Enter>键为止
	//input,_ :=reader.ReadString('\n')
	input, err := reader.ReadString('\n')

	//对返回错误进行处理，返回错误存储在err中

	//从input中删除换行符
	input = strings.TrimSpace(input)
	//使用strconv包将ParseFloat函数转换为float64
	grade, err := strconv.ParseFloat(input, 64)

	//报告错误并停止程序运行
	if err != nil {
		log.Fatal(err)
	}

	//定义全局变量
	var status string

	//如果大于60分则通过
	if grade >= 60 {
		status = "passing"
	} else {
		status = "failed"
	}

	//输出成绩，A grade of is passing/failing
	fmt.Println("A grade of ", grade, "is", status)
}
