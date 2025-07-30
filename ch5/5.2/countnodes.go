package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func main() {
	_, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "countnodes: %v\n", err)
		os.Exit(1)
	}
}

func count(nodes map[string]int, n *html.Node) map[string]int {
	if n == nil {
		return nodes
	}

	if n.Type == html.ElementNode {
		tagName := n.Data
		if _, ok := nodes[tagName]; !ok {
			nodes[tagName] = 0
		}
		nodes[tagName]++
	}

	if n.FirstChild != nil {
		nodes = count(nodes, n.FirstChild)
	}

	if n.NextSibling != nil {
		nodes = count(nodes, n.NextSibling)
	}

	return nodes
}
