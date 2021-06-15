package main

import (
	"flag"
	"io/ioutil"
	"log"
	"strings"

	"github.com/link"
)

var fileName *string

func init() {
	fileName = flag.String("fname", "ex1.txt", "Yml File containing URL and Paths")
	flag.Parse()
}

func readHtmlFile(str string) []byte {
	content, err := ioutil.ReadFile(str)
	if err != nil {
		log.Printf("HTML File.Get err   #%v ", err)
	}

	return content
}

func main() {
	htm := readHtmlFile(*fileName)
	html := strings.NewReader(string(htm))
	links, err := link.ParseHtml(html)
	if err != nil {
		panic(err)
	}
	link.Display(links)
}
