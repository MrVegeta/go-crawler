package usecase

//go:generate mockgen -source=./crawler_ctrl.go -destination=../mock/mock_crawler_ctrl.go -package=mock

type CrawlerCtrl interface {
	Crawl(url string) []string
	Extract(url string) ([]string, error)
}