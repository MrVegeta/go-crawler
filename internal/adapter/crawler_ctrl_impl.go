package adapter

import (
	"../usecase"
	"fmt"
	"golang.org/x/net/html"
)

//go:generate mockgen -source=./crawler_ctrl_impl.go -destination=../mock/mock_crawler_ctrl_impl.go -package=mock

type crawlerCtrl struct {

}

func NewCrawlerCtrl() usecase.CrawlerCtrl {
	return &crawlerCtrl{}
}

func (cc crawlerCtrl) Crawl(url string) []string {
	fmt.Println(url)
	return nil
}

func (cc crawlerCtrl) Extract(url string) ([]string, error) {
	return nil, nil
}

func (cc crawlerCtrl) forEachNode(n *html.Node, pre, post func(n *html.Node)) {
}