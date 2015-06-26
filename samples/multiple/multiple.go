package main

import (
	"encoding/json"
	"fmt"
)

// BEGIN OMIT
type Author struct{ ID, Name string }

type Book struct{ ID, Title string }

var data = []byte(`[
  {"id":"1a","name":"Neal Stephenson"},
  {"id":"1a","title":"Cryptonomicon"}
]`)

func main() {
	var response []json.RawMessage
	_ = json.Unmarshal(data, &response) // ignored error
	var author Author
	_ = json.Unmarshal(response[0], &author) // ignored error
	fmt.Printf("%+v\n", author)
	var book Book
	_ = json.Unmarshal(response[1], &book) // ignored error
	fmt.Printf("%+v\n", book)
}
