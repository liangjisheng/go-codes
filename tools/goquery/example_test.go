package goquery

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestExampleScrape(t *testing.T) {
	// Request the HTML page.
	res, err := http.Get("http://metalsucks.net")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find(".left-content article .post-title").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		title := s.Find("a").Text()
		fmt.Printf("Review %d: %s\n", i, title)
	})
}

// This example shows how to use NewDocumentFromReader from a file.
func ExampleNewDocumentFromReader_file() {
	// create from a file
	f, err := os.Open("some/file.html")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		log.Fatal(err)
	}
	// use the goquery document...
	_ = doc.Find("h1")
}

// This example shows how to use NewDocumentFromReader from a string.
func ExampleNewDocumentFromReader_string() {
	// create from a string
	data := `
<html>
	<head>
		<title>My document</title>
	</head>
	<body>
		<h1>Header</h1>
	</body>
</html>`

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}
	header := doc.Find("h1").Text()
	fmt.Println(header)

	// Output: Header
}

func ExampleSingle() {
	html := `
<html>
  <body>
    <div>1</div>
    <div>2</div>
    <div>3</div>
  </body>
</html>
`
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatal(err)
	}

	// By default, the selector string selects all matching nodes
	multiSel := doc.Find("div")
	fmt.Println(multiSel.Text())

	// Using goquery.Single, only the first match is selected
	singleSel := doc.FindMatcher(goquery.Single("div"))
	fmt.Println(singleSel.Text())

	// Output:
	// 123
	// 1
}
