package link

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

func ParseHtml(data io.Reader) ([]linkData, error) {
	containt, err := html.Parse(data)
	if err != nil {
		return nil, err
	}

	dfs(containt, "")
	return nil, nil

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

type linkData struct {
	Href string
	Text string
}
