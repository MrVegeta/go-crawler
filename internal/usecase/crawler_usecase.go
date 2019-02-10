package usecase

import "os"

//go:generate mockgen -source=./crawler_usecase.go -destination=../mock/mock_crawler_usecase.go -package=mock


type CrawlerUseCase interface {
	Start(url string)
	Stop()
}

type crawlerUseCase struct {
	cc CrawlerCtrl
}

// NewCrawlerUseCase ...
func NewCrawlerUseCase(cc CrawlerCtrl) CrawlerUseCase {
	return &crawlerUseCase{
		cc: cc,
	}
}

func (cu crawlerUseCase) Start(url string) {
	worklist := make(chan []string)

	// Start with the command-line arguments.
	go func() { worklist <- os.Args[1:] }()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				go func(link string) {
					worklist <- cu.cc.Crawl(link)
				}(link)
			}
		}
	}
}

func (cu crawlerUseCase) Stop() {
	panic("implement me")
}