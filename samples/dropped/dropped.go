package main

import (
	"encoding/json"
	"fmt"
)

// START OMIT
type Author struct{ ID, Name string }

type Book struct{ ID, Title string }

var data = []byte(`[
  {"id": "1a", "name":"Neal Stephenson"},
  {"id": "1a", "title":"Cryptonomicon"}
]`)

func main() {
	var responses = []struct {
		Author
		Book
	}{}
	_ = json.Unmarshal(data, &responses) // ignored error
	fmt.Printf("%+v\n", responses[0].Author)
	fmt.Printf("%+v\n", responses[0].Book)
}

// END OMIT
