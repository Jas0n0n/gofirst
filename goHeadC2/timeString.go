package main

//导入函数包
import (
	"fmt"
	"strings"
)

func main() {

	//需要进行字符替换的对象
	broken := "G#,Let's r#cks!"

	//将设置一个strings.Replacer,将每个 "#" 替换为 "o"
	replacer := strings.NewReplacer("#", "o")

	//对string.replacer调用repalce方法，并传入需要替换的对象
	fixed := replacer.Replace(broken)

	//输出 Go,Let's rocks!
	fmt.Println(fixed)

	//
}
