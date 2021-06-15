package link

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
)

func ParseHtml(data io.Reader) ([]linkData, error) {
	containt, err := html.Parse(data)
	if err != nil {
		return nil, err
	}

	nodes := linkNodes(containt)
	var links []linkData
	for _, node := range nodes {
		links = append(links, buildLinks(node))
	}

	// dfs(containt, "")
	return links, nil

}

func buildLinks(node *html.Node) linkData {
	var return_link linkData
	for _, attr := range node.Attr {
		if attr.Key == "href" {
			return_link.Href = "Href: " + attr.Val
			break
		}
	}
	return_link.Text = "Text: " + text(node)
	return return_link
}

func text(node *html.Node) string {
	if node.Type == html.TextNode {
		return node.Data
	}

	if node.Type != html.ElementNode {
		return ""
	}

	var ret_text string
	for node_bite := node.FirstChild; node_bite != nil; node_bite = node_bite.NextSibling {
		ret_text += text(node_bite) + " "
	}
	return strings.Join(strings.Fields(ret_text), " ")
}

func linkNodes(node *html.Node) []*html.Node {
	var return_item []*html.Node
	if node.Type == html.ElementNode && node.Data == "a" {
		return []*html.Node{node}
	}

	for node_bite := node.FirstChild; node_bite != nil; node_bite = node_bite.NextSibling {
		return_item = append(return_item, linkNodes(node_bite)...)
	}
	return return_item
}

func dfs(node *html.Node, padding string) {
	found := node.Data
	if node.Type == html.ElementNode {
		found = "<" + found + ">"
	}
	fmt.Println(padding, found)
	for row := node.FirstChild; row != nil; row = row.NextSibling {
		dfs(row, padding+"  ")
	}
}

func Display(links []linkData) {
	for _, data := range links {
		fmt.Println(data.Href)
		fmt.Println(data.Text)
	}
}

type linkData struct {
	Href string
	Text string
}
