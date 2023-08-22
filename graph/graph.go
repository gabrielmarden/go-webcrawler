package graph

import (
	"github.com/gabrielmarden/webcrawler/data"
)

//Traverse a graph using Breadth First Search algorithm to reach out all nodes.
//As input, the PROCESSOR functions to handle workload, the initial WORKLIST of URLs to search, the LIMIT to control the number of results
//and the KEYWORD to be searched in the pages found in each node
func TraverseBFS(processor func(item string, keyword string) ([]string, bool), worklist []string, limit int, keyword string) data.Set {
	nodes := data.NewSet()
	seen := data.NewSet()
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if limit > 0 && nodes.Length() >= limit {
				return nodes
			}

			if !seen.Contains(item) {
				seen.Add(item)
			} else {
				continue
			}

			if list, ok := processor(item, keyword); !nodes.Contains(item) {
				if ok {
					nodes.Add(item)
				}
				worklist = append(worklist, list...)
			}
		}
	}

	return nodes
}
