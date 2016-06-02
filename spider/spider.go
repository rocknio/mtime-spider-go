package spider

import (
    "errors"
    "fmt"
    "time"
)

const (
    uninit = iota
    inited
    started
    end
)

type Spider struct {
    start_urls          []string
    start_task_count    int
    finish_task_count   int
    status              int
    quit_chan           chan int
    start_time          time.Time
    end_time            time.Time
}

func (s *Spider) Init() (err error) {
    if s.status != uninit {
        err_string := fmt.Sprintf("Spider is already inited! Spider status is %+v", s.status)
        return errors.New(err_string)
    }

    s.start_urls = make([]string, 0, 100)
    s.quit_chan = make(chan int, 100)
    s.start_task_count = 0
    s.finish_task_count = 0
    s.status = inited

    return nil
}

func (s *Spider) Set_start_urls(url string) (err error) {
    s.start_urls = append(s.start_urls, url)
    return nil
}

func (s *Spider) Start_crawl() (err error) {
    s.status = started
    s.start_time = time.Now()

    for _, value := range s.start_urls {
        s.start_task_count++
        go Do_crawl(value, s.quit_chan)
    }

    // 判断是否所有协程都执行完
    for true {
        <-s.quit_chan
        s.finish_task_count++

        if s.finish_task_count == s.start_task_count {
            break
        }
    }

    s.status = end
    s.end_time = time.Now()

    return nil
}
