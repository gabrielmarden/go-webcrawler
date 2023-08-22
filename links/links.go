package links

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

//Receives the URL and KEYWORD as argument, and with this information will search throughout the page all anchor tags "a" for the URLs and the pages
//that contains the keyword searched. Returns a slice of links found, a boolean value related to the searched keyword, and an error in case of failures.
func Extract(url string, keyword string) ([]string, bool, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, false, fmt.Errorf("checking url %s, but receive status %d", url, resp.StatusCode)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, false, fmt.Errorf("error during parsing %s as HTML: %v", url, err)
	}

	var links []string
	ok := false
	visitedNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}

				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue
				}

				if link.Host == resp.Request.Host {
					links = append(links, link.String())
				}
			}
		}

		if n.Type == html.TextNode && strings.Contains(strings.ToLower(n.Data), strings.ToLower(keyword)) {
			ok = true
		}
	}

	forEachNode(doc, visitedNode, nil)

	return links, ok, nil

}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

//Search the page found using the URL argument, for the specific KEYWORD. Returns a boolean to show this fact, and an error in case of failures.
func Scan(url string, keyword string) (bool, error) {
	resp, err := http.Get(url)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("checking url %s, but receive status %d", url, resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, fmt.Errorf("error during reading %s: %v", url, err)
	}

	pageContent := string(body)
	return strings.Contains(strings.ToLower(pageContent), strings.ToLower(keyword)), nil
}
