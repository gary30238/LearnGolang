package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 打開文件
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("open file error", err)
		return
	}

	// 關閉文件
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')

		fmt.Print(line)
		if err != nil {
			break
		}
	}
}
