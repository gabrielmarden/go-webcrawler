package main

import (
	"flag"
	"fmt"

	"github.com/gabrielmarden/webcrawler/config"
	"github.com/gabrielmarden/webcrawler/graph"
	"github.com/gabrielmarden/webcrawler/links"
	"github.com/gabrielmarden/webcrawler/util"
)

var (
	url       = flag.String("url", "", "base url to start the webcrawler")
	keyword   = flag.String("keyword", "", "keyword searched in the pages")
	maxResult = flag.Int("max", -1, "number of max results found for the webcrawler")
)

func main() {

	flag.Parse()

	c, err := config.NewConfig(*url, *keyword, *maxResult)
	if err != nil {
		fmt.Println(err)
		return
	}

	if errs := c.Validate(config.ValidateRequiredParameter, config.ValidateKeyword, config.ValidateURL); len(errs) > 0 {
		fmt.Println("errors during validation of the input parameters. please check below:")
		for i, err := range errs {
			fmt.Printf("[%d] error: %v", i, err)
		}
		return
	}

	links := graph.TraverseBFS(crawl, []string{c.URL}, c.MaxResult, c.Keyword)
	fmt.Printf("webcrawler: the search found %d results", links.Length())

	util.WriteDataToFile(links.GetAll(), "webcrawler")
}

func crawl(url string, keyword string) ([]string, bool) {
	list, ok, err := links.Extract(url, keyword)
	if err != nil {
		fmt.Println(err)
	}
	return list, ok

}
