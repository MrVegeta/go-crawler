package main

import (
	"os"

	".."
)


func main() {
	cw  := crawler.NewCrawler()
	cw.Start(os.Args[1])
}
