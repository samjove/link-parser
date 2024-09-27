package main

import (
	"fmt"
	"strings"

	"github.com/samjove/link-parser/parser"
)

var example = `
<html>
<body>
	<h1>Hello</h1>
	<a href="/google.com">Test Link</a>
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