package parser

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

// Represents a HTML link
type Link struct {
	URL string
	Label string
}

// Takes an HTML doc and returns a slice of Links parsed from it
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	traverse(doc, "")
	return nil, nil 
}

func traverse(n *html.Node, padding string) {
	data := n.Data
	if n.Type == html.ElementNode {
		data = "<" + data + ">"
	}
	fmt.Println(padding, data)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		traverse(c, padding + "  ")
	}
}