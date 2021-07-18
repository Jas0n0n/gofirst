package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	//获取当前时间的整数形式
	seconds := time.Now().Unix()

	//给予不同的"种子"以生成不同的数值
	rand.Seed(seconds)

	// "math/rand" 生成0-99的随机数
	target := rand.Intn(100) + 1

	//提问
	fmt.Println("我设置了一个1-100的随机数")

	//从键盘读取输入流
	reader := bufio.NewReader(os.Stdin)

	//设置猜中状态，默认为false
	isguessed := false

	//guesses为设定的可猜测次数
	for guesses := 10; guesses > 0; guesses-- {

		//请尝试输入一个数
		fmt.Println("猜猜它是多少:")

		//读取字符串值，以enter键结尾
		input, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		//调用strings的去除空格方法
		input = strings.TrimSpace(input)

		//调用strcov的A to i 进行string 到 integer的转换
		guessNum, err := strconv.Atoi(input)

		//如果转换失败进行错误提示
		if err != nil {
			log.Fatal(err)
		}

		//如果猜到的数字大于目标值
		if guessNum > target {
			fmt.Println("不对,你猜的数字比它大")
			//猜错了，次数减1
			fmt.Println("你还剩", guesses-1, "次尝试机会")
		} else if guessNum < target {

			fmt.Println("不对,你猜的数字比它小")
			//猜错了，次数减1
			fmt.Println("你还剩", guesses-1, "次尝试机会")
		} else {
			isguessed = true
			fmt.Println("Yes,you got it!")
			//猜错了，次数减1
			fmt.Println("The number is ", guessNum)
			break
		}
	}
	//如果没有猜到，给出提示
	if !isguessed {
		fmt.Println("对不起，你没有猜中!", "它是：", target)
	}

}
