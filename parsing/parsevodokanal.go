package parsing

import (
	"strings"
	"golang.org/x/net/html"
)

func ExtractText(n *html.Node) []string {
	var result []string
	var extract func(*html.Node)
	
	extract = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "td" {
			text := strings.TrimSpace(getNodeText(node))
			if text != "" {
				result = append(result, text)
			}
		}
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			extract(c)
		}
	}
	
	extract(n)
	return result
}

func getNodeText(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	var text string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		text += getNodeText(c)
	}
	return text
}