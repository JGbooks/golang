package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("测试命令行传参")
	arg := os.Args
	for i := 0; i < len(arg); i++ {
		fmt.Println(i, arg[i])
	}
}
