package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getFloat() (float64, error) {

	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)

	grade, err := strconv.ParseFloat(input, 64)

	return grade, nil

}

func main() {

	fmt.Println("请输入华式摄氏度：")

	tem, err := getFloat()

	if err != nil {
		log.Fatal(err)
	}

	celsius := (tem - 32) * 5 / 9

	fmt.Printf("%2.2f degrees celsuis", celsius)

}
