package main

import (
	"fmt"
	"os"

	"github.com/gabrielmarden/webcrawler/config"
	"github.com/gabrielmarden/webcrawler/graph"
	"github.com/gabrielmarden/webcrawler/links"
)

func main() {
	c, err := config.NewConfig(os.Args[1:])
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
	for _, link := range links.GetAll() {
		fmt.Println(link)
	}
	fmt.Printf("webcrawler: the search found %d results", links.Length())
}

func crawl(url string, keyword string) ([]string, bool) {
	list, ok, err := links.Extract(url, keyword)
	if err != nil {
		fmt.Println(err)
	}
	return list, ok

}
