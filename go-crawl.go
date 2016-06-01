package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	// 获取参数
	if len(os.Args) < 2 {
		fmt.Println("请输入爬取的起始网页！")
		return
	}

	var start_url string
	start_url = os.Args[1]

	// 调用spider包的interface设置参数

	// quit channel
	quit_ch = make(chan int)

	// 执行爬取
	start_time := time.Now()

	<-quit_ch
	log.Printf(
		"%s\t%s",
		start_url,
		time.Since(start_time),
	)
}
