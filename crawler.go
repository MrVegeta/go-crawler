package crawler

import (
	"github.com/MrVegeta/go-crawler/internal/adapter"
	"github.com/MrVegeta/go-crawler/internal/usecase"
)

type Crawler interface {
	Start(url string)
	Stop()
}

type crawler struct {
	cu usecase.CrawlerUseCase
}


//
func NewCrawler() Crawler {
	cc := adapter.NewCrawlerCtrl()
	cu := usecase.NewCrawlerUseCase(cc)
	return &crawler{
		cu : cu,
	}
}

func (c crawler) Start(url string) {
	c.cu.Start(url)
}

func (c crawler) Stop() {
	c.cu.Stop()
}