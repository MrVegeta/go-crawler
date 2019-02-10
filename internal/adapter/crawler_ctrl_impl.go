package adapter

import (
	"fmt"
	"golang.org/x/net/html"
	"gopl.io/ch5/links"
	"log"
	"net/http"

	"../usecase"
)

//go:generate mockgen -source=./crawler_ctrl_impl.go -destination=../mock/mock_crawler_ctrl_impl.go -package=mock

type crawlerCtrl struct {

}

func NewCrawlerCtrl() usecase.CrawlerCtrl {
	return &crawlerCtrl{}
}

func (cc crawlerCtrl) Crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func (cc crawlerCtrl) Extract(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				links = append(links, link.String())
			}
		}
	}
	cc.forEachNode(doc, visitNode, nil)
	return links, nil
}

func (cc crawlerCtrl) forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		cc.forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}