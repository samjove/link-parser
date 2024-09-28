package parser

import (
	"io"
	"strings"

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
	linkNodes := findLinks(doc)
	var links []Link
	for _, link := range linkNodes {
		links = append(links, buildLink(link))
	}
	return links, nil 
}

func findLinks(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}
	var links []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = append(links, findLinks(c)...)
	}
	return links
}

func buildLink(n *html.Node) Link {
	var link Link
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			link.URL = attr.Val
			break
		}
	}
	link.Label = findText(n)
	return link
}

func findText(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	if n.Type != html.ElementNode {
		return ""
	}

	var text string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		text += findText(c) + " "
	}

	return strings.Join(strings.Fields(text), " ")
}