package spider

import (
    "net/http"
    "log"
    "io/ioutil"
)

func Do_crawl(url string, quit chan int) {
    resp, err := http.Get(url)
    if err != nil {
        log.Panicf("http get fail! url = %v, err = %v", url, err)
    } else {
        defer resp.Body.Close()
        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            log.Panicf("Get Http Body fail! err = %v", err)
        }
        log.Printf("resp = %v\n body = %v\n err = %v", resp, string(body), err)
    }


    quit <- 1
}
