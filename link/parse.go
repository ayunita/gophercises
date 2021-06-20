package link

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
)

// Link represents a link (<a href="...">) in an HTML document.
type Link struct {
	Href string
	Text string
}

// Parse will take in an HTML document and
// will return a slice of links parsed from it.
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	nodes := linkNodes(doc)
	var links []Link
	for _, node := range nodes {
		links = append(links, buildLink(node))
	}
	return links, nil
}

// text will take html node and
// will return all text inside the node in one string
func text(n *html.Node) string {
	// Base case 1: if text node, then return the string
	if n.Type == html.TextNode {
		return n.Data
	}
	// Base case 2: if !element node, then return an empty string
	if n.Type != html.ElementNode {
		return ""
	}

	var ret string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		fmt.Println(c)
		ret += text(c) + " "
	}

	return strings.Join(strings.Fields(ret), " ")
}

// buildLink will take html node and
// will return Link object
func buildLink(n *html.Node) Link {
	var ret Link
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			ret.Href = attr.Val
		}
	}
	ret.Text = text(n)
	return ret
}

// linkNodes will take html node and
// will return slices of link (<a>) html node
func linkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}
	var ret []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, linkNodes(c)...)
	}
	return ret
}
