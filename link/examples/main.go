package main

import (
	"fmt"
	"gophercises/link"
	"os"
)

var ex1 = `
<html>
<body>
  <h1>Hello!</h1>
  <p>Lorem ipsum: <a href="/other-page">A <strong>link</strong> to <span>another page</span></a></p>
  <a href="/other-page2">A link to a second page</a>
</body>
</html>
`

func main() {
	r, err := os.Open("examples/ex4.html")
	if err != nil {
		panic(err)
	}
	//r := strings.NewReader(ex1)

	links, err := link.Parse(r)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", links)
}
