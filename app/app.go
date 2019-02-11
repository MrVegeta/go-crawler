package main

import (
	"os"

	"github.com/MrVegeta/go-crawler"
)


func main() {
	cw  := crawler.NewCrawler()
	cw.Start(os.Args[1])
}
