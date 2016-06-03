package spider

type Item struct {
    Id         int
    Title      string
    Link       string
    Directors  string
    Actors     []string
    image_urls []string
}

type Items []Item
