package spider

import (
    "fmt"
    "strings"

    "github.com/PuerkitoBio/goquery"
)

func parseHtml(url string, pipeline chan int) (item *Item, err error) {
    html_doc, err := goquery.NewDocument(url)
    if err != nil {
        fmt.Printf("parse html fail! err = %v", err)
        return new(Item), err
    }
    parse_item := new(Item)

    // 片名
    parse_item.Title = html_doc.Find("head title").Text()
    movie_info_l := html_doc.Find("dl.info_l")
    info_node := movie_info_l.Find("dd.__r_c_")
    info_node.Each(func(i int, s *goquery.Selection) {
        if strings.Contains(s.Text(), "导演") {
            parse_item.Directors = s.Text()
            parse_item.Directors = strings.Replace(parse_item.Directors, " ", "", -1)
            parse_item.Directors = strings.Replace(parse_item.Directors, "\n", "", -1)
        } else if strings.Contains(s.Text(), "编剧") {
            parse_item.Actors = append(parse_item.Actors, s.Text())
        }
    })

    return parse_item, nil
}

func Do_crawl(url string, quit chan int, pipeline_ch chan int) {
    //resp, err := http.Get(url)
    //if err != nil {
    //    fmt.Printf("http get fail! url = %v, err = %v", url, err)
    //} else {
    //    defer resp.Body.Close()
    //    body, err := ioutil.ReadAll(resp.Body)
    //    if err != nil {
    //        fmt.Printf("Get Http Body fail! err = %v", err)
    //    }
    //    fmt.Printf("resp = %v\n body = %s\n err = %v\n\n\n", resp, string(body), err)
    //
    //    // 解析html
    //    parseHtml(string(body), pipeline_ch)
    //}

    // 解析html
    parseHtml(url, pipeline_ch)

    //quit <- 1
}
