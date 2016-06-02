package main

import (
    "fmt"
    "log"
    "mtime-spider-go/spider"
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
    spider_info := new(spider.Spider)
    err := spider_info.Init()
    if err != nil {
        fmt.Println("初始化失败，err = %+v", err)
        return
    }

    // 执行爬取
    start_time := time.Now()
    spider_info.Set_start_urls(start_url)
    spider_info.Start_crawl()

    log.Printf(
        "%s\t%s",
        start_url,
        time.Since(start_time),
    )
}
