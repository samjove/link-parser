package main

import (
	"fmt"
	"strings"

	"github.com/samjove/link-parser/parser"
)

var example = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Sample HTML with Links</title>
</head>
<body>
    <h1>Welcome to My Website</h1>
    <p>This is a sample paragraph with a link to <a href="https://www.example.com">Example</a>.</p>
    <p>Check out my <a href="https://www.portfolio.com">Portfolio</a> for more projects.</p>
    <p>For more information, visit <a href="https://www.documentation.com">Documentation</a>.</p>
    <footer>
        <p>Contact us at <a href="mailto:info@example.com">info@example.com</a>.</p>
        <p>Follow us on <a href="https://www.socialmedia.com">Social Media</a>.</p>
    </footer>
</body>
</html>
`

func main() {
	r := strings.NewReader(example)
	links, err := parser.Parse(r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", links)
}