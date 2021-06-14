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
<a href="/other-page">A link to another page</a>
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
