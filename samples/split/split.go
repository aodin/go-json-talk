package main

import (
	"encoding/json"
	"fmt"
)

// BEGIN OMIT
type Author struct{ Name string }

type Book struct{ Title string }

var data = []byte(`[
  {"name":"Neal Stephenson"},
  {"title":"Cryptonomicon"}
]`)

func main() {
	var responses = []struct {
		Author
		Book
	}{}
	_ = json.Unmarshal(data, &responses) // ignored error
	fmt.Printf("%+v\n", responses[0].Author)
	fmt.Printf("%+v\n", responses[1].Book)
}
