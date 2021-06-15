package main

import (
	"fmt"
	"strings"

	"github.com/link"
)

func main() {
	htm := `
<html>
<body>
<h1>Hello!</h1>
<a href="/other-page">A link to another page<span>span text</span></a>
<a href="/other-page2">A link to another page2</a>
</body>
</html>
`
	html := strings.NewReader(htm)
	links, err := link.ParseHtml(html)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", links)
}
